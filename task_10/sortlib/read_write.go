package sortlib

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// Выбор файла или os.Stdin
func chooseReader() (io.Reader, *os.File, error) {
	var (
		r    io.Reader
		file *os.File
		err  error
	)

	args := flag.Args()

	if len(args) > 0 {
		file, err = os.Open(args[0])
		if err != nil {
			return nil, nil, fmt.Errorf("не удалось открыть файл %q: %w", args[0], err)
		}
		r = file
	} else {
		r = os.Stdin
	}

	return r, file, nil
}

// Чтение файла
func readInfo(r io.Reader) ([]string, error) {
	br := bufio.NewReaderSize(r, 64*1024)
	var lines []string

	for {
		line, err := br.ReadString('\n')
		if len(line) > 0 {
			lines = append(lines, strings.TrimRight(line, "\r\n"))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("err read info: %w", err)
		}
	}
	return lines, nil
}

// Запись
func writeInfo(buf []string) error {
	w := bufio.NewWriter(os.Stdout)
	for _, s := range buf {
		if _, err := fmt.Fprintln(w, s); err != nil {
			return fmt.Errorf("err write info: %v", err)
		}
	}
	return w.Flush()
}
