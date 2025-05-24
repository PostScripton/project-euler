package solution

const benchmarkNumber = 100

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
		want: 2640,
	},
	{
		args: args{
			until: 100,
		},
		want: 25164150,
	},
	{
		args: args{
			until: 1000,
		},
		want: 250166416500,
	},
}
