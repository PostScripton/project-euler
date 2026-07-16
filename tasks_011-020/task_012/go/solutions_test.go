package solution

const benchmarkNumber = 500

type args struct {
	divisors int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			divisors: 5,
		},
		want: 28,
	},
	{
		args: args{
			divisors: 500,
		},
		want: 76576500,
	},
}
