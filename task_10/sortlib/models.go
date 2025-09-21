package sortlib

// Структура с флагами
type SortFlags struct {
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
}

// Для работы с месяцами
var months = map[string]int{
	"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
	"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
	"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
}

// Структура для хранения оригинала и данных для сортировки
type sortItem struct {
	raw    string
	keyStr string
	keyNum int
	hasNum bool
}
