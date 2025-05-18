package solution

import (
	"testing"
)

func TestSecondSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			if got := SecondSolution(tt.args.exceedNumber); got != tt.want {
				t.Errorf("SecondSolution(%d) = %v, want %v", tt.args.exceedNumber, got, tt.want)
			}
		})
	}
}

func BenchmarkSecondSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecondSolution(benchmarkNumber)
	}
}
