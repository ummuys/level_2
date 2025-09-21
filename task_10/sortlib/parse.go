package sortlib

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func readFlags() (SortFlags, error) {
	var sf SortFlags

	flag.BoolVar(&sf.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&sf.N, "n", false, "сортировка по числовому значению")
	flag.BoolVar(&sf.U, "u", false, "убрать дубликаты")
	flag.BoolVar(&sf.C, "c", false, "проверить отсортированность")
	flag.BoolVar(&sf.M, "M", false, "сортировка по месяцам")
	flag.BoolVar(&sf.B, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&sf.H, "h", false, "сортировка по суффиксам (например, 2K, 1G)")
	flag.IntVar(&sf.K, "k", 0, "сортировать по столбцу")

	flag.Parse()

	return sf, nil
}

// Выборка из сырых строчек нужных данных
func prepareKey(s string, sf SortFlags) (sortItem, error) {
	// s - raw, field - clear
	field := s

	// Сдвиг по табам
	if sf.K > 0 {
		field = getTabField(s, sf.K)
	}

	// Удалить пробелы слева (так работает оригинальный сорт)
	if sf.B {
		field = strings.TrimLeft(field, " ")
	}

	// Работа с месяцами
	if sf.M {
		mon, ok := months[field]
		if !ok {
			return sortItem{}, fmt.Errorf("invalid month %q in line %q", field, s)
		}
		return sortItem{raw: s, keyNum: mon, hasNum: true}, nil
	}

	// Конверт в дату
	if sf.H {
		n, err := parseHumanSize(field)
		if err != nil {
			if sf.K > 0 {
				return sortItem{}, fmt.Errorf("can't convert human size key %q at line %q: %v", field, s, err)
			}
			return sortItem{}, fmt.Errorf("can't convert human size %q: %v", field, err)
		}
		return sortItem{raw: s, keyNum: n, hasNum: true}, nil
	}

	// Конверт в числа
	if sf.N {
		n, err := atoiStrict(field)
		if err != nil {
			if sf.K > 0 {
				return sortItem{}, fmt.Errorf("can't convert key %q at line %q: %v", field, s, err)
			}
			return sortItem{}, fmt.Errorf("can't convert %q: %v", s, err)
		}
		return sortItem{raw: s, keyNum: n, hasNum: true}, nil
	}

	return sortItem{raw: s, keyStr: field}, nil
}

// Уникальные строчки
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

// Парс даты
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
