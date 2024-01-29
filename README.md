2024/1/29に開催されたGDG Tokyoのイベントのライブコーディングされたコードを自分のまとめ用に作った非公式のrepositoryです。

## 勉強会URL
- https://gdg-tokyo.connpass.com/event/307689/

## 動作環境
- go1.21.6

24/1/29現在、Go 1.20以降が必須 [参考](https://ai.google.dev/tutorials/go_quickstart?hl=ja#prerequisites)。

## 本レポジトリの初回設定
1. APIキーの取得方法 (本レポジトリの01以降のスクリプトで必要)
下記のページを参考。
[APIキーを設定する](https://ai.google.dev/tutorials/go_quickstart?hl=ja#set-up-api-key)

2. APIキーの登録（本レポジトリではローカルの環境設定を使う）
端末で下記を実行する
```
export $GEMINI_APIKEY="ここに自分のAPIキーを書く"
```

3. SDKパッケージをインストールする
```
go get github.com/google/generative-ai-go
```
[引用元](https://ai.google.dev/tutorials/go_quickstart?hl=ja#add-sdk)


## 各フォルダの初回設定
1. modの初期化
```
go mod init example.com/$(basename `pwd`)
```

2. go mod tidyの実行
```
go mod tidy
```

## 01_initialize_modelの出力結果例
```
&{0xc000305ba0 models/gemini-pro {<nil> [] <nil> <nil> <nil> <nil>} []}
```

## 参考
GoのSDKのドキュメント[リンク](https://pkg.go.dev/github.com/google/generative-ai-go#section-readme)


## エラー対応
```
... generativelanguagepb/citation.pb.go:27:2: missing go.sum entry for module providing package google.golang.org/genproto/googleapis/api/annotations (imported by cloud.google.com/go/ai/generativelanguage/apiv1/generativelanguagepb); to add:
```
上記のようなエラーメッセージが出た場合は下記の `go mod tidy` を実行する
