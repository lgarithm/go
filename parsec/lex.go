package parsec

type Stream[T any] []T

type TokenStream = Stream[string]

type Lexer = func(string) []string

func Lex(text string, ls ...Tokenizer) (TokenStream, string) {
	var tokens TokenStream
	for len(text) > 0 {
		var token string
		token, text = tokenize(text, ls...)
		if len(token) == 0 {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens, text
}

func RemoveBlankTokens(s TokenStream) TokenStream { return Select(NoneBlank, s) }

func NatureLex(text string) []string {
	ss, _ := Lex(text, Spaces, NoSpaces)
	return ss
}
