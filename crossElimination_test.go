package elimination

import (
	"testing"
)

func Test_transferCardsToBoard(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transferCardsToBoard()
		})
	}
}

func Test_triggerPoint(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triggerPoint(tt.args.x, tt.args.y)
		})
	}
}
