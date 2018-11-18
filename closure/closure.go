package closure

import (
	"context"

	"github.com/lgarithm/go/control"
)

// Closure provides a scoped context to run a function
type Closure struct {
	Clear func() error
	Open  func() (context.Context, error)
	Close func() error
}

// Run runs a function in the Closure
func (c Closure) Run(f func(ctx context.Context) error) error {
	if c.Clear != nil {
		if err := c.Clear(); err != nil {
			return control.LogError(err)
		}
	}
	defer func() { control.LogError(c.Close()) }()
	ctx, err := c.Open()
	if err != nil {
		return control.LogError(err)
	}
	return control.LogError(f(ctx))
}
