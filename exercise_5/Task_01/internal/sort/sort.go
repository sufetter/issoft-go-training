package sort

import (
	"sort"
)

func Sort(data []any, f func(i int, j int) bool) ([]any, error) {
	sort.Slice(data, f)
	return data, nil
}
