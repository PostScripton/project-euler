package solution

import (
	"fmt"
	"testing"
)

func TestThirdSolution(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d digits", tt.args.digits), func(t *testing.T) {
			t.Parallel()
			if got := ThirdSolution(tt.args.digits); got != tt.want {
				t.Errorf("ThirdSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkThirdSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThirdSolution(benchmarkNumber)
	}
}
