package grid

import (
	"slices"

	pl "github.com/chijiokejoseph/go-xo/players"
)

type GridManager struct {
	Player   []int
	Computer []int
	Grid     *Grid
	WinSlots [][3]int
}

func Contains(src [][3]int, test []int) bool {
	testTrue := []bool{true, true, true}

	var isTrue []bool

	for _, collection := range src {
		isTrue = make([]bool, 3)
		for i, item := range collection {
			isTrue[i] = slices.Contains(test, item)
		}
		if slices.Equal(isTrue, testTrue) {
			return true
		}
	}
	return false
}

// *Positions.GradeComputer checks if the computer has secured a complete set
// or rather has won the user. Returns nil if the computer has not won and there are
// still spaces on the board.
//
// Returns players.Lose if the computer wins,or returns players.Draw if there are no
// more spaces on the game's board for that round.
func (gm *GridManager) GradeComputer(pos int) pl.Status {
	gm.Computer = append(gm.Computer, pos)
	if Contains(gm.WinSlots, gm.Computer) {
		return pl.Lose
	} else if len(gm.Grid.FreeSlots) == 0 {
		return pl.Draw
	} else {
		return pl.Pending
	}
}

// *Positions.GradeUser checks if the user has secured a complete set
// or rather has won the user. Returns nil if the user has not won and there are
// still spaces on the board.
//
// Returns players.Win if the user wins, or returns players.Draw if there are no
// more spaces on the game's board for that round.
func (gm *GridManager) GradeUser(pos int) pl.Status {
	gm.Player = append(gm.Player, pos)
	if Contains(gm.WinSlots, gm.Player) {
		return pl.Win
	} else if len(gm.Grid.FreeSlots) == 0 {
		return pl.Draw
	} else {
		return pl.Pending
	}
}
