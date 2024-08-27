package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"deedles.dev/shaw"
	"deedles.dev/shaw/dict"
)

func loadDict(path string) (*dict.Dict, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return dict.Parse(file)
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

	for w, err := range shaw.Translate(os.Stdin, dict) {
		if err != nil {
			slog.Error("read input", "err", err)
			os.Exit(1)
		}

		fmt.Print(w)
	}
}
