package machine

import (
	"io"
	"os"
)

// Source is the I/O bound type used to read files, extract bytecode, and
// initialize Pure machine from it.
//
// Example:
//
//	src, _ := SourceFromFile("executable")
//	proc, _ := src.Main()
//	proc.Feed(Unit{})
type Source struct {
	r io.Reader
}

func SourceFromFile(name string) (*Source, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return SourceFromReader(file), nil
}

func SourceFromReader(r io.Reader) *Source {
	return &Source{
		r: r,
	}
}

func (src *Source) Main() (Proc, error) {
	const start = 0
	const argc = 0

	data, err := src.data()
	if err != nil {
		return Proc{}, err
	}

	code, err := src.code()
	if err != nil {
		return Proc{}, err
	}

	return MakeProc(ReadOnly(data), ReadOnly(code), start, argc), nil
}

func (src *Source) data() ([]uint8, error) {
	length, err := src.readU64()
	if err != nil {
		return nil, err
	}
	return src.readN(length)
}

func (src *Source) code() ([]uint8, error) {
	return io.ReadAll(src.r)
}

func (src *Source) readU64() (uint64, error) {
	buf := make([]uint8, SizeU64)
	_, err := io.ReadFull(src.r, buf)
	return U8x8AsU64(buf), err
}

func (src *Source) readN(n uint64) (buf []uint8, err error) {
	buf = make([]uint8, n)
	_, err = io.ReadFull(src.r, buf)
	return
}
