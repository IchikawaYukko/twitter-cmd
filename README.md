[日本語](READMEja.md)
# twitter-cmd
Simple CLI Twitter client.

Post something to Twitter from command line :-) May be useful in shell script use ?

## Build
1. Launch build environment

    `docker-compose run build-env`

2. Build

    `make deps`

    `make` or `make build-windows`

Or you can download pre compiled executable from [releases](https://github.com/IchikawaYukko/twitter-cmd/releases).

## Usage
1. Prepare OAuth token as environment variables below.

    `TWITTER_CONSUMER_KEY=`

    `TWITTER_CONSUMER_SECRET=`

    `TWITTER_ACCESS_TOKEN=`

    `TWITTER_ACCESS_TOKEN_SECRET=`

2. Tweet !!

    Linux: `./twitter Hello World!`

    Windows: `twitter.exe Hello World!`

3. Tweet with image

    `./twitter -m filename1 Hello World?`

    `./twitter -m filename1 -m filename?2 Hello World?`

    (4 files can be specified at one time)
