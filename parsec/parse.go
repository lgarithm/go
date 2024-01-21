package parsec

type Parse[T any] func(TokenStream) (*T, TokenStream)

func TODO[T any](ts TokenStream) (*T, TokenStream) {
	// fmt.Printf("TODO")
	return nil, ts
}

func PredNext(p Pred[string], s TokenStream) bool {
	if len(s) == 0 {
		return false
	}
	return p(s[0])
}

func MatchNext(x string, s TokenStream) bool { return PredNext(Equals(x), s) }

type void = struct{}

var empty void

func Epilson(s TokenStream) (*void, TokenStream) { return &empty, s }

func Eof(s TokenStream) (*void, TokenStream) {
	if len(s) == 0 {
		return &empty, s
	}
	return nil, s
}

func Any(s TokenStream) (*string, TokenStream) {
	if len(s) == 0 {
		return nil, s
	}
	return &s[0], s[1:]
}

type match1 string

func (m match1) parse(s TokenStream) (*void, TokenStream) {
	if MatchNext(string(m), s) {
		return &empty, s[1:]
	}
	return nil, s
}

func Match1(t string) Parse[void] { return match1(t).parse }

var Match = Match1

func CanParse[T any](Parse[T]) bool { return true }
