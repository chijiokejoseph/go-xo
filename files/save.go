package files

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/chijiokejoseph/go-xo/utils"
)

// Save is a function that takes in a Record with and
// saves its fields to the json file specified as filename.
//
// # It automatically creates `filename` if it does not exist
//
// Returns:
//   - error: nil if operation succeeds else an error
func Save(newRecord Record, filename string) error {
	// TODO: FIX bug in Save function that prevents data from being written to the `data/records.json` file

	// open file in create or write modes
	file, err := os.OpenFile(filename, os.O_CREATE+os.O_WRONLY, 0755)
	if err != nil {
		return utils.WrapErr(err, ErrNoCreate)
	}

	// must close file
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// attempt to load content from file (empty if just created)
	db, err := Load(filename)
	if err != nil {
		return err
	}

	// use Records struct to append new record to existing
	db.Add(newRecord)

	// serialize records data to json
	stream, err := json.MarshalIndent(db.records, "", strings.Repeat(" ", 4))
	if err != nil {
		return utils.WrapErr(err, ErrMarshal)
	}

	// attempt to write json serialized bytes to fiel
	_, err = file.Write(stream)
	if err != nil {
		return utils.WrapErr(err, ErrNoWrite)
	}

	return nil
}
