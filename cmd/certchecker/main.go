package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/antihax/optional"
	rainbond "github.com/goodrain/openapi-go"
	"github.com/hongyaa-tech/rainbond-cert-controller/config"
	"github.com/hongyaa-tech/rainbond-cert-controller/notify"
	"github.com/hongyaa-tech/rainbond-cert-controller/sslcheck"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
}

const (
	CHECK_DAYS = 25 // if domain cert exipres less than 25 days, then alert
)

func main() {
	rainbond_client := rainbond.NewAPIClient(rainbond.NewConfiguration())
	ctx := context.WithValue(context.Background(), rainbond.ContextAPIKey, rainbond.APIKey{
		Key: config.Cfg.Rainbond.ApiKey,
	})
	// list all tenants

	gwRules, _, err := rainbond_client.OpenapiGatewayApi.OpenapiV1HttpdomainsList(ctx, &rainbond.OpenapiGatewayApiOpenapiV1HttpdomainsListOpts{
		AutoSsl: optional.NewBool(true),
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
