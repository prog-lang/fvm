package machine

import (
	"io"
	"os"
)

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

func (src *Source) MakeCmd() (Cmd, error) {
	const start = 0
	data, err := src.data()
	if err != nil {
		return Cmd{}, err
	}
	code, err := src.code()
	if err != nil {
		return Cmd{}, err
	}
	return MakeCmd(NewROM(data), NewROM(code), start), nil
}

func (src *Source) data() ([]uint8, error) {
	length, err := src.readI32()
	if err != nil {
		return nil, err
	}
	return src.readN(length)
}

func (src *Source) code() ([]uint8, error) {
	return io.ReadAll(src.r)
}

func (src *Source) readI32() (int32, error) {
	buf := make([]uint8, SizeI32)
	_, err := io.ReadFull(src.r, buf)
	return U8x4AsI32(buf), err
}

func (src *Source) readN(n int32) (buf []uint8, err error) {
	buf = make([]uint8, n)
	_, err = io.ReadFull(src.r, buf)
	return
}
