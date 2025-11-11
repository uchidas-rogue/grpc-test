# grpc-test

## install

### Goをインストール
homepageからインストール

### protocol bufferのインストール

```zsh
# protbuf本体
brew install protobuf

# go用のツールinstall
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# path追加
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> $HOME/.zshrc
source ~/.zshrc
```

vscodeへ拡張インストール
* Go (by Go Team at Google)
* Tooltitude for Protobuf (proto files)


### サーバーセットアップ

credsフォルダへ移動して、期間長めで適当な自己証明書を作成する
```
cd creds
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 36500 -nodes -subj "/CN=localhost"
```

サーバー起動
```
make run-orders
```
