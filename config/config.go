package config

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	rainbond "github.com/goodrain/openapi-go"
	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Rainbond            Rainbond          `json:"rainbond"`
	Acme                Acme              `json:"acme"`
	AuthList            map[string]Auth   `json:"auth_list"`
	NotifyList          map[string]Notify `json:"notify_list"`
	DisableCheckCluster string            `json:"disable_check_cluster"`
}

var once sync.Once
var Cfg = &Config{}

type Rainbond struct {
	ApiKey string `json:"api_key"`
}

type Acme struct {
	Email    string `json:"email"`
	KeyType  string `json:"key_type"`
	CADirUrl string `json:"ca_dir_url"`
	RootPath string `json:"root_path"`
}

type Auth struct {
	Provider string            `json:"provider"`
	Env      map[string]string `json:"env"`
}

type Notify struct {
	Type        string `json:"type"`
	URL         string `json:"url,omitempty"`
	Channel     string `json:"channel,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	Secret      string `json:"secret,omitempty"`
}

//type CheckNotify map[string]string

func Load() {
	once.Do(func() {
		configor.New(&configor.Config{
			Debug:              false,
			AutoReload:         false,
			AutoReloadInterval: time.Minute,
		}).Load(Cfg, "cfg.json")
		loadSSLAuthFromRainbond()
	})
}

func loadSSLAuthFromRainbond() {
	rainbond_client := rainbond.NewAPIClient(rainbond.NewConfiguration())
	ctx := context.WithValue(context.Background(), rainbond.ContextAPIKey, rainbond.APIKey{
		Key: Cfg.Rainbond.ApiKey,
	})
	ret, _, err := rainbond_client.OpenapiEntrepriseApi.OpenapiV1ConfigsList(ctx)
	if err != nil {
		logrus.Error("load ssl auto info from rainbond with error " + err.Error())
		return
	}
	json.Unmarshal(ret.AutoSsl.Value, &(Cfg.AuthList))
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
