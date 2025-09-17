package sortlib

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func chooseReader() (io.Reader, *os.File, error) {
	var (
		r    io.Reader
		file *os.File
		err  error
	)

	if len(os.Args) > 1 {
		file, err = os.Open(os.Args[1])
		if err != nil {
			return nil, nil, fmt.Errorf("bad name file")
		}
		r = file
	} else {
		r = os.Stdin
	}
	return r, file, nil
}

func readInfo(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	var buf []string

	for s.Scan() {
		buf = append(buf, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("err read info: %v", err)
	}

	return buf, nil
}

func writeInfo(buf []string) error {
	w := bufio.NewWriter(os.Stdout)
	for _, s := range buf {
		if _, err := fmt.Fprintln(w, s); err != nil {
			return fmt.Errorf("err write info: %v", err)
		}
	}
	return w.Flush()
}
