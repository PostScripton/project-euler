package solution

import (
	"fmt"
	"testing"
)

func TestFirstSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d digits", tt.args.digits), func(t *testing.T) {
			t.Parallel()
			if got := FirstSolution(tt.args.digits); got != tt.want {
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
