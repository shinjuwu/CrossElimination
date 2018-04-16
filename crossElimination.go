package elimination

import (
	"container/list"
	"fmt"
)

//EliminationTask is used to illustract a Elimination game eliminate task
type EliminationTask struct {
	board [5][5]int
	cards *list.List
}
type node struct {
	index     int
	cardIndex int
}

var seq = [][]int{
	{5, 6, 7, 8, 9},
	{4, 19, 20, 21, 10},
	{3, 18, 25, 22, 11},
	{2, 17, 24, 23, 12},
	{1, 16, 15, 14, 13},
}
var board = [5][5]rune{}
var test = []rune{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

//New is use for create a elimination task
func New(initCards *list.List) *EliminationTask {
	et := &EliminationTask{
		board: [5][5]int{},
		cards: initCards,
	}
	return et
}

func (et *EliminationTask) transferCardsToBoard() {

}

func eliminate(cards *list.List) *list.List {
	return nil
}

func transferCardsToBoard() {
	for k, v := range seq {
		for k1, v1 := range v {
			board[k][k1] = test[v1-1]
		}
	}

	for k, v := range board {
		fmt.Printf("index: %d ->Cards: ", k)
		for _, v1 := range v {
			fmt.Printf(" %v", v1)
		}
		fmt.Println()
	}
}
