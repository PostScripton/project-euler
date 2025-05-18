package solution

const benchmarkNumber = 600_851_475_143

type args struct {
	number int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			number: 13_195,
		},
		want: 29,
	},
	{
		args: args{
			number: 600_851_475_143,
		},
		want: 6857,
	},
}
