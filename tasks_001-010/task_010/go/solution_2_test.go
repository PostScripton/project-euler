package solution

import (
	"fmt"
	"testing"
)

func TestSecondSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("until %d", tt.args.until), func(t *testing.T) {
			t.Parallel()
			if got := SecondSolution(tt.args.until); got != tt.want {
				t.Errorf("SecondSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSecondSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SecondSolution(benchmarkNumber)
	}
}
