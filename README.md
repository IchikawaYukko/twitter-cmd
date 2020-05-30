# twitter-cmd
Post something to Twitter from command line :-) May be useful in shell script use ?

## Build
1. Launch build environment

    `docker-compose run build-env`

2. Build

    `make deps`

    `make` or `make build-windows`

## Usage
1. Prepare OAuth token as environment variables below.

    `TWITTER_CONSUMER_KEY=`

    `TWITTER_CONSUMER_SECRET=`

    `TWITTER_ACCESS_TOKEN=`

    `TWITTER_ACCESS_TOKEN_SECRET=`

2. Tweet !!

    `./twitter Hello World!`
