package algs4

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
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

func (in *In) ReadDouble() (float64, error) {
	line, err := in.Inbuf.ReadString('\n')

	if err != nil {
		return float64(0), err
	}

	dbl, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return float64(0), err
	}

	return dbl, nil
}

func (in *In) ReadInt() (int64, error) {
	line, err := in.Inbuf.ReadString('\n')

	if err != nil {
		return int64(0), err
	}

	i64, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
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

func ReadInts(inbuf io.Reader) ([]int, error) {
	ints := make([]int, 0)

	var i64 int64

	fileReader := bufio.NewReader(inbuf)
	line, err := fileReader.ReadString('\n')

	for err == nil {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			i64, err = strconv.ParseInt(line, 10, 32)

			if err != nil {
				return nil, err
			}

			ints = append(ints, int(i64))
		}

		line, err = fileReader.ReadString('\n')
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	return ints, nil
}
