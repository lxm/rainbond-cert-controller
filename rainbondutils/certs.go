package rainbondutils

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/antihax/optional"
	"github.com/go-acme/lego/v3/certificate"
	rainbond "github.com/goodrain/openapi-go"
)

type TeamCertInfo map[string]rainbond.CertificatesR

func ListTeamCerts(client *rainbond.APIClient, ctx context.Context, teamInfo rainbond.TeamInfo) TeamCertInfo {
	var currentPage = 1
	var pageSize = 100
	var totalPage int
	certList := make(map[string]rainbond.CertificatesR)
	for {
		res, _, err := client.OpenapiTeamApi.OpenapiV1TeamsCertificatesList(ctx, teamInfo.TenantId, &rainbond.OpenapiTeamApiOpenapiV1TeamsCertificatesListOpts{
			Page:     optional.NewFloat32(float32(currentPage)),
			PageSize: optional.NewFloat32(float32(pageSize)),
		})
		totalPage = int(math.Ceil(float64(res.Total) / float64(pageSize)))
		if err != nil {
			fmt.Println(err.Error())
			break
			continue
		}
		for _, cert := range res.List {
			certList[cert.Alias] = cert
		}
		if totalPage <= currentPage {
			break
		} else {
			currentPage++
		}
	}
	return certList
}

func UpdateOrCreateTeamCert(client *rainbond.APIClient, ctx context.Context, tenantId string, teamCert *rainbond.CertificatesR, cert *certificate.Resource) (rainbond.TeamCertificatesR, *http.Response, error) {
	// client.OpenapiTeamApi.OpenapiV1TeamsCertificatesUpdate()
	// client.OpenapiTeamApi.OpenapiV1TeamsCertificatesCreate()
	teamCertData := rainbond.TeamCertificatesC{
		Alias:           cert.Domain,
		Certificate:     string(cert.Certificate),
		PrivateKey:      string(cert.PrivateKey),
		CertificateType: "服务端证书",
	}
	if teamCert != nil {
		fmt.Println("update cert")
		return client.OpenapiTeamApi.OpenapiV1TeamsCertificatesUpdate(ctx, fmt.Sprintf("%d", teamCert.Id), tenantId, teamCertData)
	} else {
		fmt.Println("create cert")
		// fmt.Println(teamCertData)
		return client.OpenapiTeamApi.OpenapiV1TeamsCertificatesCreate(ctx, tenantId, teamCertData)
	}
}
