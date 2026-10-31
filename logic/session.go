package logic

import (
	"errors"
	"fmt"
	"log"

	"github.com/chijiokejoseph/go-xo/config"
	"github.com/chijiokejoseph/go-xo/files"
	"github.com/chijiokejoseph/go-xo/grid"
	"github.com/chijiokejoseph/go-xo/inputs"
	pl "github.com/chijiokejoseph/go-xo/players"
)

func remark(status pl.Status, user pl.Player) string {
	switch status {
	case pl.Win:
		return fmt.Sprintf("Congrats 🎉🎉🎉 %s. You won this round", user.Name)
	case pl.Lose:
		return fmt.Sprintf("😭 Three Missed Calls from IQ %s.\nLeave a message: You are a failure, you can never make it", user.Name)
	case pl.Draw:
		return "There comes a time in life when you realize that\n" + "You and dog have the same brain"
	default:
		log.Fatal("unreachable state reached designates program bugs")
		return ""
	}
}

func Session(username string, record *files.Record) error {
	// increment number of rounds played by 1
	record.Rounds++

	userHandle, computerHandle, err := config.InputHandle(pl.Handles)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	msg := fmt.Sprintf("%s. You are %s | Computer is %s", username, userHandle, computerHandle)
	inputs.InfoPrint(msg)

	// instantiate user and computer handles
	user := pl.Player{
		Name:   username,
		Handle: userHandle,
	}
	computer := pl.Player{
		Name:   "Computer",
		Handle: computerHandle,
	}

	// instantiate grid manager
	gm := grid.NewGridManager()

	// play a round
	status, err := Round(gm, user, computer)

	switch {
	case errors.Is(err, inputs.ErrQuit):
		return err
	case err != nil:
		return err
	}

	// get remark message
	msg = remark(status, user)

	// display remark message
	fmt.Println()
	fmt.Println()
	inputs.SubMenuPrint(msg)

	// check for user victory
	switch status {
	case pl.Draw:
		// update round stats
		record.Draws++

	case pl.Win:
		record.Wins++
	case pl.Lose:
		record.Losses++
	}

	return nil
}
