package shaw_test

import (
	"io"
	"strings"
	"testing"

	"deedles.dev/shaw"
	"deedles.dev/shaw/dict/dicts"
)

func TestTranslate(t *testing.T) {
	const before = "This is a test."
	afterbytes, err := io.ReadAll(shaw.Translate(strings.NewReader(before), dicts.American()))
	if err != nil {
		t.Fatal(err)
	}
	after := string(afterbytes)
	if after != "𐑞𐑦𐑕 𐑦𐑟 𐑩 𐑑𐑧𐑕𐑑." {
		t.Fatalf("%q", after)
	}
}
