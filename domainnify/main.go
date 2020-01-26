package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprint("%v", *s)
}

func (s *stringSlice) Set(v string) error {
	*s = append(*s, v)
	return nil
}

const allowedChars = "abcdefghijklmnopqrstuvwxyz01234567890_-"

func main() {
	// tオプションで複数のTLDを指定できるようにする。
	// 例: go run main.go -t com -t net
	var tlds stringSlice
	flag.Var(&tlds, "t", "add top level domain")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
