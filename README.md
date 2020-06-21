[日本語](READMEja.md)
# twitter-cmd
Simple CLI Twitter client.

Post something to Twitter from command line :-) May be useful in shell script use ?

## Build
1. Launch build environment

    `docker-compose run build-env`

1. Set OAuth CONSUMER_KEY and CONSUMER_SECRET

    1. Get your key/secret from Twitter Developpers > Apps page.
    1. Set it to consumer_token.go

1. Build

    `make deps`

    `make` or `make build-windows`

Or you can download pre compiled executable from [releases](https://github.com/IchikawaYukko/twitter-cmd/releases).

## Usage (Get OAuth token)

1. Start OAuth Authentication

    `./twitter 1`

    This will show Authenticate URL like this

```
Open this URL and login with your Twitter account.
https://api.twitter.com/oauth/authenticate?oauth_token=xxxxxxxxxx
```
2. Open Authenticate URL by browser.

3. Enter PIN

```
Enter PIN: *******
Token saved to .ichikawayukko-twitter_cmd
Erase this file to re-authenticate.
```

# Usage (Tweet)
1. Tweet !!

    Linux: `./twitter Hello World!`

    Windows: `twitter.exe Hello World!`

1. Tweet with image

    `./twitter -m filename1 Hello World?`

    `./twitter -m filename1 -m filename?2 Hello World?`

    (4 files can be specified at one time)
