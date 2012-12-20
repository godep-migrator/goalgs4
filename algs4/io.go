package algs4

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

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
