package config

import (
	"errors"
	"fmt"

	"github.com/chijiokejoseph/go-xo/inputs"
	pl "github.com/chijiokejoseph/go-xo/players"
)

// InputHandle prompts the user to select either of
// Handles "X" or "O" as their handle for a game round.
//
// Returns:
//   - pl.Handle: the handle selected by the player
//   - pl.Handle: the opposite handle to be used by the
//     computer
//   - error: nil if operation successful else an error
func InputHandle(options []pl.Handle) (pl.Handle, pl.Handle, error) {
	// convert options of []pl.Handle to []string
	menu := make([]string, len(options))
	for i, option := range options {
		menu[i] = string(option)
	}

	for {
		// prompt user to select a handle
		choice, err := inputs.InputOption(menu, (&inputs.Config{}).WithPrompt("Choose Handle").WithTitle("Player Handle Menu"))
		switch {
		// if user cancels current operation return
		case errors.Is(err, inputs.ErrQuit):
			return pl.Handle(""), pl.Handle(""), err
		// on other errors, display error warning and retry
		case err != nil:
			inputs.WarnPrint(err.Error())
			continue
		}

		// extract player handle
		playerIdx := choice.Index
		player := options[playerIdx]

		// set computer handle to the value following player handle
		computerIdx := (len(menu) - 1) - playerIdx
		computer := options[computerIdx]

		// display the selected player handle
		inputs.InfoPrint(fmt.Sprintf("You have selected Player '%s'\n", player))

		return player, computer, nil
	}
}
