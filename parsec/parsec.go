package parsec

type match_n []string

func (m match_n) parse(s TokenStream) (*void, TokenStream) {
	for _, t := range m {
		var x *void
		if x, s = Try(Match1(t))(s); x == nil {
			return nil, s
		}
	}
	return &empty, s
}

func MatchN(ts ...string) Parse[void] { return match_n(ts).parse }

type many[T any] Parse[T]

func (m many[T]) parse(s TokenStream) (*[]T, TokenStream) {
	var p = Try(Parse[T](m))
	var xs []T
	for {
		var x *T
		if x, s = p(s); x == nil {
			return &xs, s
		}
		xs = append(xs, *x)
	}
}

func Many[T any](p Parse[T]) Parse[[]T] { return many[T](p).parse }

type Tuple3[A any, B any, C any] struct {
	a A
	b B
	c C
}

func T2[A any, B any, C any](t Tuple3[A, B, C]) C { return t.c }

type p3[A any, B any, C any] struct {
	a Parse[A]
	b Parse[B]
	c Parse[C]
}

func (p p3[A, B, C]) parse(s TokenStream) (*Tuple3[A, B, C], TokenStream) {
	x, s := p.a(s)
	if x == nil {
		return nil, s
	}
	y, s := p.b(s)
	if y == nil {
		return nil, s
	}
	z, s := p.c(s)
	if z == nil {
		return nil, s
	}
	return &Tuple3[A, B, C]{*x, *y, *z}, s
}

func Triple[A any, B any, C any](a Parse[A], b Parse[B], c Parse[C]) Parse[Tuple3[A, B, C]] {
	return p3[A, B, C]{a, b, c}.parse
}

func Map[A any, B any](f func(A) B, p Parse[A]) Parse[B] {
	return func(s TokenStream) (*B, TokenStream) {
		a, s := p(s)
		if a == nil {
			return nil, s
		}
		b := f(*a)
		return &b, s
	}
}

func Between[A any, B any, T any](a Parse[A], b Parse[B], p Parse[T]) Parse[T] {
	return Map(func(t Tuple3[A, T, B]) T { return t.b }, Triple(a, p, b))
}

func Sides[A any, B any, C any](t Tuple3[A, B, C]) (A, C) { return t.a, t.c }

func bracket[T any](a string, b string, p Parse[T]) Parse[T] { return Between(Match(a), Match(b), p) }

func Quoted[T any](q string, p Parse[T]) Parse[T] { return bracket(q, q, p) }

func Paren[T any](p Parse[T]) Parse[T]   { return bracket(`(`, `)`, p) }
func Brace[T any](p Parse[T]) Parse[T]   { return bracket(`{`, `}`, p) }
func Bracket[T any](p Parse[T]) Parse[T] { return bracket(`[`, `]`, p) }
func Chevron[T any](p Parse[T]) Parse[T] { return bracket(`<`, `>`, p) }

func AnyElse(ss []string) Parse[string] { return SymbolSet(ss).Others }

func FollowBy[B any, T any](b Parse[B], p Parse[T]) Parse[T] { return Between(Epilson, b, p) }

func LeadBy[A any, T any](a Parse[A], p Parse[T]) Parse[T] { return Between(a, Epilson, p) }

func TillEnd[T any](p Parse[T]) Parse[T] { return FollowBy(Eof, p) }

func Try[T any](p Parse[T]) Parse[T] {
	return func(s TokenStream) (*T, TokenStream) {
		if x, s1 := p(s); x != nil {
			return x, s1
		}
		return nil, s
	}
}

type sepBy[S any, T any] struct {
	q Parse[S]
	p Parse[T]
}

func (s sepBy[S, T]) parse(ts TokenStream) (*[]T, TokenStream) {
	var xs []T
	x0, ts1 := s.p(ts)
	if x0 == nil {
		return &xs, ts
	}
	xs = append(xs, *x0)
	x1s, ts2 := Many(LeadBy(s.q, s.p))(ts1)
	if x1s == nil {
		return &xs, ts1
	}
	xs = append(xs, *x1s...)
	return &xs, ts2
}

func SepBy[S any, T any](q Parse[S], p Parse[T]) Parse[[]T] {
	return sepBy[S, T]{q: q, p: p}.parse
}

func CL[T any](p Parse[T]) Parse[[]T] { return SepBy(Match(`,`), p) }

func FirstMatch[T any](ps ...Parse[T]) Parse[T] {
	return func(s TokenStream) (*T, TokenStream) {
		for _, p := range ps {
			if x, s := Try(p)(s); x != nil {
				return x, s
			}
		}
		return nil, s
	}
}
