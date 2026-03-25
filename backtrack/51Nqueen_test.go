package backtrack

import "testing"

func TestSolveNQueensCounts(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{n: 1, want: 1},
		{n: 2, want: 0},
		{n: 3, want: 0},
		{n: 4, want: 2},
	}

	for _, tc := range tests {
		t.Run("n="+string(rune('0'+tc.n)), func(t *testing.T) {
			got := len(solveNQueens(tc.n))
			if got != tc.want {
				t.Fatalf("n=%d got=%d want=%d", tc.n, got, tc.want)
			}
		})
	}
}
