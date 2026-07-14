package solution

import (
	"testing"
)

func TestFirstSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			if got := FirstSolution(tt.args.exceedNumber); got != tt.want {
				t.Errorf("FirstSolution(%d) = %v, want %v", tt.args.exceedNumber, got, tt.want)
			}
		})
	}
}

func BenchmarkFirstSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstSolution(benchmarkNumber)
	}
}
