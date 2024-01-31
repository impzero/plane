package calc

import (
	"reflect"
	"testing"
)

func TestCalculateMeanDirection(t *testing.T) {
	type args struct {
		vectors [][2]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "two vectors opposite direction, same length",
			args: args{vectors: [][2]float64{{-1.0, 0.0}, {1.0, 0.0}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateMeanDirection(tt.args.vectors); got != tt.want {
				t.Errorf("CalculateMeanDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVector(t *testing.T) {
	type args struct {
		vectors [][2]float64
	}
	tests := []struct {
		name string
		args args
		want [2]float64
	}{
		{
			name: "two opposite vectors should cancel each out",
			args: args{vectors: [][2]float64{{-2.0, 1.0}, {2.0, -1.0}}},
			want: [2]float64{0.0, 0.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVector(tt.args.vectors); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
