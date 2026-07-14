package solution

import (
	"fmt"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-th prime number", tt.args.nthPrimeNumber), func(t *testing.T) {
			t.Parallel()
			if got := FirstSolution(tt.args.nthPrimeNumber); got != tt.want {
				t.Errorf("FirstSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFirstSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstSolution(benchmarkNumber)
	}
}
