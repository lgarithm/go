package tr

import (
	"fmt"
	"io"
)

func title(w io.Writer, t string) {
	if len(t) > 0 {
		fmt.Fprintf(w, "\tSummery: %s\n", t)
	}
	fmt.Fprintf(w, "%s\n", hr)
}
