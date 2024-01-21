package parsec

import (
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Split[T any] func(T) (T, T)

type Tokenizer = Split[string]

func TakeWhile(p RunePred, text string) (string, string) {
	for i, c := range text {
		if !p(c) {
			return text[:i], text[i:]
		}
	}
	return text, ""
}

func Ident0(p RunePred) Tokenizer {
	return func(text string) (string, string) { return TakeWhile(p, text) }
}

var (
	Spaces   Tokenizer = Ident0(unicode.IsSpace)
	NoSpaces Tokenizer = Ident0(Not(unicode.IsSpace))
	Digits   Tokenizer = Ident0(unicode.IsDigit)
)

func tokenize(text string, ts ...Tokenizer) (string, string) {
	for _, f := range ts {
		if t0, t1 := f(text); len(t0) > 0 {
			return t0, t1
		}
	}
	return "", text
}

func Symbols(s []string) Tokenizer {
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) > len(s[j]) || len(s[i]) == len(s[j]) && s[i] < s[j]
	})
	return SymbolSet(s).Tokenize
}

func Ident(p0, p1 RunePred) Tokenizer {
	return func(text string) (string, string) {
		if len(text) == 0 || !p0([]rune(text)[0]) {
			return "", text
		}
		return TakeWhile(p1, text)
	}
}

func NoneBlank(t string) bool {
	for _, c := range t {
		if !unicode.IsSpace(c) {
			return true
		}
	}
	return false
}

func QuotedStr(text string) (string, string) {
	if len(text) == 0 || []rune(text)[0] != '"' {
		return "", text
	}
	q, err := strconv.QuotedPrefix(text)
	if err != nil {
		return "", text
	}
	if !strings.HasPrefix(text, q) {
		return "", text
	}
	return q, text[len(q):]
}
