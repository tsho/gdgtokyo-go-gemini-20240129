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

## 各フォルダの分け方
勝手ながら下記の4つに分けました。

- 0 [この勉強会で使う共通モジュールの作成](https://www.youtube.com/watch?v=3ZomyM0expM&t=20m46s)
- 1 [モデルの初期化](https://www.youtube.com/watch?v=3ZomyM0expM&t=25m40s)
- 2 [テキストのみの入力からテキストを生成する](https://www.youtube.com/watch?v=3ZomyM0expM&t=36m35s)
- 3 [テキストと画像の入力からテキスト生成(マルチモーダル)](https://www.youtube.com/watch?v=3ZomyM0expM&t=45m16s)



## 各フォルダの初回設定
1. go mod tidyの実行
```
go mod tidy
```

2. 実行
```
go run .
```

* 3のmutlmodalは画像をこちらのgithubからダウンロードする
[hi.pngとangry.png](https://github.com/tenntenn/gopher-stickers/tree/master/png)


### もし自分で新しいフォルダを作った場合

modの初期化
```
go mod init example.com/$(basename `pwd`)
```

## 01_initialize_modelの出力結果例
```
&{0xc000305ba0 models/gemini-pro {<nil> [] <nil> <nil> <nil> <nil>} []}
```

## 02_generate_textの出力結果例
```
Goのマスコットキャラクターは「Gopher(ゴファー)」です。Gopherは、穴を掘って暮らす哺乳類の一種、....
```
* 生成AIは同じ結果がでることを保証しないので出力結果は異なる可能性があります。

## 03_multimodal_text_imagesの出力結果例
```
2つの画像の違いは、左の画像のGopherが笑顔で手を振っているのに対し、右の画像のGopherは怒った顔で何かを訴えていることです。共通点は、どちらもGopherというキャラクターであることです。
```
* 生成AIは同じ結果がでることを保証しないので出力結果は異なる可能性があります。


## 参考
GoのSDKのドキュメント[リンク](https://pkg.go.dev/github.com/google/generative-ai-go#section-readme)


## エラー対応
```
... generativelanguagepb/citation.pb.go:27:2: missing go.sum entry for module providing package google.golang.org/genproto/googleapis/api/annotations (imported by cloud.google.com/go/ai/generativelanguage/apiv1/generativelanguagepb); to add:
```
上記のようなエラーメッセージが出た場合は下記の `go mod tidy` を実行する
