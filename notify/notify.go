package notify

import (
	"errors"
	"fmt"

	"github.com/CatchZeng/dingtalk/client"
	"github.com/CatchZeng/dingtalk/message"
	"github.com/hongyaa-tech/rainbond-cert-controller/config"
	"github.com/sirupsen/logrus"
)

func SendNotify(notifyName, msg string) error {
	notifyCfg, ok := config.Cfg.NotifyList[notifyName]
	if !ok {
		logrus.Error("config type error ", notifyName)
		return errors.New("config type error")
	}
	switch notifyCfg.Type {
	case "dingtalk":
		return notifyDingtalk(notifyCfg, msg)
	default:
		return errors.New("no support type " + notifyCfg.Type)
	}
}

func notifyDingtalk(notifyCfg config.Notify, msgStr string) error {
	dingtalk := client.DingTalk{
		AccessToken: notifyCfg.AccessToken,
		Secret:      notifyCfg.Secret,
	}
	msg := message.NewTextMessage().SetContent(msgStr)
	_, err := dingtalk.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
