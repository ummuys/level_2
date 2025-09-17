package sortlib

import (
	"fmt"
	"sort"
)

func ChooseSort(buf []string, fileName string, sf SortFlags) ([]string, string, error) {
	if sf.U {
		buf = uniqueStrings(buf)
	}

	rows := make([]sortItem, len(buf))
	for i, s := range buf {
		item, err := prepareKey(s, sf)
		if err != nil {
			return nil, "", err
		}
		rows[i] = item
	}

	if sf.C {
		if msg := checkSorted(rows, sf, fileName); msg != "" {
			return nil, msg, nil
		}
		return nil, "", nil
	}

	sortSlice(rows, sf)

	res := make([]string, len(rows))
	for i := range rows {
		res[i] = rows[i].raw
	}
	return res, "", nil
}

func sortSlice(rows []sortItem, sf SortFlags) {
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].hasNum {
			if rows[i].keyNum != rows[j].keyNum {
				if sf.R {
					return rows[i].keyNum > rows[j].keyNum
				}
				return rows[i].keyNum < rows[j].keyNum
			}
			if sf.R {
				return rows[i].raw > rows[j].raw
			}
			return rows[i].raw < rows[j].raw
		}
		if rows[i].keyStr != rows[j].keyStr {
			if sf.R {
				return rows[i].keyStr > rows[j].keyStr
			}
			return rows[i].keyStr < rows[j].keyStr
		}

		if sf.R {
			return rows[i].raw > rows[j].raw
		}
		return rows[i].raw < rows[j].raw
	})
}

func checkSorted(rows []sortItem, sf SortFlags, fileName string) string {
	for i := 0; i < len(rows)-1; i++ {
		a, b := rows[i], rows[i+1]
		if a.hasNum {
			if sf.R {
				if a.keyNum < b.keyNum {
					return fmt.Sprintf("wbsort: %s:%d: disorder: %v", fileName, i+2, a.keyNum)
				}
			} else {
				if a.keyNum > b.keyNum {
					return fmt.Sprintf("wbsort: %s:%d: disorder: %v", fileName, i+2, b.keyNum)
				}
			}
		} else {
			if sf.R {
				if a.keyStr < b.keyStr {
					return fmt.Sprintf("wbsort: %s:%d: disorder: %v", fileName, i+2, a.keyStr)
				}
			} else {
				if a.keyStr > b.keyStr {
					return fmt.Sprintf("wbsort: %s:%d: disorder: %v", fileName, i+2, b.keyStr)
				}
			}
		}
	}
	return ""
}
