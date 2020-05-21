package main

import (
	"context"
	"fmt"
	"strings"

	rainbond "github.com/goodrain/openapi-go"
	"github.com/hongyaa-tech/rainbond-cert-controller/config"
	"github.com/hongyaa-tech/rainbond-cert-controller/notify"
	"github.com/hongyaa-tech/rainbond-cert-controller/sslcheck"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
}

const (
	CHECK_DAYS = 25 // if domain cert exipres less than 25 days, then alert
)

func main() {
	// cronSvc := cron.New(cron.WithSeconds())
	cronSvc := cron.New()
	cronSvc.AddFunc("*/10 * * * *", clusterCertCheck)
	cronSvc.Run()
}

func clusterCertCheck() {
	rainbond_client := rainbond.NewAPIClient(rainbond.NewConfiguration())
	ctx := context.WithValue(context.Background(), rainbond.ContextAPIKey, rainbond.APIKey{
		Key: config.Cfg.Rainbond.ApiKey,
	})
	// list all tenants

	gwRules, _, err := rainbond_client.OpenapiGatewayApi.OpenapiV1HttpdomainsList(ctx, &rainbond.OpenapiGatewayApiOpenapiV1HttpdomainsListOpts{
		// AutoSsl: optional.NewBool(true),
	})
	if err != nil {
		msg := fmt.Sprintf("init certcheker list gatewat rules error ", err.Error())
		logrus.Error(msg)
		notify.SendNotify("default", msg)
	}
	for _, gwRule := range gwRules {
		if strings.HasPrefix(gwRule.DomainName, "*.") {
			logrus.Info(fmt.Sprintf("wildcard domain:[%s] do not support check", gwRule.DomainName))
			continue
		}
		if strings.Compare(gwRule.Protocol, "https") != 0 {
			continue
		}
		if strings.Contains(config.Cfg.DisableCheckCluster, gwRule.RegionName) {
			logrus.Info(fmt.Sprintf("rule: %s cluster: %s disabled by config, ignore", gwRule.DomainName, gwRule.RegionName))
			continue
		}
		expire, err := sslcheck.GetCertsExpire(gwRule.DomainName, "443")
		if err != nil {
			msg := fmt.Sprintf("domain:%s check with error:%s", gwRule.DomainName, err.Error())
			logrus.Error(msg)
			go notify.SendNotify("default", msg)
			continue
		}
		if expire < CHECK_DAYS*86400 {
			msg := fmt.Sprintf("domain:%s will expire in %d days(config: %d days), check auto sign", gwRule.DomainName, expire/86400, CHECK_DAYS)
			logrus.Info(msg)
			notify.SendNotify("default", msg)
		}
	}
}
