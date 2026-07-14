package solution

import (
	"testing"
)

func TestThirdSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			if got := ThirdSolution(tt.args.number); got != tt.want {
				t.Errorf("ThirdSolution(%d) = %v, want %v", tt.args.number, got, tt.want)
			}
		})
	}
}

func BenchmarkThirdSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThirdSolution(benchmarkNumber)
	}
}
