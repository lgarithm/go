package parsec

var (
	Select = Filter[string]
)

func Filter[T any](p Pred[T], s []T) []T {
	var s1 []T
	for _, t := range s {
		if p(t) {
			s1 = append(s1, t)
		}
	}
	return s1
}
