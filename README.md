# rainbond-cert-controller

用于自动化申请rainbond集群https策略证书

## 参数配置

### 控制台配置文件
```
{
    "aliyun_hongyaa":{
        "provider":"alidns",
        "env":{
            "ALICLOUD_POLLING_INTERVAL":"2",
            "ALICLOUD_SECRET_KEY":"ali sk",
            "ALICLOUD_PROPAGATION_TIMEOUT":"300",
            "ALICLOUD_ACCESS_KEY":"ali ak"
        }
    },
    "dnspod_hongyaa":{
        "provider":"dnspod",
        "env":{
            "DNSPOD_PROPAGATION_TIMEOUT":"100",
            "DNSPOD_HTTP_TIMEOUT":"100",
            "DNSPOD_API_KEY":"apiid,apikey"
        }
    }
}
```
其中provider和env参考 [lego-dns](https://go-acme.github.io/lego/dns/)

### 控制器运行时配置

以下为环境变量

```
RAINBOND_API_KEY rainbond openapi key
ACME_EMAIL let's encrypt email
ACME_KEY_TYPE 可选，默认为RSA4096
ACME_DIR_URL 可选，默认为https://acme-v02.api.letsencrypt.org/directory
ACME_SRORAGE_PATH 可选，用于存放认证信息，默认/opt/rainbond-cert-controller/storage
DINGTALK_AK 可选，用于钉钉通知
DINGTALK_SK 可选，用于钉钉通知
```

docker run参考

```
docker run --rm -e ACME_EMAIL=luxingmin@hongyaa.com.cn \
-e RAINBOND_API_KEY=xxx \
-e DINGTALK_AK=xxx \
-e DINGTALK_SK=xxx \
hongyaa/rainbond-cert-controller:latest
```

## 功能列表
* 自动申请证书
* 到期自动续期
* 钉钉/Slack通知申请状态

## TODO
* 证书状态检查及异常通知