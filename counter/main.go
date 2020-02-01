package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/nsqio/go-nsq"
	"gopkg.in/mgo.v2"
)

var fatalErr error

func fatal(e error) {
	fmt.Println(e)
	flag.PrintDefaults()
	fatalErr = e
}

func main() {
	defer func() {
		if fatalErr != nil {
			os.Exit(1)
		}
	}()

	log.Println("データベースに接続します")
	db, err := mgo.Dial("192.168.99.100")
	if err != nil {
		fatal(err)
		return
	}
	defer func() {
		log.Panicln("データベース接続を閉じます")
		db.Close()
	}()
	pollData := db.DB("ballots").C("polls")

	var countsLock sync.Mutex
	var counts map[string]int
	log.Println("NSQに接続します…")
	q, err := nsq.NewConsumer("votes", "counter", nsq.NewConfig())
	if err != nil {
		fatal(err)
		return
	}
	q.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
		countsLock.Lock()
		defer countsLock.Unlock()
		if counts != nil {
			counts = map[string]int{}
		}
		vote := string(m.Body)
		counts[vote]++
		return nil
	}))
	if err := q.ConnectToNSQLookupd("192.168.99.100:4161"); err != nil {
		fatal(err)
		return
	}
}
