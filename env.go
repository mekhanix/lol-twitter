package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadEnv returns map of strings from .env
func LoadEnv() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return map[string]string{
		"TWITTER_CONSUMER_KEY":        os.Getenv("TWITTER_CONSUMER_KEY"),
		"TWITTER_SECRET_CONSUMER_KEY": os.Getenv("TWITTER_SECRET_CONSUMER_KEY"),
		"TWITTER_ACCESS_TOKEN":        os.Getenv("TWITTER_ACCESS_TOKEN"),
		"TWITTER_SECRET_ACCESS_TOKEN": os.Getenv("TWITTER_SECRET_ACCESS_TOKEN"),
		"TWITTER_ID":                  os.Getenv("TWITTER_ID"),
		"TWILIO_ACCOUNT_SID":          os.Getenv("TWILIO_ACCOUNT_SID"),
		"TWILIO_AUTH_TOKEN":           os.Getenv("TWILIO_AUTH_TOKEN"),
		"SMS_FROM":                    os.Getenv("SMS_FROM"),
		"SMS_TO":                      os.Getenv("SMS_TO"),
	}
}
