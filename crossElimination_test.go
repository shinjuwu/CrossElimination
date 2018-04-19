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
		ca node
	}
	ca := node{
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
		want [][]node
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

func Test_elimination(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elimination()
		})
	}
}

func Test_updateBordList(t *testing.T) {
	type args struct {
		destorys []int
		stars    []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{[]int{21, 22, 23, 15, 14, 13}, []int{22, 14}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateBordList(tt.args.destorys, tt.args.stars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBordList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStarDestoryCards(t *testing.T) {
	transferCardsToBoard()
	type args struct {
		star node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"case1", args{node{100, 0, 0}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStarDestoryCards(tt.args.star); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStarDestoryCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
