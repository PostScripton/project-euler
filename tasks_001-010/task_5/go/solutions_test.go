package solution

const benchmarkNumber = 20

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
		want: 2520,
	},
	{
		args: args{
			until: 20,
		},
		want: 232792560,
	},
}
