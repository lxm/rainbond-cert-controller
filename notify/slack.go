package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hongyaa-tech/rainbond-cert-controller/config"
)

func notifySlack(notifyCfg config.Notify, msgStr string) error {
	defaultMessage := Default{
		Channel: notifyCfg.Channel,
		Title:   msgStr,
	}
	buffer := bytes.NewBuffer(nil)
	if err := json.NewEncoder(buffer).Encode(defaultMessage.Conver()); err != nil {
		return err
	}
	re, err := http.Post(notifyCfg.URL, "application/json", buffer)
	if err != nil {
		return err
	}
	if re.StatusCode >= 300 {
		defer re.Body.Close()
		re, _ := ioutil.ReadAll(re.Body)
		return fmt.Errorf(string(re))
	}
	return nil
}

//Default default
type Default struct {
	Channel    string       `json:"channel"`
	Title      string       `json:"title"`
	Describe   string       `json:"describe"`
	Extensions []*Extension `json:"extensions"`
}

//Extension ext
type Extension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//Conver conver
func (d *Default) Conver() *Message {
	var texts []*Text
	for _, item := range d.Extensions {
		if item.Name == "" {
			continue
		}
		texts = append(texts, &Text{
			Type: "mrkdwn",
			Text: fmt.Sprintf("*%s:*\n%s", item.Name, item.Value),
		})
	}
	defectMessage := &Message{
		Channel: d.Channel,
		Blocks: []*Block{
			&Block{
				Type: "section",
				Text: &Text{
					Type: "mrkdwn",
					Text: fmt.Sprintf("*%s*\n%s", d.Title, d.Describe),
				},
			},
		},
	}
	if len(texts) > 0 {
		defectMessage.Blocks = append(defectMessage.Blocks, &Block{
			Type:   "section",
			Fields: texts,
		})
	}
	return defectMessage
}

//Message slack message
type Message struct {
	Channel string   `json:"channel"`
	Blocks  []*Block `json:"blocks"`
}

//Block block
type Block struct {
	Type   string  `json:"type"`
	Text   *Text   `json:"text,omitempty"`
	Fields []*Text `json:"fields,omitempty"`
}

//Text text
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
