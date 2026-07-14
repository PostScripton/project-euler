package solution

const benchmarkNumber = 10001

type args struct {
	nthPrimeNumber int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			nthPrimeNumber: -1,
		},
		want: 0,
	},
	{
		args: args{
			nthPrimeNumber: 0,
		},
		want: 0,
	},
	{
		args: args{
			nthPrimeNumber: 1,
		},
		want: 2,
	},
	{
		args: args{
			nthPrimeNumber: 2,
		},
		want: 3,
	},
	{
		args: args{
			nthPrimeNumber: 6,
		},
		want: 13,
	},
	{
		args: args{
			nthPrimeNumber: 10001,
		},
		want: 104743,
	},
}
