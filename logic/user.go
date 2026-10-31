package logic

import (
	"errors"
	"fmt"
	"slices"

	"github.com/chijiokejoseph/go-xo/grid"
	"github.com/chijiokejoseph/go-xo/inputs"
	pl "github.com/chijiokejoseph/go-xo/players"
)

// InputPos prompts the user to select a valid position on the displayed board.
// repeating the prompt for a certain number of trials after which it returns an error
// if the user has not yet entered a valid input.
func InputPos(validRange []int) (int, error) {

	prompt := "Select a free board position: "

	for {
		// prompt player to select a board position
		choice, err := inputs.InputNum(prompt)
		switch {
		// handle player quit requests
		case errors.Is(err, inputs.ErrQuit):
			return 0, err
		// warn and reprompt on any other errors
		case err != nil:
			inputs.Warn(err.Error())
			continue
		}

		// if player choice is not within valid range
		if !slices.Contains(validRange, choice) {

			// create collection of valid positions in []any
			collection := make([]any, len(validRange))
			for i, value := range validRange {
				collection[i] = value
			}

			// display warning message
			inputs.Warn(fmt.Sprintf(
				"Invalid input entered. Available positions on the board are %d",
				collection...,
			))

			continue
		}

		return choice, nil
	}
}

// UserTurn executes the game turn for the user, calling selectPos under the hood
// to get user input for the position on the board that the user intends to pick.
//
// Returns:
//   - int: player's selected position
//   - error: nil if successful else error
func UserTurn(board *grid.Grid, user pl.Player) (int, error) {
	Sleep(1000)

	inputs.InfoPrint("It's your turn " + user.Name)

	freeSlots := board.FreeSlots

	pos, err := InputPos(freeSlots)
	if err != nil {
		return 0, err
	}

	posIndex := slices.Index(freeSlots, pos)
	freeSlots = slices.Delete(freeSlots, posIndex, posIndex+1)

	board.FreeSlots = freeSlots
	board.Place(pos, user.Handle)
	board.Display()

	return pos, nil
}
