package sortlib_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ummuys/level_2/task_10/sortlib"
)

func TestChooseSort_BasicAndCombos(t *testing.T) {
	cases := []struct {
		numTest int
		input   []string
		flags   sortlib.SortFlags
		expect  []string
	}{
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
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%02d", tc.numTest), func(t *testing.T) {
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

func TestChooseSort_FlagK_ColumnSorting(t *testing.T) {
	cases := []struct {
		numTest int
		input   []string
		flags   sortlib.SortFlags
		expect  []string
	}{
		{
			numTest: 101,
			input: []string{
				"alice\t10",
				"bob\t2",
				"carol\t30",
			},
			flags:  sortlib.SortFlags{K: true, KN: 2, N: true},
			expect: []string{"bob\t2", "alice\t10", "carol\t30"},
		},
		{
			numTest: 102,
			input: []string{
				"b\tx",
				"a\tz",
				"c\ty",
			},
			flags:  sortlib.SortFlags{K: true, KN: 2},
			expect: []string{"b\tx", "c\ty", "a\tz"},
		},
		{
			numTest: 103,
			input: []string{
				"a\t10",
				"b\t2",
				"c\t2",
			},
			flags:  sortlib.SortFlags{K: true, KN: 2, N: true, R: true},
			expect: []string{"a\t10", "b\t2", "c\t2"},
		},
		{
			numTest: 104,
			input: []string{
				"a\t10\tJan",
				"b\t2\tFeb",
				"c\t30\tMar",
			},
			flags:  sortlib.SortFlags{K: true, KN: 3, M: true, R: true},
			expect: []string{"c\t30\tMar", "b\t2\tFeb", "a\t10\tJan"},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.numTest), func(t *testing.T) {
			out, _, err := sortlib.ChooseSort(tc.input, "", tc.flags)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(out, tc.expect) {
				t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, tc.expect)
			}
		})
	}
}

func TestChooseSort_FlagM_MonthNames(t *testing.T) {
	in := []string{"Apr", "Jan", "Dec", "Feb"}
	out, _, err := sortlib.ChooseSort(in, "", sortlib.SortFlags{M: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"Jan", "Feb", "Apr", "Dec"}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, want)
	}

	in2 := []string{"Apr", "Jan", "Dec", "Feb"}
	out2, _, err2 := sortlib.ChooseSort(in2, "", sortlib.SortFlags{M: true, R: true})
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}
	want2 := []string{"Dec", "Apr", "Feb", "Jan"}
	if !reflect.DeepEqual(out2, want2) {
		t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out2, want2)
	}
}

func TestChooseSort_FlagB_TrailingBlanks(t *testing.T) {
	in := []string{"apple   ", "  banana", "apple", "banana  "}
	out, _, err := sortlib.ChooseSort(in, "", sortlib.SortFlags{B: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// При B пробелы в конце/начале не влияют на порядок сравнения.
	// Мы не требуем их удаления в выводе, только порядок.
	want := []string{"apple", "apple   ", "  banana", "banana  "}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, want)
	}
}

func TestChooseSort_FlagC_CheckOnly_WithCombos(t *testing.T) {
	_, msg, err := sortlib.ChooseSort([]string{"a\t1", "b\t2", "c\t3"}, "", sortlib.SortFlags{K: true, C: true, KN: 2, N: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg != "" {
		t.Fatalf("expected empty message, got: %q", msg)
	}

	_, msg2, err2 := sortlib.ChooseSort([]string{"a\t2", "b\t1"}, "", sortlib.SortFlags{K: true, C: true, KN: 2, N: true})
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}
	if msg2 == "" {
		t.Fatalf("expected disorder message, got empty")
	}
}

func TestChooseSort_FlagH_HumanNumeric(t *testing.T) {
	cases := []struct {
		numTest int
		input   []string
		flags   sortlib.SortFlags
		expect  []string
	}{
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
		{
			numTest: 203,
			input:   []string{"1K", "1024", "1K"},
			flags:   sortlib.SortFlags{H: true, U: true},
			expect:  []string{"1K"},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.numTest), func(t *testing.T) {
			out, _, err := sortlib.ChooseSort(tc.input, "", tc.flags)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(out, tc.expect) {
				t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, tc.expect)
			}
		})
	}
}

func TestChooseSort_Combinations_KNH(t *testing.T) {
	in := []string{
		"a\t2K\tJan",
		"b\t1500\tFeb",
		"c\t1M\tMar",
		"d\t999\tJan",
	}
	out, _, err := sortlib.ChooseSort(in, "", sortlib.SortFlags{K: true, KN: 2, H: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{
		"d\t999\tJan",
		"b\t1500\tFeb",
		"a\t2K\tJan",
		"c\t1M\tMar",
	}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, want)
	}
}

func TestChooseSort_Combinations_KMUR(t *testing.T) {
	in := []string{
		"a\t10\tJan",
		"b\t10\tJan",
		"c\t10\tFeb",
		"d\t10\tMar",
		"e\t10\tFeb",
	}
	out, _, err := sortlib.ChooseSort(in, "", sortlib.SortFlags{K: true, KN: 3, M: true, R: true, U: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Reverse по месяцам: Mar > Feb > Jan; U=true — по значению (месяцу) оставить один из дубликатов
	want := []string{
		"d\t10\tMar",
		"c\t10\tFeb",
		"a\t10\tJan",
	}
	if !reflect.DeepEqual(out, want) {
		t.Fatalf("unexpected result:\n got: %#v\nwant: %#v", out, want)
	}
}
