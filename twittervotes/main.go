package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Session

type poll struct {
	Options []string
}

func dialdb() error {
	var err error
	log.Println("MongoDBにダイアル中:localhost")
	db, err = mgo.Dial("localhost")
	return err
}

func closedb() {
	db.Close()
	log.Panicln("データベース接続が閉じられました")
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("ballots").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}
