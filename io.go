package algs4

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"unicode"
)

type In struct {
	Inbuf *bufio.Reader
	Fd    *os.File
}

var (
	Stdin *In
)

func init() {
	*(&Stdin) = &In{
		Inbuf: bufio.NewReader(os.Stdin),
		Fd:    os.Stdin,
	}
}

func NewIn(filename string) (*In, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	newIn := &In{
		Inbuf: bufio.NewReader(fd),
		Fd:    fd,
	}
	return newIn, nil
}

func NewInFromReader(inbuf io.Reader) *In {
	return &In{
		Inbuf: bufio.NewReader(inbuf),
		Fd:    nil,
	}
}

func (in *In) ReadDouble() (float64, error) {
	str, err := in.ReadString()
	if err != nil {
		return float64(0), err
	}

	dbl, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return float64(0), err
	}

	return dbl, nil
}

func (in *In) ReadInt() (int64, error) {
	str, err := in.ReadString()
	if err != nil {
		return int64(0), err
	}

	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(0), err
	}

	return i64, nil
}

func (in *In) IsEmpty() bool {
	_, err := in.Inbuf.Peek(1)
	if err != nil {
		return true
	}

	return false
}

func (in *In) ReadString() (string, error) {
	if in.IsEmpty() {
		return "", io.EOF
	}

	tok := &bytes.Buffer{}
	inTok := false

	for !in.IsEmpty() {
		b, err := in.Inbuf.ReadByte()
		if err != nil {
			return string(bytes.TrimSpace(tok.Bytes())), err
		}

		if unicode.IsSpace(rune(b)) {
			if inTok {
				return string(tok.Bytes()), nil
			}

			continue
		}

		inTok = true
		tok.WriteByte(b)
	}

	return "", io.EOF
}

func ReadInts(inbuf io.Reader) ([]int64, error) {
	ints := make([]int64, 0)

	in := NewInFromReader(inbuf)
	for !in.IsEmpty() {
		i64, err := in.ReadInt()
		if err != nil {
			return nil, err
		}

		ints = append(ints, i64)
	}

	return ints, nil
}
