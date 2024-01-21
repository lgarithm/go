package parsec

import "os"

type Pos struct {
	Off      int
	Row, Col int
}

func ZeroPos() Pos {
	return Pos{
		Row: 1,
		Col: 1,
	}
}

func (p *Pos) PassRune(c rune) {
	p.Off += 1
	if c == '\n' {
		p.Row += 1
		p.Col = 1
	} else {
		p.Col += 1
	}
}

func (p *Pos) Pass(s string) {
	for _, c := range s {
		p.PassRune(c)
	}
}

type Token struct {
	String string
	Pos    Pos
}

func WithPos(ss TokenStream) []Token {
	p := ZeroPos()
	var ts []Token
	for _, t := range ss {
		ts = append(ts, Token{
			String: t,
			Pos:    p,
		})
		p.Pass(t)
	}
	return ts
}

func TokenizeFile(filename string, lex ...Tokenizer) ([]Token, string, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	s, t := Lex(string(bs), lex...)
	return WithPos(s), t, nil
}
