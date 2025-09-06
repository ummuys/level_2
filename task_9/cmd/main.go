package main

import (
	"errors"
	"fmt"
	"strings"
)

var errFirstDigit error = errors.New("string cannot start with a digit")

func unpack(str string) (string, error) {

	if str == "" {
		return "", nil
	} else if str[0] >= '0' && str[0] <= '9' {
		return "", errFirstDigit
	}

	var (
		prev   rune
		num    int
		hasNum bool
		res    strings.Builder
	)

	res.Grow(len(str)) // Помогаем уменьшить кол-во аллокаций

	for _, s := range str {
		if s >= '0' && s <= '9' {
			if prev == 92 {
				prev = s
				continue
			}
			hasNum = true
			num *= 10
			num += int(s - '0')
		} else {
			if hasNum {
				res.WriteString(strings.Repeat(string(prev), num))
				hasNum = false
				num = 0
			} else if prev != 0 {
				res.WriteRune(prev)
			}
			prev = s
		}
	}

	//Доесть последний символ
	if hasNum {
		res.WriteString(strings.Repeat(string(prev), num))
		hasNum = false
		num = 0
	} else if prev != 0 {
		res.WriteRune(prev)
	}

	return res.String(), nil
}

func main() {
	str, err := unpack("qwe\\4")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
