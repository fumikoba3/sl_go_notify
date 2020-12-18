package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type slackParams struct {
	Text      string `json:"text"`
	UserName  string `json:"username"`
	Channel   string `json:"channel"`
	IconEmoji string `json:"icon_emoji"`
	// IconUrl   string `json:"icon_url"`
}

func createParam(message string, channelName string) slackParams {

	username := os.Getenv("userName")
	iconEmoji := os.Getenv("iconEmoji")
	// iconUrl := os.Getenv("iconUrl")
	//iconEmoji<iconUrl
	//faultlineは以下url
	//TODO https://faultline.github.io/faultline/icon.png'

	sp := &slackParams{
		Text:      message,
		UserName:  username,
		Channel:   channelName,
		IconEmoji: iconEmoji,
	}

	return *sp
}

func Notify(message string, channelName string) error {

	slackUrl := os.Getenv("slackUrl")

	sp := createParam(message, channelName)
	jsonSp, _ := json.Marshal(sp)
	fmt.Println(string(jsonSp))

	resp, err := http.PostForm(
		slackUrl,
		url.Values{"payload": {string(jsonSp)}},
	)

	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return nil
}
