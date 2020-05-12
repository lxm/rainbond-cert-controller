package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/antihax/optional"
	rainbond "github.com/goodrain/openapi-go"
	"github.com/hongyaa-tech/rainbond-cert-controller/certs"
	"github.com/hongyaa-tech/rainbond-cert-controller/config"
	"github.com/hongyaa-tech/rainbond-cert-controller/rainbondutils"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
}

func main() {
	rainbond_client := rainbond.NewAPIClient(rainbond.NewConfiguration())
	ctx := context.WithValue(context.Background(), rainbond.ContextAPIKey, rainbond.APIKey{
		Key: config.Cfg.Rainbond.ApiKey,
	})
	// list all tenants
	tenants, err := listAllTeams(rainbond_client, ctx)

	if err != nil {
		logrus.Fatal("listAllTeams failed\t" + err.Error())
	}
	_ = tenants
	gwRules, _, err := rainbond_client.OpenapiGatewayApi.OpenapiV1HttpdomainsList(ctx, &rainbond.OpenapiGatewayApiOpenapiV1HttpdomainsListOpts{
		AutoSsl: optional.NewBool(true),
	})

	var allTeamCertInfo map[string]rainbondutils.TeamCertInfo
	allTeamCertInfo = make(map[string]rainbondutils.TeamCertInfo)

	for _, tenant := range tenants {
		allTeamCertInfo[tenant.TenantName] = rainbondutils.ListTeamCerts(rainbond_client, ctx, tenant)
	}

	for _, gwRule := range gwRules {
		if gwRule.AutoSsl == true {
			var needRequestCert bool = false
			var existCert *rainbond.CertificatesR = nil
			teamCertInfo, ok := allTeamCertInfo[gwRule.TeamName]
			var certID int32 = 0
			if !ok {
				logrus.Info(fmt.Printf("team %s has no certs", gwRule.TeamName))
				needRequestCert = true
			} else {
				existCert = getExistCert(gwRule.DomainName, teamCertInfo)
				if existCert != nil {
					certID = existCert.Id
					timeExpire, err := time.Parse("2006-01-02T15:04:05", existCert.EndData)
					if err != nil {
						logrus.Error("parse cert expire error " + err.Error())
						needRequestCert = true
					} else if timeExpire.Unix()-time.Now().Unix() < int64(30*86400) { // renew in 30 days
						needRequestCert = true
					}
				} else {
					needRequestCert = true
				}
			}
			if certID == 0 {
				needRequestCert = true
			}
			if needRequestCert {
				certResource, err := certs.RequestCert(gwRule.DomainName, gwRule.AutoSslConfig)
				if err != nil {
					logrus.Error(fmt.Sprintf("RequestCert for domain %s err %s", gwRule.DomainName, err.Error()))
					continue
				}

				rainbondCertInfo, _, err := rainbondutils.UpdateOrCreateTeamCert(rainbond_client, ctx, gwRule.TenantId, existCert, certResource)
				if err != nil {
					logrus.Error(fmt.Sprintf("UpdateOrCreateTeamCert for domain %s err %s", gwRule.DomainName, err.Error()))
					continue
				}
				certID = rainbondCertInfo.Id
			}
			if gwRule.CertificateId == 0 {
				_, _, err := rainbond_client.OpenapiGatewayApi.OpenapiV1TeamsRegionsAppsHttpdomainsUpdate(ctx, fmt.Sprintf("%d", gwRule.AppId), gwRule.RegionName, gwRule.HttpRuleId, gwRule.TeamName, rainbond.PostHttpGatewayRule{
					CertificateId: certID,
					ServiceId:     gwRule.ServiceId,
					DomainName:    gwRule.DomainName,
				})
				if err != nil {
					gerr := err.(rainbond.GenericSwaggerError)
					logrus.Error(fmt.Sprintf("update domain rule for domain %s err %s", gwRule.DomainName, string(gerr.Body())))
				}
			}
		}
	}
}

func getExistCert(domain string, certInfo rainbondutils.TeamCertInfo) *rainbond.CertificatesR {
	// try wildcard
	if strings.HasPrefix("*", domain) {
		cert, ok := certInfo[domain]
		if ok {
			return &cert
		}
	} else {
		cert, ok := certInfo["*."+domain]
		if ok {
			return &cert
		}
	}
	cert, ok := certInfo[domain]
	if ok {
		return &cert
	} else {
		return nil
	}
}
func listAllTeams(client *rainbond.APIClient, ctx context.Context) (tenants []rainbond.TeamInfo, err error) {
	page := 1
	pageSize := 100
	for {
		teamResp, _, err := client.OpenapiTeamApi.OpenapiV1TeamsList(ctx, &rainbond.OpenapiTeamApiOpenapiV1TeamsListOpts{
			Query:    optional.EmptyString(),
			Page:     optional.NewString(fmt.Sprintf("%d", page)),
			PageSize: optional.NewString("100"),
		})
		if err != nil {
			return tenants, err
		}
		// str, err := json.Marshal(teamResp)
		tenants = append(tenants, teamResp.Tenants...)
		page++
		if teamResp.Total <= int32(pageSize*page) {
			break
		}
	}
	return tenants, nil
}
