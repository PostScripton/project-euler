package solution

const benchmarkNumber = 3

type args struct {
	digits int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			digits: 2,
		},
		want: 9009,
	},
	{
		args: args{
			digits: 3,
		},
		want: 906609,
	},
	{
		args: args{
			digits: 4,
		},
		want: 99000099,
	},
}
