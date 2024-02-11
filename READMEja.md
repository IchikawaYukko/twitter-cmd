[English](README.md)
# twitter-cmd
コマンドライン簡易ツイッタークライアント！

コマンドラインからツイートできる何か！ シェルスクリプト用に便利！？

## ビルド
1. ビルド環境の立ち上げ

    `docker-compose run build-env`

1. CONSUMER_KEY と CONSUMER_SECRET の設定

    1. Twitter Developpers > Apps ページがら、自分用の key/secret を取得
    1. 取得した key/secret を consumer_token.go に書き込む

1. 依存パッケージを取得

    `make deps`

1. ビルド

    `make` または `make build-windows`

[releases](https://github.com/IchikawaYukko/twitter-cmd/releases) からビルド済みバイナリをダウンロードしてもいい。

## 使い方 (OAuth トークンの取得)

1. OAuth 認証を開始

    `./twitter 1`

    実行すると、以下のような認証URLが表示されます。

```
Open this URL and login with your Twitter account.
https://api.twitter.com/oauth/authenticate?oauth_token=xxxxxxxxxx
```
2. 表示された URL をブラウザで開く

3. 認証後に表示された PIN を入力

```
Enter PIN: *******
Token saved to .ichikawayukko-twitter_cmd
Erase this file to re-authenticate.
```

## 使い方 (ツイート)
1. つぶやく!!

    Linux: `./twitter ゎーぉ！`

    Windows: `twitter.exe ゎーぉ！`

1. 画像付きツイート

    `./twitter -m ファイル名1 ゎーぉ！`

    `./twitter -m ファイル名1 -m ファイル名2 ゎーぉ！`

    (ファイルは4つまで指定できるよ)
