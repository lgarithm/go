package parsec

import "strings"

type SymbolSet []string

func (s SymbolSet) Tokenize(text string) (string, string) {
	for _, p := range s {
		if strings.HasPrefix(text, p) {
			return p, text[len(p):]
		}
	}
	return "", text
}

func (ss SymbolSet) Others(s TokenStream) (*string, TokenStream) {
	if len(s) == 0 {
		return nil, s
	}
	for _, x := range ss {
		if x == s[0] {
			return nil, s
		}
	}
	return &s[0], s[1:]
}
