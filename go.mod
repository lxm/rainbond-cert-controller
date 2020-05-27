module github.com/hongyaa-tech/rainbond-cert-controller

go 1.13

replace github.com/goodrain/openapi-go => ./packages/clients/go

require (
	github.com/CatchZeng/dingtalk v1.0.0
	github.com/antihax/optional v1.0.0
	github.com/go-acme/lego/v3 v3.5.0
	github.com/goodrain/openapi-go v0.0.0-00010101000000-000000000000
	github.com/jinzhu/configor v1.1.1
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2
)
