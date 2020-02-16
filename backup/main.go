package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/matryer/filedb"
)

type path struct {
	Path string
	Hash string
}

func (p path) String() string {
	return fmt.Sprintf("%s [%s]", p.Path, p.Hash)
}

func main() {
	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	dbpath := flag.String("db", "./backupdata", "データベースのディレクトリへのパス")
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fatalErr = errors.New("エラー: コマンドを指定してください")
		return
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

	switch strings.ToLower(args[0]) {
	case "list":
		var path path
		col.ForEach(func(i int, data []byte) bool {
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return true
			}
			fmt.Printf("= %s\n", path)
			return false
		})
	case "add":
		if len(args[1:]) == 0 {
			fatalErr = errors.New("追加するパスを指定してください")
			return
		}
		for _, p := range args[1:] {
			path := &path{Path: p, Hash: "まだアーカイブされていません"}
			if err := col.InsertJSON(path); err != nil {
				fatalErr = err
				return
			}
			fmt.Printf("+ %s\n", path)
		}
	case "remove":
		// https://github.com/matryer/filedb の 実装が誤っているため、Windowsでは正しく動作しない。
		// 上記コードを参考にして独自実装しておく

		// 残すものを配列pathsにピックアップ
		paths := []path{}
		col.ForEach(func(i int, data []byte) bool {
			var path path
			err := json.Unmarshal(data, &path)
			if err != nil {
				fatalErr = err
				return true
			}

			remain := true
			for _, p := range args[1:] {
				if path.Path == p {
					remain = false
					break
				}
			}

			if remain {
				paths = append(paths, path)
			}
			return false
		})

		// 全件削除
		col.RemoveEach(func(i int, data []byte) (bool, bool) { return true, false })

		// 先の処理で残したものを再書き込み
		for _, path := range paths {
			if err := col.InsertJSON(path); err != nil {
				fatalErr = err
				return
			}
		}
	}
}
