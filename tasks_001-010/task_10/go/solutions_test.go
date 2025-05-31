package solution

const benchmarkNumber = 2_000_000

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
		want: 17,
	},
	{
		args: args{
			until: 2_000_000,
		},
		want: 142913828922,
	},
}
