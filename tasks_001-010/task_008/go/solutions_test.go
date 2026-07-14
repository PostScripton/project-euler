package solution

const benchmarkNumber = 13

type args struct {
	adjacentDigits int
}

var tests = []struct {
	args args
	want int
}{
	{
		args: args{
			adjacentDigits: -1,
		},
		want: 0,
	},
	{
		args: args{
			adjacentDigits: 0,
		},
		want: 0,
	},
	{
		args: args{
			adjacentDigits: 1,
		},
		want: 9,
	},
	{
		args: args{
			adjacentDigits: 4,
		},
		want: 5832,
	},
	{
		args: args{
			adjacentDigits: 13,
		},
		want: 23514624000,
	},
	{
		args: args{
			adjacentDigits: 20,
		},
		want: 240789749760000,
	},
}
