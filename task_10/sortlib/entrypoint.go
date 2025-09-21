package sortlib

import (
	"errors"
	"fmt"
)

// Основная функция
func WBSort() (err error) {

	flags, err := readFlags()
	if err != nil {
		return
	}

	r, f, err := chooseReader()
	if err != nil {
		return err
	}

	fName := "input_write"
	if f != nil {
		fName = f.Name()
		defer func() {
			if fErr := f.Close(); fErr != nil {
				err = errors.Join(err, fmt.Errorf("close %q: %w", fName, fErr))
			}
		}()
	}

	buf, err := readInfo(r)
	if err != nil {
		return
	}

	arr, msg, err := ChooseSort(buf, fName, flags)
	if err != nil {
		return
	}

	out := arr
	if msg != "" {
		out = []string{msg}
	}
	err = writeInfo(out)

	return
}
