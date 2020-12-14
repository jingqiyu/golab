package train

import (
	"reflect"
	"testing"
)

func TestTurnRight(t *testing.T) {
	/*type args struct {
		nowDirection Direction
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "turn Right ",
			args: args{
				nowDirection: "N",
			},
			want: Direction("E"),
		},
		{
			name: "turn Right ",
			args: args{
				nowDirection: "E",
			},
			want: Direction("S"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TurnRight(tt.args.nowDirection); got != tt.want {
				t.Errorf("TurnRight() = %v, want %v", got, tt.want)
			}
		})
	}*/
}

func TestInit(t *testing.T) {
	type args struct {
		pos       position
		direction Direction
	}
	tests := []struct {
		name string
		args args
		want CarStatus
	}{
		{
			name: "init 0 , 0 N ",
			args: args{
				pos:       position{0, 0},
				direction: "N",
			},
			want: CarStatus{
				X:         0,
				Y:         0,
				Direction: "N",
			},
		},
		{
			name: "init 0 , 0 S ",
			args: args{
				pos:       position{0, 0},
				direction: "S",
			},
			want: CarStatus{
				X:         0,
				Y:         0,
				Direction: "S",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.pos, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
