package lab

import (
	"testing"
)

func TestHitToggle(t *testing.T) {
	type args struct {
		buss string
		uid  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "WhiteList",
			args: args{buss: "supportComment", uid: "1916000086559"},
			want: true,
		},

		{
			name: "WhiteListWrong",
			args: args{buss: "supportComment", uid: "19160000865"},
			want: false,
		},

		{
			name: "WhiteListWrong",
			args: args{buss: "supportComment", uid: "19160000861"},
			want: false,
		},

		{
			name: "WhiteListWrong",
			args: args{buss: "supportComment", uid: "1916000086"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HitToggle(tt.args.buss, tt.args.uid); got != tt.want {
				t.Errorf("HitToggle() = %v, want %v", got, tt.want)
			}
		})
	}
}
