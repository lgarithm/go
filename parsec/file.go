package parsec

import "os"

func LexFile(filename string, lex ...Tokenizer) (TokenStream, string, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	s, t := Lex(string(bs), lex...)
	return RemoveBlankTokens(s), t, nil
}

func ParseFile[T any](filename string, tokenizers []Tokenizer, p Parse[T]) (*T, error) {
	ts, _, err := LexFile(filename, tokenizers...)
	if err != nil {
		return nil, err
	}
	x, _ := TillEnd(p)(ts)
	if x == nil {
		return nil, ErrFailed
	}
	return x, nil
}

func MakeFileParser[T any](tokenizers []Tokenizer, p Parse[T]) func(string) (*T, error) {
	return func(filename string) (*T, error) {
		return ParseFile(filename, tokenizers, p)
	}
}
