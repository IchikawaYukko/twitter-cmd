// Newline format must be LF
package main

import (
	"encoding/base64"
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

	message, files := parse_message(Args)
	
	base64.StdEncoding.EncodeToString(data)
	Println(files)

	tweet, err := api.PostTweet(message, nil)
	if err != nil {
		panic(err)
	}

	Print(tweet.Text)
}

func parse_message(args []string) (string, []string) {
	var text string = ""
	var image_filename []string
	var image_flag bool

	for i, v := range args {
		if i == 0 {
			continue	// skip first arg value (It is name of executable)
		}
		if v == "-d" {	// if -d flag found, next args is image filename
			image_flag = true
			continue
		}
		if image_flag == true {
			image_filename = append(image_filename, v)
			image_flag = false
		} else {
			text += " " + v
		}
	}
	return text, image_filename
}

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(Getenv("TWITTER_ACCESS_TOKEN"), Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
