package main

import (
	"backup/backup"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			log.Fatalln(fatalErr)
		}
	}()

	interval := flag.Int("interval", 10, "チェックの間隔(秒単位)")
	archive := flag.String("archive", "archive", "アーカイブの保存先")
	dbpath := flag.String("db", "./db", "filedbへのパス")
	flag.Parse()

	m := &backup.Monitor{
		Destination: *archive,
		Archiver:    backup.ZIP,
		Paths:       make(map[string]string),
	}

	db, err := filedb.Dial(*dbpath)
	if err != nil {
		fatalErr = err
		return
	}
	defer db.Close()

	col, err := db.C("paths")
	if err != nil {
		fatalErr = err
		return
	}

	var path path
	col.ForEach(func(_ int, data []byte) bool {
		if err := json.Unmarshal(data, &path); err != nil {
			fatalErr = err
			return true
		}
		m.Paths[path.Path] = path.Hash
		return false
	})

	if fatalErr != nil {
		return
	}

	if len(m.Paths) < 1 {
		fatalErr = errors.New("パスがありません。backupツールを使って追加してください")
		return
	}

	check(m, col)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
Loop:
	for {
		select {
		case <-time.After(time.Duration(*interval) * time.Second):
			check(m, col)
		case <-signalChan:
			fmt.Println()
			log.Printf("終了します")
			break Loop
		}
	}
}

func check(m *backup.Monitor, col *filedb.C) {
	log.Println("チェックします…")
	counter, err := m.Now()
	if err != nil {
		log.Panicln("バックアップに失敗しました")
	}

	if counter == 0 {
		log.Println("変更はありません")
		return
	}

	log.Printf("%dこのディレクトリをアーカイブしました\n", counter)
	var path path
	col.SelectEach(func(_ int, data []byte) (bool, []byte, bool) {
		if err := json.Unmarshal(data, &path); err != nil {
			log.Println("JSONデータの書き出しに失敗しました。次の項目に進みます: ", err)
			return true, data, false
		}

		path.Hash, _ = m.Paths[path.Path]
		newdata, err := json.Marshal(&path)
		if err != nil {
			log.Println("JSONデータの書き出しに失敗しました。次の項目に進みます: ", err)
			return true, data, false
		}
		return true, newdata, false
	})

}
