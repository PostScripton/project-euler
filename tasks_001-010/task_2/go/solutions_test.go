package solution

const benchmarkNumber = 4_000_000

type args struct {
	exceedNumber int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			exceedNumber: 4_000_000,
		},
		want: 4613732,
	},
}
