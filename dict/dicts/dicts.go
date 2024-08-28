package dicts

import (
	"compress/gzip"
	"embed"
	"sync"

	"deedles.dev/shaw/dict"
)

//go:embed *.dict.gz
var fs embed.FS

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

var american = sync.OnceValue(func() *dict.Dict {
	file := must(fs.Open("amer.dict.gz"))
	r := must(gzip.NewReader(file))
	defer r.Close()
	return must(dict.Parse(r))
})

func American() *dict.Dict {
	return american()
}
