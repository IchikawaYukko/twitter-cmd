# twitter-cmd [English](README.md)
コマンドラインからツイートできる何か！ シェルスクリプト用に便利！？

## ビルド
1. ビルド環境の立ち上げ

    `docker-compose run build-env`

2. ビルド

    `make deps`

    `make` または `make build-windows`

## 使い方
1. 以下の OAuth トークンを環境変数に準備する

    `TWITTER_CONSUMER_KEY=`

    `TWITTER_CONSUMER_SECRET=`

    `TWITTER_ACCESS_TOKEN=`

    `TWITTER_ACCESS_TOKEN_SECRET=`

2. つぶやく!!

    Linux: `./twitter ゎーぉ！`

    Windows: `twitter.exe ゎーぉ！`

3. 画像付きツイート
    `./twitter -m ファイル名1 ゎーぉ！`
    `./twitter -m ファイル名1 -m ファイル名2 ゎーぉ！`
    (ファイルは4つまで指定できるよ)
