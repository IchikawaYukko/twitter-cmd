package main

import (
	"bufio"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"os"
)

var token_savefile string = ".ichikawayukko-twitter_cmd"

func get_token() (string, string) {
	api := anaconda.NewTwitterApi("", "")

	uri, cred, err := api.AuthorizationURL("oob")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}
	fmt.Println("Open this URL and login with your Twitter account.\n" + uri)

	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter PIN: ")
	stdin.Scan()
	pin := stdin.Text()

	cred, _, err = api.GetCredentials(cred, pin)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}

	return cred.Token, cred.Secret
}

func save_token(access_token string, access_secret string) {
	file, err := os.OpenFile(token_savefile, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}
	defer file.Close()

	fmt.Fprintln(file, access_token)
	fmt.Fprintln(file, access_secret)
	fmt.Fprintln(file, "")

	fmt.Println("Token saved to " + token_savefile)
	fmt.Println("Erase this file to re-authenticate.")
}

func load_token() (string, string) {
	file, err := os.OpenFile(token_savefile, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(-1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	access_token, _, _ := reader.ReadLine()
	access_secret, _, _ := reader.ReadLine()

	return string(access_token), string(access_secret)
}
