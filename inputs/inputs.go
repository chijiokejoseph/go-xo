package inputs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

)

var ErrQuit = errors.New("user has canceled the current operation")
var ErrParseFailed = errors.New("parsing user response to number failed")

var quitKeywords = []string{
	"quit",
	"exit",
}

func addQuitMsg(prompt string) string {
	newMsg := "Enter ('quit' or 'exit') to cancel current operation"
	return newMsg + "\n" + prompt
}

func InputStr(prompt string) (string, error) {
	// add quit message to prompt
	prompt = addQuitMsg(prompt)

	// display prompt
	fmt.Print(prompt)

	// read content from stdin
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	response = strings.Trim(response, "\n")
	response = strings.Trim(response, "\r")
	response = strings.Trim(response, " ")

	// convert response to lowercase
	lowered := strings.ToLower(response)
	// if lowercase response matches any quit keyword
	if slices.Contains(quitKeywords, lowered) {
		return "", ErrQuit
	}

	return response, nil
}

func InputNum(prompt string) (int, error) {
	// get response from user
	response, err := InputStr(prompt)
	if err != nil {
		return 0, err
	}

	// parse user response to integer
	responseNum, err := strconv.Atoi(response)
	if err != nil {
		return 0, ErrParseFailed
	}

	return responseNum, nil
}
