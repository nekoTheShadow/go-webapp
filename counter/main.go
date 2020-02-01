package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
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
}
