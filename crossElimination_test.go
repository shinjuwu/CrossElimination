package elimination

import (
	"reflect"
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
		ca cardPos
	}
	ca := cardPos{
		posX: 0,
		posY: 0,
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{ca}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triggerPoint(tt.args.ca)
		})
	}
}

func Test_scanBoard(t *testing.T) {
	tests := []struct {
		name string
		want [][]cardPos
	}{
		// TODO: Add test cases.
		{"case1", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
