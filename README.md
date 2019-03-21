# go-playground

## build
docker-compose build

## run
docker-compose up -d
docker exec -it goplayground_app_1 go run study.go
docker exec -it goplayground_app_1 go test packageName

## キーワード
- type キーワードによる型宣言
- strutキーワードを使った構造体
- 関数のレシーバ型
- newを利用したメモリ確保
- ...型名　での可変長引数

## ダックタイピング
インタフェースが必要とするメソッドをすべて実装した時点で、自動的にそのインタフェースを実装したとみなされます。


## 文法
- := は関数ないに限り使用可能
- deferによる遅延実行

## 演算子
http://cuto.unirita.co.jp/gostudy/post/go-operator/
### アドレス演算子
- & 右辺の変数のアドレスを得る
- * 右辺の変数に格納されたアドレスを解決する

## ユニットテスト
- assertは無い
- *_test.go というファイル名


### 送受信演算子
- <- 左辺のチェネルへ右辺の値を送信する
- =<- 右辺のチェネルから値を受信する

## パッケージ
- 名前空間を分けるためにpackageを利用する
- importで使用する
- 大文字で始まるものはimportで参照可能なパブリックメンバ
- 小文字で始まるものはimportでも参照不可能なプライベートメンバ

## TODO
- Study testing framework
- create cli sample

## reference
http://cuto.unirita.co.jp/gostudy/