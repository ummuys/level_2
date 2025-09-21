package sortlib

// Структура с флагами
type SortFlags struct {
<<<<<<< HEAD
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
=======
	K  bool // have, work
	KN int
	N  bool // have, work
	R  bool // have, work
	U  bool // have, work
	M  bool
	B  bool // have, work
	C  bool // have, work
	H  bool
>>>>>>> dbd42c121a2be6d7cd474a67df6921aaa89f04a7
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
