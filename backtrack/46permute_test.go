package backtrack

import (
	"reflect"
	"sort"
	"testing"
)

func normalizePermutations(x [][]int) [][]int {
	out := make([][]int, len(x))
	for i := range x {
		out[i] = append([]int(nil), x[i]...)
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
	return out
}

func TestPermuteExample(t *testing.T) {
	got := normalizePermutations(permute([]int{1, 2, 3}))
	want := normalizePermutations([][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	})

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestSolveNQueensN2HasNoSolutions(t *testing.T) {
	got := solveNQueens(2)

	if len(got) != 0 {
		t.Fatalf("solveNQueens(2) = %v, want no solutions", got)
	}
}
