package sortlib

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFlags() (SortFlags, error) {
	nextFlag := false
	nextNum := false
	sf := SortFlags{}
	for _, elem := range os.Args {
		for _, s := range elem {
			if s == '-' {
				nextFlag = true
				continue
			} else if nextFlag {
				switch s {
				case 'k':
					sf.K = true
					nextNum = true
				case 'r':
					sf.R = true
				case 'n':
					sf.N = true
				case 'u':
					sf.U = true
				case 'c':
					sf.C = true
				case 'M':
					sf.M = true
				case 'b':
					sf.B = true
				case 'h':
					sf.H = true
				default:
					nextFlag = false
				}
			} else if nextNum {
				iNum, err := strconv.Atoi(elem)
				if err != nil {
					return sf, fmt.Errorf("invalid num for flag k: %s", elem)
				}
				sf.KN = iNum
				nextNum = false
			}
		}
	}
	return sf, nil
}

func prepareKey(s string, sf SortFlags) (sortItem, error) {
	field := s

	if sf.KN > 0 {
		field = getTabField(s, sf.KN)
	}

	if sf.B {
		field = strings.TrimRight(field, " ")
	}

	if sf.M {
		mon, ok := months[field]
		if !ok {
			return sortItem{}, fmt.Errorf("invalid month %q in line %q", field, s)
		}
		return sortItem{raw: s, keyNum: mon, hasNum: true}, nil
	}

	if sf.H {
		n, err := parseHumanSize(field)
		if err != nil {
			if sf.KN > 0 {
				return sortItem{}, fmt.Errorf("can't convert human size key %q at line %q: %v", field, s, err)
			}
			return sortItem{}, fmt.Errorf("can't convert human size %q: %v", field, err)
		}
		return sortItem{raw: s, keyNum: n, hasNum: true}, nil
	}

	if sf.N {
		n, err := atoiStrict(field)
		if err != nil {
			if sf.KN > 0 {
				return sortItem{}, fmt.Errorf("can't convert key %q at line %q: %v", field, s, err)
			}
			return sortItem{}, fmt.Errorf("can't convert %q: %v", s, err)
		}
		return sortItem{raw: s, keyNum: n, hasNum: true}, nil
	}

	return sortItem{raw: s, keyStr: field}, nil
}

func uniqueStrings(in []string) []string {
	uniq := make(map[string]struct{}, len(in))
	res := make([]string, 0, len(in)/2)
	for _, s := range in {
		if _, ok := uniq[s]; !ok {
			uniq[s] = struct{}{}
			res = append(res, s)
		}
	}
	return res
}

func parseHumanSize(s string) (int, error) {
	x := strings.TrimSpace(s)
	if x == "" {
		return 0, fmt.Errorf("empty size")
	}

	if len(x) >= 2 && (x[len(x)-1] == 'B' || x[len(x)-1] == 'b') {
		x = x[:len(x)-1]
	}

	mult := 1.0
	if len(x) > 0 {
		last := x[len(x)-1]
		switch last {
		case 'K', 'k':
			mult = 1024
			x = x[:len(x)-1]
		case 'M', 'm':
			mult = 1024 * 1024
			x = x[:len(x)-1]
		case 'G', 'g':
			mult = 1024 * 1024 * 1024
			x = x[:len(x)-1]
		case 'T', 't':
			mult = 1024 * 1024 * 1024 * 1024
			x = x[:len(x)-1]
		}
	}

	val, err := strconv.ParseFloat(strings.TrimSpace(x), 64)
	if err != nil {
		return 0, err
	}
	res := int(math.Round(val * mult))
	return res, nil
}
