# 前提

本書ではnsqとmongoの構築にあたっては、ローカル環境へのインストールを行っているが、今回はDockerを利用した。

- Windows 10 Home
- Docker ToolBox

なおDockerのホストIPは次の通りである: `192.168.99.100`

# mongo

mongoについては、以下の手順でベースとなるソフトウェアのインストールを行う。

```bash
# イメージのダウンロード
docker pull mongo

# コンテナの作成
docker run                               ^
  -d                                     ^
  --name go-webapp-mongo                 ^
  -p 27017:27017                         ^
  -e MONGO_INITDB_ROOT_USERNAME=dev      ^
  -e MONGO_INITDB_ROOT_PASSWORD=password ^
  mongo

# コンテナへのアタッチ
docker exec -it go-webapp-mongo bash
```

その後、本書の手順にある通り、DBとコレクションを作成&初期化する。

```bash
# mongoの起動
mongo -u "dev" -p "password"

# DBの作成
use ballots

# コレクションの作成&初期化
db.polls.insert({
  "title" : "今の気分は?",
  "options" : [
    "happy",
    "sad",
    "fail",
    "win"
]})
```

# nsq

インストールは次の通り

```bash
# イメージのダウンロード
docker pull nsqio/nsq

# コンテナの作成 (nsqlookup)
docker run -d                 ^
  --name go-webapp-nsqlookupd ^
  -p 4160:4160 -p 4161:4161   ^
  nsqio/nsq /nsqlookupd

# コンテナの作成 (nsqd)
docker run -d                               ^
  --name go-webapp-nsqd                     ^
  -p 4150:4150 -p 4151:4151                 ^
  nsqio/nsq /nsqd                           ^
  --broadcast-address=192.168.99.100        ^
  --lookupd-tcp-address=192.168.99.100:4160
```

この手順の場合、`nsq_tail`は以下の手順で利用すること:

```
docker run --rm nsqio/nsq /nsq_tail --lookupd-http-address=192.168.99.100:4161 -topic votes
```
