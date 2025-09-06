package main

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		numTest  int
		in       string
		out      string
		hasError bool
		err      error
	}{
		{1, "", "", false, nil},
		{2, "abcd", "abcd", false, nil},
		{3, "a4bc2d5e", "aaaabccddddde", false, nil},
		{4, "й2ё", "ййё", false, nil},
		{5, "a12", "aaaaaaaaaaaa", false, nil},
		{6, "a0", "", false, nil},
		{7, "1", "", true, errFirstDigit},
		{8, "qwe\\4\\5", "qwe45", false, nil},
		{9, "qwe\\45", "qwe44444", false, nil},
		{10, "a\\3b2", "a3bb", false, nil},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()
			got, err := unpack(tc.in)

			if tc.hasError {
				if err == nil {
					t.Fatalf("testNum №%d: expected error %v, got nil", tc.numTest, tc.err)
				}
				if !errors.Is(err, tc.err) {
					t.Fatalf("testNum №%d: expected error %v, got %v", tc.numTest, tc.err, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("testNum №%d: unexpected error: %v", tc.numTest, err)
			}
			if got != tc.out {
				t.Fatalf("testNum №%d: got %q, want %q", tc.numTest, got, tc.out)
			}
		})
	}
}
