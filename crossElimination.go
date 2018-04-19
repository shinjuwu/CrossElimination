package elimination

import (
	"fmt"
)

//EliminationTask is used to illustract a Elimination game eliminate task
type EliminationTask struct {
	board [5][5]int
	cards []int
}

type node struct {
	symbol int
	posX   int
	posY   int
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
var test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 44, 44, 100, 20, 23, 100, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
var matchX = []node{}
var matchY = []node{}
var matchedCardBoard = [5][5]int{}
var destoryCard = []int{}

//var matchCard int

//New is use for create a elimination task
func New(initCards []int) *EliminationTask {
	et := &EliminationTask{
		board: [5][5]int{},
		cards: initCards,
	}
	return et
}

func (et *EliminationTask) transferCardsToBoard() {

}

func scanBoard() [][]node {
	res := [][]node{}
	matchedCardBoard = [5][5]int{}
	for k, v := range board {
		for k1 := range v {
			card := node{
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

func triggerPoint(card node) {
	fmt.Println("=========================")
	transferCardsToBoard()
	fmt.Println("Trigger Point:", card)
	matchX = []node{}
	matchY = []node{}
	card.symbol = board[card.posX][card.posY]
	fmt.Println("MatchCard:", card.symbol)
	matchX = append(matchX, card)
	matchY = append(matchY, card)
	for i := 0; i < len(directArrays); i++ {
		direct := directArrays[i]
		nextCard := node{
			symbol: card.symbol,
			posX:   card.posX + direct[0],
			posY:   card.posY + direct[1],
		}
		nextPoint(nextCard, direct)
	}
	createMatchedCardBoard()

	fmt.Println("matchX: ", matchX)
	fmt.Println("matchY: ", matchY)
}

func nextPoint(card node, direct [2]int) {
	if card.posX < 0 || card.posX == len(board) ||
		card.posY < 0 || card.posY == len(board[0]) {
		fmt.Println("Return Point:", card)
		return
	}
	if board[card.posX][card.posY] != 100 && board[card.posX][card.posY] != card.symbol {
		fmt.Println("Return Point:", card)
		return
	}
	if board[card.posX][card.posY] == 100 {
		saveCards(node{symbol: 100, posX: card.posX, posY: card.posY}, direct)
	} else {
		saveCards(card, direct)
	}

	nextCard := node{
		symbol: card.symbol,
		posX:   card.posX + direct[0],
		posY:   card.posY + direct[1],
	}
	nextPoint(nextCard, direct)
}

func saveCards(card node, direct [2]int) {
	if direct[0] == 0 {
		matchY = append(matchY, card)
	} else {
		matchX = append(matchX, card)
	}
}

func createMatchedCardBoard() {
	if len(matchX) >= 3 {
		for _, v := range matchX {
			matchedCardBoard[v.posX][v.posY] = v.symbol
		}
	}

	if len(matchY) >= 3 {
		for _, v := range matchY {
			matchedCardBoard[v.posX][v.posY] = v.symbol
		}
	}
}

func isMatched(card node) bool {
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
	fmt.Println("stars pos:", starsPatternNum)
	fmt.Println("destroy pos:", destoryCardPatternNum)
	updateBordList(destoryCardPatternNum, starsPatternNum)
	fmt.Println(test)
}

func createDestoryCards(matchedLine [][]node) []int {
	res := []int{}
	for _, line := range matchedLine {
		for _, v := range line {
			if v.symbol == 100 {
				starDestory := getStarDestoryCards(v)
				res = append(res, starDestory...)
			}
			destoryCardsPatternNum := gamePattern[v.posX][v.posY]
			res = append(res, destoryCardsPatternNum)
		}
	}

	return res
}

func createStars(matchedLine [][]node) []int {
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
	transferCardsToBoard()
	fmt.Println("=============================")
	for _, pos := range destorys {
		for _, starPos := range stars {
			test[pos-1] = -1
			if pos == starPos {
				test[pos-1] = 100
				break
			}
		}
	}
	transferCardsToBoard()
	fmt.Println("===================================")
	test = deleteDestoryCards()
	transferCardsToBoard()
	fmt.Println("Num of list", len(test))
	return test
}

func deleteDestoryCards() []int {
	res := []int{}
	for _, v := range test {
		if v != -1 {
			res = append(res, v)
		}
	}
	return res
}

func getStarDestoryCards(star node) []int {
	res := []int{}
	for i := 0; i < len(directArrays); i++ {
		direct := directArrays[i]
		nextCard := node{
			symbol: 100,
			posX:   star.posX + direct[0],
			posY:   star.posY + direct[1],
		}
		if isinBoard(nextCard) {
			nextCardPatternNum := gamePattern[nextCard.posX][nextCard.posY]
			res = append(res, nextCardPatternNum)
		}

	}
	return res
}

func isinBoard(card node) bool {
	if card.posX < 0 || card.posX == len(board) ||
		card.posY < 0 || card.posY == len(board[0]) {
		return false
	}
	return true
}
