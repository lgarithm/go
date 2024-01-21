package parsec

type Language struct {
	Tokenizers []Tokenizer
}

func (l Language) LexFile(filename string) (TokenStream, string, error) {
	return LexFile(filename, l.Tokenizers...)
}
