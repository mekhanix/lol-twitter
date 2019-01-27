package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/http"
	"net/url"
	"strings"
)

func setBody(body string) strings.Reader {
	msgData := url.Values{}
	msgData.Set("To", Env["SMS_TO"])
	msgData.Set("From", Env["SMS_FROM"])
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())
	return msgDataReader
}

func sendSms(variants []anaconda.Variant) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + Env["TWILIO_ACCOUNT_SID"] + "/Messages.json"
	var body string
	for _, variant := range variants {
		body += "\n" + variant.Url
	}
	if len(body) > 0 {
		msgData := setBody(body)
		client := &http.Client{}
		req, _ := http.NewRequest("POST", urlStr, &msgData)
		req.SetBasicAuth(Env["TWILIO_ACCOUNT_SID"], Env["TWILIO_AUTH_TOKEN"])
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := client.Do(req)

		if err != nil {
			fmt.Printf("%s", err.Error())
		}
		fmt.Println(resp.StatusCode)
	}

}
