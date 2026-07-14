package solution

const benchmarkNumber = 1000

type args struct {
	until int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			until: 10,
		},
		want: 23,
	},
	{
		args: args{
			until: 1000,
		},
		want: 233168,
	},
}
