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
	posX int
	posY int
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
var test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 77, 77, 77, 16, 17, 18, 19, 99, 99, 99, 99, 24, 25}
var matchX = []cardPos{}
var matchY = []cardPos{}
var matchedCardBoard = [5][5]int{}
var destoryCard = []int{}

//var matchCard int

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

func scanBoard() [][]cardPos {
	res := [][]cardPos{}
	matchedCardBoard = [5][5]int{}
	for k, v := range board {
		for k1 := range v {
			card := cardPos{
				posX: k,
				posY: k1,
			}
			if !isMatched(card) {
				triggerPoint(card)
				if len(matchX) >= 3 {
					res = append(res, matchX)
				}
				if len(matchY) >= 3 {
					res = append(res, matchY)
				}
				fmt.Println(matchedCardBoard)
			}
		}
	}
	fmt.Println("res:", res)
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

func triggerPoint(card cardPos) {
	fmt.Println("=========================")
	transferCardsToBoard()
	fmt.Println("Trigger Point:", card)
	matchX = []cardPos{}
	matchY = []cardPos{}
	matchCard := board[card.posX][card.posY]
	fmt.Println("MatchCard:", matchCard)
	matchX = append(matchX, card)
	matchY = append(matchY, card)
	for i := 0; i < len(directArrays); i++ {
		direct := directArrays[i]
		nextCardPos := cardPos{
			posX: card.posX + direct[0],
			posY: card.posY + direct[1],
		}
		nextPoint(matchCard, nextCardPos, direct)
	}
	createMatchedCardBoard(matchCard)

	fmt.Println("matchX: ", matchX)
	fmt.Println("matchY: ", matchY)
}

func nextPoint(matchCard int, card cardPos, direct [2]int) {
	if card.posX < 0 || card.posX == len(board) ||
		card.posY < 0 || card.posY == len(board[0]) ||
		board[card.posX][card.posY] != matchCard {
		fmt.Println("Return Point:", card)
		return
	}

	saveCards(card, direct)
	nextCard := cardPos{
		posX: card.posX + direct[0],
		posY: card.posY + direct[1],
	}

	nextPoint(matchCard, nextCard, direct)
}

func saveCards(card cardPos, direct [2]int) {
	if direct[0] == 0 {
		matchY = append(matchY, card)
	} else {
		matchX = append(matchX, card)
	}
}

func createMatchedCardBoard(matchCard int) {
	if len(matchX) >= 3 {
		for _, v := range matchX {
			matchedCardBoard[v.posX][v.posY] = matchCard
		}
	}

	if len(matchY) >= 3 {
		for _, v := range matchY {
			matchedCardBoard[v.posX][v.posY] = matchCard
		}
	}
}

func isMatched(card cardPos) bool {
	if matchedCardBoard[card.posX][card.posY] != 0 {
		return true
	}
	return false
}

func elimination() {
	matchedLine := scanBoard()
	if len(matchedLine) == 0 {
		return
	}
	starsPatternNum := createStars(matchedLine)
	destoryCardPatternNum := createDestoryCards(matchedLine)

	updateBordList(destoryCardPatternNum, starsPatternNum)
}

func createDestoryCards(matchedLine [][]cardPos) []int {
	res := []int{}
	for _, line := range matchedLine {
		for _, pos := range line {
			destoryCardsPatternNum := gamePattern[pos.posX][pos.posY]
			res = append(res, destoryCardsPatternNum)
		}
	}

	return res
}

func createStars(matchedLine [][]cardPos) []int {
	res := []int{}
	for _, line := range matchedLine {
		for key, pos := range line {
			index := len(line) / 2
			if key == index {
				starCardPos := pos
				cardPatternNum := gamePattern[starCardPos.posX][starCardPos.posY]
				res = append(res, cardPatternNum)
			}

		}
	}
	return res
}

func updateBordList(destorys []int, stars []int) []int {
	for _, pos := range destorys {

	}
}
