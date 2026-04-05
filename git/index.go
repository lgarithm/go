package git

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

// var fmt = message.NewPrinter(language.English)

// https://mincong.io/2018/04/28/git-index/
type Index struct {
	Entries []Entry
}

type Header struct {
	Magic   uint32
	Version uint32
	N       uint32
}

func (h Header) String() string {
	return fmt.Sprintf("%x %d %d", h.Magic, h.Version, h.N)
}

type SHA1 [20]byte

type GitTime struct {
	Sec  uint32
	Nano uint32
}

func (t GitTime) Time() time.Time {
	return time.Unix(int64(t.Sec), int64(t.Nano))
}

type EntryHeader struct {
	CTime  GitTime // 8, 0
	MTime  GitTime // 8, 16
	Device uint32  // 4, 20
	Info   uint32  // 4, 24
	Mode   uint32  // 4, 28
	UID    uint32  // 4, 32
	GID    uint32  // 4, 36
	Size   uint32  // 4, 40
	SHA1   SHA1    // 20, 60
	Flags  uint16  // 2, 62
}

var t0 = time.Now()

func (h EntryHeader) String() string {
	// var s = ago.S
	s := func(d time.Duration) string { return d.String() }
	f := func(t GitTime) string {
		return s(t0.Sub(t.Time()))
	}
	return fmt.Sprintf("%24s %24s size:%12d", f(h.CTime), f(h.MTime), h.Size)
}

type Entry struct {
	EntryHeader
	Path string
}

func (h Entry) String() string {
	return fmt.Sprintf("%s %s", h.EntryHeader, h.Path)
}

var endian = binary.BigEndian

func LoadIndexFile(p string) (*Index, error) {
	var idx Index
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var h Header
	if err := binary.Read(r, endian, &h); err != nil {
		return nil, err
	}
	for range h.N {
		e, err := readEntry(r)
		if err != nil {
			break
		}
		idx.Entries = append(idx.Entries, *e)
	}
	// TODO: parse extension
	// var sum SHA1
	// if err := binary.Read(r, endian, &sum); err != nil {
	// 	return nil, err
	// }
	return &idx, nil
}

func readEntry(r *bufio.Reader) (*Entry, error) {
	var e Entry
	if err := binary.Read(r, endian, &e.EntryHeader); err != nil {
		return nil, err
	}
	var n int
	e.Path, n = readStr(r)
	n += 62
	if p := (8 - n%8) % 8; p > 0 {
		var pad [8]byte
		if _, err := io.ReadFull(r, pad[:p]); err != nil {
			return nil, err
		}
	}
	return &e, nil
}

func readStr(r *bufio.Reader) (string, int) {
	var bs []byte
	for {
		b, err := r.ReadByte()
		if err != nil {
			return string(bs), len(bs) + 1
		}
		if b == 0 {
			break
		}
		bs = append(bs, b)
	}
	return string(bs), len(bs) + 1
}

func IndexPath(p string) string { return path.Join(p, `.git/index`) }
