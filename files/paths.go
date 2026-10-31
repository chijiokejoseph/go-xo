package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chijiokejoseph/go-xo/inputs"
)

func DataPath() (string, error) {
	// get current working directory
	dir, err := os.Getwd()
	if err != nil {
		msg := inputs.Error("Failed to get current working directory")
		fmt.Println(msg)
		return "", err
	}

	// construct data directory path
	data := filepath.Join(dir, "data")

	// create data directory if it doesn't exist
	err = os.MkdirAll(data, os.ModeDir)
	if err != nil {
		msg := inputs.Error("Failed to create 'data' folder in current working directory")
		fmt.Println(msg)
		return "", err
	}

	return data, nil
}
