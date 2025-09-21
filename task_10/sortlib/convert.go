package sortlib

import (
	"strconv"
	"strings"
)

// Удалить табы
func getTabField(s string, k int) string {
	if k <= 0 {
		return s
	}
	parts := strings.Split(s, "\t")
	if i := k - 1; i >= 0 && i < len(parts) {
		return parts[i]
	}
	return ""
}

// Конверт в числа
func atoiStrict(s string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, err
	}
	return n, nil
}
