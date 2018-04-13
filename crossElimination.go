package elimination

import "container/list"

//EliminationTask is used to illustract a Elimination game eliminate task
type EliminationTask struct {
	board [5][5]int
	cards *list.List
}

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
