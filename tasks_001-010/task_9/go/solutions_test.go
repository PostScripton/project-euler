package solution

const benchmarkNumber = 1000

type args struct {
	sum int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			sum: 1000,
		},
		want: 31875000,
	},
}
