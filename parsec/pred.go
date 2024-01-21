package parsec

import "strings"

type Pred[T any] func(T) bool

type RunePred = Pred[rune]

func Not[T any](p Pred[T]) Pred[T] { return func(x T) bool { return !p(x) } }

func Equals[T comparable](x T) Pred[T] { return func(y T) bool { return x == y } }

func HasPrefix(p string) Pred[string] { return func(s string) bool { return strings.HasPrefix(s, p) } }

func Initial(c rune) Pred[string] {
	return func(s string) bool {
		rs := []rune(s)
		return len(rs) > 0 && rs[0] == c
	}
}
