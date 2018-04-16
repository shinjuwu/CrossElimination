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

type cardPos struct {
	positionX int
	positionY int
}

var gamePattern = [][]int{
	{21, 20, 19, 18, 17},
	{22, 7, 6, 5, 16},
	{23, 8, 1, 4, 15},
	{24, 9, 2, 3, 14},
	{25, 10, 11, 12, 13},
}

var directArrays = [4][2]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}
var board = [5][5]int{}
var test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 1, 1, 1, 1}
var matchCards = []cardPos{}

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

func eliminate() map[cardPos][]cardPos {
	res := map[cardPos][]cardPos{}
	for k, v := range board {
		for k1 := range v {
			triggerPoint(k, k1)
		}
	}
	return res
}

func transferCardsToBoard() {
	for k, v := range gamePattern {
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

func triggerPoint(x int, y int) {
	transferCardsToBoard()
	for i := 0; i < len(directArrays); i++ {
		direct := directArrays[i]
		nextCardPos := cardPos{
			positionX: x + direct[0],
			positionY: y + direct[1],
		}
		matchCard := board[x][y]

		nextPoint(matchCard, nextCardPos, direct)
	}
	fmt.Println(matchCards)
}

func nextPoint(matchCard int, cardPosition cardPos, direct [2]int) {
	if cardPosition.positionX < 0 || cardPosition.positionX == len(board) ||
		cardPosition.positionY < 0 || cardPosition.positionY == len(board[0]) ||
		board[cardPosition.positionX][cardPosition.positionY] != matchCard {
		return
	}

	saveCards(cardPosition)
	nextCard := cardPos{
		positionX: cardPosition.positionX + direct[0],
		positionY: cardPosition.positionY + direct[1],
	}

	nextPoint(matchCard, nextCard, direct)
}

func saveCards(card cardPos) {
	matchCards = append(matchCards, card)
}
