package dicts_test

import (
	"testing"

	"deedles.dev/shaw/dict/dicts"
)

func TestAmerican(t *testing.T) {
	if dicts.American() == nil {
		t.Fatal("failed to load")
	}
}
