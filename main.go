// Newline format must be LF
package main

import (
	. "fmt"
	"github.com/ChimeraCoder/anaconda"
	. "os"
)

func main() {
	api := GetTwitterApi()

	if len(Args) < 2 {
		Println("Usage: twitter status_text\n\nPlease specify status and Environment variable shown below.\nTWITTER_CONSUMER_KEY\nTWITTER_CONSUMER_SECRET\nTWITTER_ACCESS_TOKEN\nTWITTER_ACCESS_TOKEN_SECRET")
		Exit(255)
	}

	tweet, err := api.PostTweet(parse_message(Args), nil)
	if err != nil {
		panic(err)
	}

	Print(tweet.Text)
}

func parse_message(args []string) string {
	var text string = ""
	for i, v := range args {
		if i == 0 {
			continue	// skip first arg value (It is name of executable)
		}
		text += " " + v
	}
	return text
}

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
