package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const otherWord = "*"

func main() {
	// sprinkle.exe のフルパスを取得する。
	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// sprinkle.exe と同じディレクトリにある otherword.txt を読み込む。
	fp, err := os.Open(filepath.Join(filepath.Dir(exe), "otherword.txt"))
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := bufio.NewScanner(fp)

	taransforms := []string{}
	for reader.Scan() {
		taransforms = append(taransforms, reader.Text())
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := taransforms[rand.Intn(len(taransforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
