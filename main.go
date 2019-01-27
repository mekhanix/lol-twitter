package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

// Env contain envar from .env
var Env = LoadEnv()

func main() {
	anaconda.SetConsumerKey(Env["TWITTER_CONSUMER_KEY"])
	anaconda.SetConsumerSecret(Env["TWITTER_SECRET_CONSUMER_KEY"])
	api := anaconda.NewTwitterApiWithCredentials(
		Env["TWITTER_ACCESS_TOKEN"], Env["TWITTER_SECRET_ACCESS_TOKEN"],
		Env["TWITTER_CONSUMER_KEY"], Env["TWITTER_SECRET_CONSUMER_KEY"],
	)

	v := url.Values{}
	v.Set("follow", Env["TWITTER_ID"])
	s := api.PublicStreamFilter(v)
	for t := range s.C {
		switch value := t.(type) {
		case anaconda.Tweet:
			{
				if value.Text == "lol" && value.QuotedStatus != nil {
					entities := value.QuotedStatus.ExtendedEntities.Media
					for _, entity := range entities {
						variants := entity.VideoInfo.Variants
						sendSms(variants)
					}
				}
			}
		}
	}
}
