package files

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/chijiokejoseph/go-xo/utils"
)

// Load tries to parse the data stored in filename, where
// filename is the path to a valid json file.
func Load(filename string) (*Database, error) {

	file, err := os.OpenFile(filename, os.O_CREATE+os.O_RDONLY, 0755)
	if err != nil {
		return nil, utils.WrapErr(err, ErrNoCreate)
	}

	// register callback to close file
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// read contents of file
	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, utils.WrapErr(err, ErrNoRead)
	}

	// if file is newly created and empty return empty db
	if len(contents) == 0 {
		return &Database{}, nil
	}

	var records []Record
	err = json.Unmarshal(contents, &records)
	if err != nil {
		return nil, utils.WrapErr(err, ErrUnMarshal)
	}

	return &Database{records}, nil
}
