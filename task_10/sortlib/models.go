package sortlib

type SortFlags struct {
	K  bool
	KN int
	N  bool
	R  bool
	U  bool
	M  bool
	B  bool
	C  bool
	H  bool
}

// TODO:
var months = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

type sortItem struct {
	raw    string
	keyStr string
	keyNum int
	hasNum bool
}
