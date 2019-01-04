package git

import (
	"crypto/sha1"
	"fmt"
)

// HashObject implements the git hash-object command
func HashObject(bs []byte) string {
	h := sha1.New()
	h.Write([]byte("blob "))
	h.Write([]byte(fmt.Sprintf("%d", len(bs))))
	h.Write([]byte{0})
	h.Write(bs)
	return fmt.Sprintf("%x\n", h.Sum(nil))
}
