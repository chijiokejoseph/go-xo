package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/chijiokejoseph/go-xo/files"
	"github.com/chijiokejoseph/go-xo/inputs"
	"github.com/chijiokejoseph/go-xo/logic"
)

var dbChan = make(chan *files.Database)
var errChan = make(chan error)

func Goodbye() {
	fmt.Println()
	fmt.Println()
	inputs.InfoPrint("The great Go is sad to see you go 😔😔😔")
	inputs.InfoPrint("But all good things must come to an end 🚫")
	inputs.InfoPrint("Safe Travels 🚀")
	fmt.Println()
}

func setup() {
	go func() {
		data, err := files.DataPath()
		// if cannot get data path
		if err != nil {
			// stop and close db channel
			close(dbChan)
			// send error to errChan
			errChan <- err
			return
		}

		path := filepath.Join(data, "records.json")
		db, err := files.Load(path)
		// if cannot load file records
		if err != nil {
			// stop and close db channel
			close(dbChan)
			// send err to errChan
			errChan <- err
			return
		}

		dbChan <- db
		errChan <- nil
	}()
}

func cleanup(record files.Record) {
	go func() {
		data, err := files.DataPath()
		// if cannot get data path
		if err != nil {
			// stop and close db channel
			close(dbChan)
			// send error to errChan
			errChan <- err
			return
		}

		path := filepath.Join(data, "records.json")
		err = files.Save(record, path)
		// if cannot load file records
		if err != nil {
			// stop and close db channel
			close(dbChan)
			// send err to errChan
			errChan <- err
			return
		}

		errChan <- nil
	}()
}

func main() {
	// run resource setup
	setup()

	// display introduction messages to the user
	inputs.MenuPrint("GO-X-O")
	fmt.Println()
	fmt.Println()
	inputs.InfoPrint("This is a simple tic-tac-toe game built in Go 🚀")
	fmt.Println()
	fmt.Println()

	// get username
	username, err := inputs.InputStr("Enter your name: ")

	switch {
	case errors.Is(err, inputs.ErrQuit):
		// if user quits the operation
		inputs.WarnPrint("You have canceled the current operation")
		fmt.Println()

		// ask the user if he wishes to exit
		choice, err := inputs.InputBool((&inputs.Config{}).WithTitle("Exit Menu").WithPrompt("Do you wish to exit the game: "))
		// if there is an error, exit
		if err != nil {
			Goodbye()
			return
		}
		// if user chooses to exit
		if choice {
			Goodbye()
			return
		}
		// else set username default to player
		username = "Player"

	case err != nil:
		// for any other errors, display error and exit
		fmt.Println()
		msg := fmt.Sprintf("Program stopped due to error: %s\n", err)
		inputs.ErrorPrint(msg)
		Goodbye()
		return
	}

	// display another message to the user
	fmt.Println()
	fmt.Println("Hello Challenger " + username + ". Your name appears formidable. We wish you victory!")
	fmt.Println()
	fmt.Println()

	record := files.Record{
		Name: username,
	}

game:
	for {
		err := logic.Session(username, &record)

		switch {
		case errors.Is(err, inputs.ErrQuit):
			// if user quits the operation
			inputs.WarnPrint("You have canceled the current operation")
			fmt.Println()

			// ask the user if he wishes to exit
			choice, err := inputs.InputBool((&inputs.Config{}).WithTitle("Exit Menu").WithPrompt("Do you wish to exit the game: "))
			// if there is an error, exit
			if err != nil {
				break game
			}
			// if user chooses to exit
			if choice {
				break game
			}

		case err != nil:
			// for any other errors, display error and exit
			fmt.Println()
			msg := fmt.Sprintf("Program stopped due to error: %s\n", err)
			inputs.ErrorPrint(msg)
			break game
		}

		// ask the user if he wishes to exit
		fmt.Println()
		choice, err := inputs.InputBool((&inputs.Config{}).WithTitle("Exit Menu").WithPrompt("Do you wish to exit the game: "))

		// if there is an error, exit
		if err != nil {
			break game
		}
		// if user chooses to exit
		if choice {
			break game
		}

		fmt.Println()
		fmt.Println()

	}

	// read data from db channel
	db, ok := <-dbChan
	if !ok {
		err := <-errChan
		msg := fmt.Sprintf("Creating or Reading existing records failed due to error: %s", err)
		inputs.ErrorPrint(msg)

		fmt.Println()
		fmt.Println(msg)

		Goodbye()
		return
	}

	// defer closing of db channel
	defer func() {
		close(dbChan)
	}()

	// initiate separate routine to save file data
	cleanup(record)

	// calculate current record score
	record.Score = record.Wins - (record.Draws + record.Losses)
	record.HighScore = record.Score

	// calculate previous high score
	highScore := db.MaxHighScore()

	// check if user has beaten the preious highscore
	if record.Score > highScore {
		logic.Sleep(500)

		// print out a special congratulatory message for beating the highscore
		fmt.Println()
		fmt.Println()
		inputs.MenuPrint("More Congratulations " + username + " 🎉🎉🎊🎊🎉🎉")
		count := 40
		fmt.Println(strings.Repeat("🎆", count))
		fmt.Println(strings.Repeat("🎉", count))
		fmt.Println(strings.Repeat("🎊", count))
		fmt.Println(
			fmt.Sprintf(
				"You really are Him, you've beat your old High Score %d\n"+
					"Now it's %d"+strings.Repeat("🥳", 7)+". Hail Honored One.",
				highScore, record.Score,
			),
		)
	}

	// start session record display by printing a gap
	fmt.Println()
	fmt.Println()

	oldRecord, err := db.Get(record.Name)

	// if there is no old record
	if err != nil {
		// print out only the current session record
		msg := fmt.Sprintf("%s. Your record is: \nWINS: %d | DRAWS: %d | LOSSES: %d | SCORE: %d", record.Name, record.Wins, record.Draws, record.Losses, record.Score)
		inputs.MenuPrint(msg)
	} else {

		// print old session records
		msg := fmt.Sprintf("%s. Your previous records are: \nWINS: %d | DRAWS: %d | LOSSES: %d | SCORE: %d", oldRecord.Name, oldRecord.Wins, oldRecord.Draws, oldRecord.Losses, oldRecord.Score)
		inputs.MenuPrint(msg)

		// print current session records
		msg = fmt.Sprintf("%s. Your current record is: \nWINS: %d | DRAWS: %d | LOSSES: %d | SCORE: %d", record.Name, record.Wins, record.Draws, record.Losses, record.Score)
		inputs.MenuPrint(msg)
	}

	Goodbye()

	err = <-errChan
	defer func() {
		close(errChan)
	}()

	if err != nil {
		inputs.ErrorPrint(err.Error())
		return
	}
}
