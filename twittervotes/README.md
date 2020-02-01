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

# コンテナの作成
docker run                      ^
  -d                            ^
  --name go-webapp-nsqd         ^
  -p 4150:4150 -p 4151:4151     ^
  nsqio/nsq /nsqd               ^
  --broadcast-address=localhost ^
  --lookupd-tcp-address=localhost:4160
```

この手順で実施すると、`nsq_tail`が正常に動作しない。そこで以下のURLにアクセスすることで`nsqd`が想定通り動作していることを確認する: `http://192.168.99.100:4151/stats`