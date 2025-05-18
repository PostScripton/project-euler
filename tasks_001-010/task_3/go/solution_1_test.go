package solution

import (
	"testing"
)

const benchmarkNumber = 600_851_475_143

func TestFirstSolution(t *testing.T) {
	tests := []struct {
		number int
		want   int
	}{
		{
			number: 13_195,
			want:   29,
		},
		{
			number: 600_851_475_143,
			want:   6857,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			got := FirstSolution(tt.number)
			if got != tt.want {
				t.Errorf("FirstSolution(%d); got: %d; want: %d", tt.number, got, tt.want)
				t.Fail()
			}
		})
	}
}

func BenchmarkFirstSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FirstSolution(benchmarkNumber)
	}
}
