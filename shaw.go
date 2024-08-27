package shaw

import (
	"bufio"
	"errors"
	"io"
	"iter"
	"os"

	"deedles.dev/shaw/dict"
	"deedles.dev/xiter"
)

// isTranslatable returns true if r is something expected in a
// translatable word.
func isTranslatable(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '\''
}

// runes returns an iterator over runes read from r.
func runes(r io.Reader) iter.Seq2[rune, error] {
	return func(yield func(rune, error) bool) {
		br := bufio.NewReader(r)
		for {
			c, _, err := br.ReadRune()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					yield(0, err)
				}
				return
			}

			if !yield(c, nil) {
				return
			}
		}
	}
}

// words returns an iterator over word segments. Despite the name, it
// returns whitespace and other non-word things as well as continuous
// runs between each. In other words,
//
//	an  example
//
// will yield
//
//	[]rune("an")
//	[]rune("  ")
//	[]rune("example")
func words(runes iter.Seq[rune]) iter.Seq[[]rune] {
	return xiter.ChunksFunc(runes, isTranslatable)
}

// translateLines translates words read from r in a line-by-line
// manner using a [bufio.Scanner]. It is primarily for translating
// from stdin. Like [words], it also yields intervening non-word
// segments.
func translateLines(r io.Reader, dict *dict.Dict) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		s := bufio.NewScanner(r)
		for s.Scan() {
			for word := range words(xiter.Runes(s.Bytes())) {
				s := string(word)
				if isTranslatable(word[0]) {
					s = dict.Translate(s)
				}
				if !yield(s, nil) {
					return
				}
			}
			if !yield("\n", nil) {
				return
			}
		}
		if err := s.Err(); err != nil {
			yield("", err)
		}
	}
}

// translateRunes translates words read from r. Like [words], it also
// yields intervening non-word segments.
func translateRunes(r io.Reader, dict *dict.Dict) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		runes := func(ryield func(rune) bool) {
			for r, err := range runes(r) {
				if err != nil {
					yield("", err)
					return
				}

				if !ryield(r) {
					return
				}
			}
		}

		for w := range words(runes) {
			s := string(w)
			if isTranslatable(w[0]) {
				s = dict.Translate(s)
			}
			if !yield(s, nil) {
				return
			}
		}
	}
}

// Translate returns an iterator that translates words read from r
// using dict. It also yields intervening non-word segments in between
// the words translated such that joining all of the yielded strings
// in order will result in a fully translated piece of text with
// whitespace and punctuation preserved.
//
// If at any point a non-nil error is yielded, it will be the last
// thing yielded.
func Translate(r io.Reader, dict *dict.Dict) iter.Seq2[string, error] {
	if r == os.Stdin {
		return translateLines(r, dict)
	}
	return translateRunes(r, dict)
}
