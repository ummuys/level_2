package main

import (
	"fmt"
	"sort"
	"strings"
)


func findAnagrams(words []string) (map[string][]string) {
	m := make(map[string][]string)

	for _, w := range words {
		w = strings.ToLower(w)
		r := []rune(w)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		wS := string(r)
		m[wS] = append(m[wS], w)
	}

	mRes := make(map[string][]string)

	for _, v := range m {
		if len(v) < 2 {
			continue
		}
		k := v[0]
		sort.Strings(v)
		mRes[k] = v
	}
	return mRes
}

func main() {
	fmt.Println(findAnagrams([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}))
}
