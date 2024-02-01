package calc

import "testing"

func Test_getCenterOffset(t *testing.T) {
	type args struct {
		size int
		i    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "size 6 i 1", args: args{size: 6, i: 1}, want: 1},
		{name: "size 6 i 3", args: args{size: 6, i: 3}, want: 0},
		{name: "size 6 i 5", args: args{size: 6, i: 5}, want: 2},
		{name: "size 4 i 3", args: args{size: 5, i: 4}, want: 2},
		{name: "size 3 i 0", args: args{size: 3, i: 0}, want: 1},
		{name: "size 3 i 1", args: args{size: 3, i: 1}, want: 0},
		{name: "size 3 i 2", args: args{size: 3, i: 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCenterOffset(tt.args.size, tt.args.i); got != tt.want {
				t.Errorf("getCenterOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
