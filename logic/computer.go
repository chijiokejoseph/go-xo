package logic

import (
	"math/rand/v2"
	"slices"

	"github.com/chijiokejoseph/go-xo/grid"
	"github.com/chijiokejoseph/go-xo/inputs"
	pl "github.com/chijiokejoseph/go-xo/players"
)

// RandIntRange picks a random integer between min and max.
// max is inclusive.
func RandIntRange(min, max int) int {
	max += 1
	return rand.IntN(max-min) + min
}

// RandSelect randomly chooses an available position on the board
// The chosen position is returned and is assigned as the computer's
// position on the board.
func RandSelect(gm *grid.GridManager) int {
	freeSlots := gm.Grid.FreeSlots
	randIndex := RandIntRange(0, len(freeSlots)-1)
	return freeSlots[randIndex]
}

// Checks if the game user is about to secure a complete set
// If that is the case, the computer picks the last position left in that set
//
// Returns the position that the computer should pick to stop the user
// else returns -1 if the user is not about to secure a complete set.
func SmartSelect(gm *grid.GridManager) int {
	var (
		potentialSpot = -1
	)

	// range over each pair of winning slots
	for _, pair := range gm.WinSlots {
		var (
			playerLocs   []int
			computerLocs []int
			remLocs      []int
		)

		// range of each slot of the winning pair
		for _, item := range pair {

			// if the slot contains a player
			present := slices.Contains(gm.Player, item)

			if present {
				// add slot to player locations for current pair
				playerLocs = append(playerLocs, item)
			} else if slices.Contains(gm.Computer, item) {
				// add slot to computer locations for current pair
				computerLocs = append(computerLocs, item)
			} else {
				// add slot to remaining locations for current pair
				remLocs = append(remLocs, item)
			}

			if len(computerLocs) == 2 && len(remLocs) == 1 {
				// if the computer is about to win and there is a position left
				// set potentialSpot for computer to win
				potentialSpot = remLocs[0]
			} else if len(playerLocs) == 2 && len(remLocs) == 1 {
				// if the player is about to win and there is a free slot left
				// Set potentialSpot to block player win
				potentialSpot = remLocs[0]
			}

			if slices.Contains(gm.Grid.FreeSlots, potentialSpot) {
				// return it if its position is available
				return potentialSpot
			} else {
				// reset potentialSpot since it is not valid
				potentialSpot = -1
				continue
			}
		}
	}

	return potentialSpot
}

// ComputerSelect simply attempts to make a smart selection
// for the computer using smartSelect, in case the user
// is attempting to secure a complete set
//
// If smartSelect returns -1, i.e., the user is not about
// to complete a set, it defaults to using randSelect to
// pick the computer's position on the board.
func ComputerSelect(ps *grid.GridManager) int {
	output := SmartSelect(ps)
	if output == -1 {
		output = RandSelect(ps)
	}
	return output
}

// ComputerTurn encompasses all the functions that need to be executed
// for the computer's turn. It includes pauses using time.Sleep to give
// the computer a feeling of taking time to make decisions
func ComputerTurn(gm *grid.GridManager, computer pl.Player) int {
	// get free slots
	freeSlots := gm.Grid.FreeSlots

	// simulate computer taking time to make a decision
	Sleep(1000)
	inputs.InfoPrint("It's the computer's turn")
	Sleep(1500)

	// get computer position
	pos := ComputerSelect(gm)

	// remove position from free slots
	posIndex := slices.Index(freeSlots, pos)
	freeSlots = slices.Delete(freeSlots, posIndex, posIndex+1)
	gm.Grid.FreeSlots = freeSlots

	// update computer position on the display
	gm.Grid.Place(pos, computer.Handle)
	gm.Grid.Display()

	return pos
}
