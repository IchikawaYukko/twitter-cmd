// Newline format must be LF
package main

import (
	"flag"
	. "fmt"
	"github.com/ChimeraCoder/anaconda"
	. "os"
)

func main() {
	api := GetTwitterApi()

	flag.Parse()
	if len(Args) == 1 {
		panic("Usage: twitter status_text\n\nPlease specify status.")
	}
	text := flag.Arg(0)
	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}

	Print(tweet.Text)
}

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
