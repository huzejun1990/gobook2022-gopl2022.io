package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

// 	"gopl2022.io/ch4/treesort"
func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	// treesort.Sort(data) // 同一个路径下，不需要加包名
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
