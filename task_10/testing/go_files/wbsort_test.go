package sortlib_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ummuys/level_2/task_10/sortlib"
)

func TestChooseSort_AllCases(t *testing.T) {
	cases := []struct {
		numTest int
		input   []string
		flags   sortlib.SortFlags
		expect  []string
	}{
		// База
		{numTest: 1, input: []string{"banana", "apple", "cherry"}, flags: sortlib.SortFlags{}, expect: []string{"apple", "banana", "cherry"}},
		{numTest: 2, input: []string{"banana", "apple", "cherry"}, flags: sortlib.SortFlags{R: true}, expect: []string{"cherry", "banana", "apple"}},
		{numTest: 3, input: []string{"apple", "banana", "apple", "cherry", "banana"}, flags: sortlib.SortFlags{U: true}, expect: []string{"apple", "banana", "cherry"}},
		{numTest: 4, input: []string{"10", "2", "30"}, flags: sortlib.SortFlags{N: true}, expect: []string{"2", "10", "30"}},
		{numTest: 5, input: []string{"10", "2", "30"}, flags: sortlib.SortFlags{N: true, R: true}, expect: []string{"30", "10", "2"}},
		{numTest: 6, input: []string{"10", "2", "10", "2", "30"}, flags: sortlib.SortFlags{N: true, U: true}, expect: []string{"2", "10", "30"}},
		{numTest: 7, input: []string{"b", "a", "b", "c", "a"}, flags: sortlib.SortFlags{R: true, U: true}, expect: []string{"c", "b", "a"}},
		{numTest: 8, input: []string{"003", "3", "2", "10", "2", "10"}, flags: sortlib.SortFlags{N: true, R: true, U: true}, expect: []string{"10", "3", "003", "2"}},
		{numTest: 9, input: []string{"only"}, flags: sortlib.SortFlags{}, expect: []string{"only"}},
		{numTest: 10, input: []string{}, flags: sortlib.SortFlags{R: true, U: true, N: false}, expect: []string{}},
		{numTest: 11, input: []string{"x", "x", "x"}, flags: sortlib.SortFlags{}, expect: []string{"x", "x", "x"}},
		{numTest: 12, input: []string{"x", "x", "x"}, flags: sortlib.SortFlags{U: true}, expect: []string{"x"}},
		{numTest: 13, input: []string{"0", "-1", "2", "-10", "0"}, flags: sortlib.SortFlags{N: true}, expect: []string{"-10", "-1", "0", "0", "2"}},
		{numTest: 14, input: []string{"0", "-1", "2", "-10", "0", "2"}, flags: sortlib.SortFlags{N: true, R: true, U: true}, expect: []string{"2", "0", "-1", "-10"}},
		{numTest: 15, input: []string{"010", "002", "30"}, flags: sortlib.SortFlags{N: true}, expect: []string{"002", "010", "30"}},
		{numTest: 16, input: []string{"", "a", " ", "b", ""}, flags: sortlib.SortFlags{}, expect: []string{"", "", " ", "a", "b"}},
		{numTest: 17, input: []string{"apple", "Banana", "Apple", "banana"}, flags: sortlib.SortFlags{}, expect: []string{"Apple", "Banana", "apple", "banana"}},
		{numTest: 18, input: []string{"a", "A", "a", "B", "b", "B"}, flags: sortlib.SortFlags{R: true, U: true}, expect: []string{"b", "a", "B", "A"}},
		{numTest: 19, input: []string{"2000000000", "1000000000", "1500000000"}, flags: sortlib.SortFlags{N: true}, expect: []string{"1000000000", "1500000000", "2000000000"}},
		{numTest: 20, input: []string{"3", "2", "10", "0"}, flags: sortlib.SortFlags{N: true, R: true}, expect: []string{"10", "3", "2", "0"}},
		{numTest: 21, input: []string{"0002", "2", "02", "10", "010", "10"}, flags: sortlib.SortFlags{N: true, U: true}, expect: []string{"0002", "02", "2", "010", "10"}},
		{numTest: 22, input: []string{"zz", "aa", "zz", "mm", "aa", "bb"}, flags: sortlib.SortFlags{U: true}, expect: []string{"aa", "bb", "mm", "zz"}},
		{numTest: 23, input: []string{"k", "k", "k"}, flags: sortlib.SortFlags{R: true, U: true}, expect: []string{"k"}},
		{
			numTest: 101,
			input: []string{
				"alice\t10",
				"bob\t2",
				"carol\t30",
			},
			flags:  sortlib.SortFlags{K: 2, N: true},
			expect: []string{"bob\t2", "alice\t10", "carol\t30"},
		},
		{
			numTest: 102,
			input: []string{
				"b\tx",
				"a\tz",
				"c\ty",
			},
			flags:  sortlib.SortFlags{K: 2},
			expect: []string{"b\tx", "c\ty", "a\tz"},
		},
		{
			numTest: 104,
			input: []string{
				"a\t10\tJan",
				"b\t2\tFeb",
				"c\t30\tMar",
			},
			flags:  sortlib.SortFlags{K: 3, M: true, R: true},
			expect: []string{"c\t30\tMar", "b\t2\tFeb", "a\t10\tJan"},
		},

		// Месяцы
		{
			numTest: 151,
			input:   []string{"Apr", "Jan", "Dec", "Feb"},
			flags:   sortlib.SortFlags{M: true},
			expect:  []string{"Jan", "Feb", "Apr", "Dec"},
		},
		{
			numTest: 152,
			input:   []string{"Apr", "Jan", "Dec", "Feb"},
			flags:   sortlib.SortFlags{M: true, R: true},
			expect:  []string{"Dec", "Apr", "Feb", "Jan"},
		},

		// B
		{
			numTest: 171,
			input:   []string{"apple   ", "  banana", "apple", "banana  "},
			flags:   sortlib.SortFlags{B: true},
			expect:  []string{"apple", "apple   ", "  banana", "banana  "},
		},

		// H
		{
			numTest: 201,
			input:   []string{"2K", "1500", "1M", "999"},
			flags:   sortlib.SortFlags{H: true},
			expect:  []string{"999", "1500", "2K", "1M"}, // 2K=2048
		},
		{
			numTest: 202,
			input:   []string{"10G", "900M", "1G", "512K"},
			flags:   sortlib.SortFlags{H: true, R: true},
			expect:  []string{"10G", "1G", "900M", "512K"},
		},

		// Комбинации K+N+H
		{
			numTest: 251,
			input: []string{
				"a\t2K\tJan",
				"b\t1500\tFeb",
				"c\t1M\tMar",
				"d\t999\tJan",
			},
			flags: sortlib.SortFlags{K: 2, H: true},
			expect: []string{
				"d\t999\tJan",
				"b\t1500\tFeb",
				"a\t2K\tJan",
				"c\t1M\tMar",
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%03d", tc.numTest), func(t *testing.T) {
			out, _, err := sortlib.ChooseSort(tc.input, "", tc.flags)
			if err != nil {
				t.Fatalf("testNum = %d: unexpected error: %v", tc.numTest, err)
			}
			if !reflect.DeepEqual(out, tc.expect) {
				t.Fatalf("testNum = %d: unexpected result:\n got:  %#v\n want: %#v", tc.numTest, out, tc.expect)
			}
		})
	}
}

func TestChooseSort_CheckOnly(t *testing.T) {
	cases := []struct {
		numTest        int
		input          []string
		flags          sortlib.SortFlags
		expectMsgEmpty bool
	}{
		{
			numTest:        301,
			input:          []string{"a\t1", "b\t2", "c\t3"},
			flags:          sortlib.SortFlags{C: true, K: 2, N: true},
			expectMsgEmpty: true,
		},
		{
			numTest:        302,
			input:          []string{"a\t2", "b\t1"},
			flags:          sortlib.SortFlags{C: true, K: 2, N: true},
			expectMsgEmpty: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%03d", tc.numTest), func(t *testing.T) {
			_, msg, err := sortlib.ChooseSort(tc.input, "", tc.flags)
			if err != nil {
				t.Fatalf("testNum = %d: unexpected error: %v", tc.numTest, err)
			}
			if tc.expectMsgEmpty && msg != "" {
				t.Fatalf("testNum = %d: expected empty message, got: %q", tc.numTest, msg)
			}
			if !tc.expectMsgEmpty && msg == "" {
				t.Fatalf("testNum = %d: expected disorder message, got empty", tc.numTest)
			}
		})
	}
}
