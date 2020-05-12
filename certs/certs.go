package certs

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-acme/lego/v3/certcrypto"
	"github.com/go-acme/lego/v3/certificate"
	"github.com/go-acme/lego/v3/challenge/dns01"
	"github.com/go-acme/lego/v3/lego"
	"github.com/go-acme/lego/v3/providers/dns"
	"github.com/go-acme/lego/v3/registration"
	"github.com/hongyaa-tech/rainbond-cert-controller/acmeaccount"
	"github.com/hongyaa-tech/rainbond-cert-controller/config"
	"github.com/sirupsen/logrus"
)

func init() {
	config.Load()
}

func RequestCert(domain, domainAuthName string) (*certificate.Resource, error) {
	accountStroage := acmeaccount.NewAccountsStorage(config.Cfg.Acme.Email, config.Cfg.Acme.RootPath)

	account, err := accountStroage.CreateOrLoadAccount()
	if err != nil {
		log.Fatal("create account failed")
	}
	//privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//account := &Account{
	//	Email: config.Cfg.Acme.Email,
	//	key:   privateKey,
	//}
	legoCfg := lego.NewConfig(account)
	legoCfg.Certificate.KeyType = certcrypto.RSA4096
	legoClient, err := lego.NewClient(legoCfg)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	domainAuth, ok := config.Cfg.AuthList[domainAuthName]
	if !ok {
		fmt.Println("auth info not found")
		return nil, errors.New("auth info not found")
		// auth info not found
	}
	for envName, envVal := range domainAuth.Env {
		logrus.Info("set env\t", envName, envVal)
		os.Setenv(envName, envVal)
	}
	provider, err := dns.NewDNSChallengeProviderByName(domainAuth.Provider)

	if err != nil {
		fmt.Println("init provider failed")
		return nil, err
	}
	err = legoClient.Challenge.SetDNS01Provider(provider, dns01.AddRecursiveNameservers([]string{"114.114.114.114:53"}), dns01.AddDNSTimeout(60*time.Second))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	reg, err := legoClient.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
	account.Registration = reg
	request := certificate.ObtainRequest{
		Domains: []string{domain},
		Bundle:  true,
	}
	certificates, err := legoClient.Certificate.Obtain(request)
	if err != nil {
		fmt.Println("obtain cert failed\t" + err.Error())
		return nil, err
	}
	// jsonStr, err := json.Marshal(certificates)
	return certificates, nil
	// fmt.Println(string(certificates.Certificate))
	// fmt.Println(string(certificates.PrivateKey))
}
