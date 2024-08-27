package dict

import (
	"bufio"
	"io"
	"iter"
	"strings"
)

type Dict struct {
	prefixes map[string]string
	suffixes map[string]string
	words    map[string]string
}

func Parse(r io.Reader) (*Dict, error) {
	dict := Dict{
		prefixes: make(map[string]string),
		suffixes: make(map[string]string),
		words:    make(map[string]string),
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) != 2 {
			continue
		}

		// TODO: Find out what these are actually supposed to do.
		fields[0] = strings.TrimSuffix(fields[0], "_")
		fields[1] = strings.TrimSuffix(fields[1], ":")

		switch fields[0][0] {
		case '^':
			dict.prefixes[fields[0][1:]] = fields[1]
		case '$':
			dict.suffixes[fields[0][1:]] = fields[1]
		default:
			dict.words[fields[0]] = fields[1]
		}
	}

	return &dict, s.Err()
}

func (d *Dict) Translate(word string) string {
	if w, ok := d.words[word]; ok {
		return w
	}

	for p, s := range ixes(word) {
		tp, ok := d.prefixes[p]
		if !ok {
			continue
		}
		ts, ok := d.suffixes[s]
		if !ok {
			continue
		}

		return tp + ts
	}

	if lower := strings.ToLower(word); lower != word {
		return d.Translate(lower)
	}

	return word
}

// ixes returns an iterator over all possible combinations of prefixes
// and suffixes that form the given word.
func ixes(word string) iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for i := 1; i < len(word)-1; i++ {
			if !yield(word[:i], word[i:]) {
				return
			}
		}
	}
}
