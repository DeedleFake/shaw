package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"deedles.dev/shaw/dict"
	"deedles.dev/xiter"
)

func loadDict(path string) (*dict.Dict, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return dict.Parse(file)
}

func isTranslatable(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '\''
}

func main() {
	dictfile := flag.String("dict", "", "dictionary file")
	flag.Parse()
	if *dictfile == "" {
		flag.Usage()
		os.Exit(2)
	}

	dict, err := loadDict(*dictfile)
	if err != nil {
		slog.Error("load dictionary", "err", err)
		os.Exit(1)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		words := xiter.ChunksFunc(xiter.Runes(s.Bytes()), isTranslatable)
		for word := range words {
			s := string(word)
			if isTranslatable(word[0]) {
				s = dict.Translate(s)
			}
			fmt.Print(s)
		}
		fmt.Println()
	}
	if err := s.Err(); err != nil {
		slog.Error("read input", "err", err)
		os.Exit(1)
	}
}
