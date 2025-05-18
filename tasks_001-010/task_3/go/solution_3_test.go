package solution

import (
	"testing"
)

func TestThirdSolution(t *testing.T) {
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

			got := ThirdSolution(tt.number)
			if got != tt.want {
				t.Errorf("ThirdSolution(%d); got: %d; want: %d", tt.number, got, tt.want)
				t.Fail()
			}
		})
	}
}

func BenchmarkThirdSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThirdSolution(benchmarkNumber)
	}
}
