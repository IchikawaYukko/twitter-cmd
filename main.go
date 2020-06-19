// Newline format must be LF
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func main() {
	api := GetTwitterApi()

	if len(os.Args) < 2 {
		fmt.Println(`
Usage:
twitter status_text
twitter -m media_filename1 status_text
twitter -m media_filename1 -m media_filename2 status_text    (Up to 4 media)
twitter "-m foobar"  (if status_text contains -m, must quote it.)

`)
		os.Exit(255)
	}

	message, files := parse_message(os.Args)
	var (
		attachment url.Values = nil
		media      []string
	)

	if 4 < len(files) {
		fmt.Println(os.Stderr, "Only 4 images can be posted at one time.")
		os.Exit(-1)
	} else if 0 < len(files) {
		for _, file := range files {
			media = append(media, upload_image(api, file))
		}

		attachment = url.Values{}
		attachment.Add("media_ids", strings.Join(media, ","))
	}

	tweet, err := api.PostTweet(message, attachment)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}

	fmt.Print(tweet.Text)
}

func upload_image(api *anaconda.TwitterApi, filename string) string {
	media, err := api.UploadMedia(load_image(filename))
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}

	return media.MediaIDString
}

func load_image(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}
	return base64.StdEncoding.EncodeToString(data)
}

func parse_message(args []string) (string, []string) {
	var (
		text           string = ""
		image_filename []string
		image_flag     bool
	)

	for i, v := range args {
		if i == 0 {
			continue // skip first arg value (It is name of executable)
		}
		if v == "-m" { // if -m flag found, next args is image filename
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
	anaconda.SetConsumerKey(consumer_token)
	anaconda.SetConsumerSecret(consumer_secret)

	// If token is missing, get it
	_, err := os.Stat(token_savefile)
	if err != nil {
		save_token(get_token())
	}

	access_token, access_secret := load_token()
	api := anaconda.NewTwitterApi(access_token, access_secret)
	return api
}
