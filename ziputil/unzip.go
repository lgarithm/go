package ziputil

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// Unzip unzips a gzip file
func Unzip(bs []byte) map[string][]byte {
	br := bytes.NewReader(bs)
	gr, err := gzip.NewReader(br)
	if err != nil {
		return nil
	}
	tr := tar.NewReader(gr)
	files := map[string][]byte{}
	for {
		hdr, err := tr.Next()
		if err != nil {
			break
		}
		bs, err := ioutil.ReadAll(tr)
		if err != nil {
			continue
		}
		files[hdr.Name] = bs
	}
	return files
}
