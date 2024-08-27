// Package dict handles parsing and using dictionary files.
//
// The dictionary file format, as explained in a comment in Dave Coffin's Python translator, is as follows:
//
// Each line of a dictionary consists of an English word, a space,
// a Shavian translation, and no comments. Special notations are:
//
//	^word ð‘¢ð‘»ð‘›	word is a prefix
//	$word ð‘¢ð‘»ð‘›	word is a suffix
//	Word ð‘¢ð‘»ð‘›	always use a namer dot
//	word_ ð‘¢ð‘»ð‘›	never use a namer dot
//	word_VB ð‘¢ð‘»ð‘›	shave this way when tagged as a verb
//	word. ð‘¢ð‘»ð‘›	shave this way when no suffix is present
//	word .ð‘¢ð‘»ð‘›	word takes no prefixes
//	word ð‘¢ð‘»ð‘›.	word takes no suffixes
//	word ð‘¢ð‘»ð‘›:	suffixes do not alter the root,
//	               	e.g. "ð‘‘ð‘¾" or "ð‘•ð‘¾" palatizing to "ð‘–ð‘©" or "ð‘ ð‘©".
//	word .		delete this word from the dictionary
//
// Words are matched case-sensitive when possible, e.g. US/us,
// WHO/who, Job/job, Nice/nice, Polish/polish.
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
	words    map[string]word
}

type word struct {
	v      string
	prefix bool
	suffix bool
}

func Parse(r io.Reader) (*Dict, error) {
	dict := Dict{
		prefixes: make(map[string]string),
		suffixes: make(map[string]string),
		words:    make(map[string]word),
	}

	s := bufio.NewScanner(r)
	for s.Scan() {
		fields := strings.Fields(s.Text())
		if len(fields) != 2 {
			continue
		}

		// TODO: Handle these.
		fields[0] = strings.TrimSuffix(fields[0], "_")
		fields[0] = strings.TrimSuffix(fields[0], "_VB")
		fields[1] = strings.TrimSuffix(fields[1], ":")

		var noprefix, nosuffix bool
		fields[1], noprefix = strings.CutPrefix(fields[1], ".")
		fields[1], nosuffix = strings.CutSuffix(fields[1], ".")

		switch fields[0][0] {
		case '^':
			dict.prefixes[fields[0][1:]] = fields[1]
		case '$':
			dict.suffixes[fields[0][1:]] = fields[1]
		default:
			dict.words[fields[0]] = word{v: fields[1], prefix: !noprefix, suffix: !nosuffix}
		}
	}

	return &dict, s.Err()
}

func (d *Dict) Translate(word string) string {
	if w, ok := d.words[word]; ok {
		return w.v
	}

	for p, s := range ixes(word) {
		if tp, ok := d.prefixes[p]; ok {
			if ts, ok := d.words[s]; ok && ts.suffix {
				return tp + ts.v
			}
		}

		if ts, ok := d.suffixes[s]; ok {
			if tp, ok := d.words[p]; ok && tp.prefix {
				return tp.v + ts
			}
		}
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
