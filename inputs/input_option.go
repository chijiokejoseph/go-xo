package inputs

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

const SPACES = 70

var Banner = strings.Repeat("=", SPACES)
var SubBanner = strings.Repeat("-", SPACES)

type Option struct {
	Index int
	Value string
}

type Config struct {
	Prompt string
	Title  string
}

func Error(msg string) string {
	return fmt.Sprintf("❌ %s\n", msg)
}

func ErrorPrint(msg string) {
	fmt.Println(Error(msg))
}

func Warn(msg string) string {
	return fmt.Sprintf("⚠️ %s\n", msg)
}

func WarnPrint(msg string) {
	fmt.Println(Warn(msg))
}

func Good(msg string) string {
	return fmt.Sprintf("✅ %s\n", msg)
}

func GoodPrint(msg string) {
	fmt.Println(Good(msg))
}

func Info(msg string) string {
	return fmt.Sprintf("ℹ️ %s\n", msg)
}

func InfoPrint(msg string) {
	fmt.Println(Info(msg))
}

func Menu(msg string) string {
	return Banner + fmt.Sprintf("\n🗓️ %-*s 🗓️\n", SPACES-10, msg) + Banner
}

func MenuPrint(msg string) {
	fmt.Println(Menu(msg))
}

func SubMenu(msg string) string {
	return SubBanner + fmt.Sprintf("\nℹ️ %-*s ℹ️\n", SPACES-10, msg) + SubBanner
}

func SubMenuPrint(msg string) {
	fmt.Println(SubMenu(msg))
}

func (c *Config) WithPrompt(prompt string) *Config {
	c.Prompt = prompt
	return c
}

func (c *Config) WithTitle(title string) *Config {
	c.Title = title
	return c
}

func NewOption() *Option {
	return &Option{}
}

// InputOption
// Provides an API for selecting an option from a slice of options.
func InputOption(menu []string, config *Config) (Option, error) {
	if len(menu) == 0 {
		log.Fatalln("menu cannot be empty")
	}

	// configure title
	title := config.Title
	if title == "" {
		title = "Menu"
	}
	// configure prompt
	prompt := config.Prompt
	fallback := fmt.Sprintf("Select an option (1 - %d): ", len(menu))
	if prompt == "" {
		prompt = fallback
	} else {
		prompt = prompt + "\n" + fallback
	}

	// display menu options
	fmt.Println(Menu(title))
	for i, option := range menu {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Println()

	choice, err := InputNum(prompt)
	if err != nil {
		return Option{}, err
	}

	if choice < 1 || choice > len(menu) {
		errMsg := fmt.Sprintf("Invalid option number entered, expected value to be within the range 1 - %d, Got: %d\n", len(menu), choice)
		fmt.Printf("❌ %s\n", errMsg)
		return Option{}, errors.New(errMsg)
	}

	return Option{Index: choice - 1, Value: menu[choice-1]}, nil
}

func InputBool(config *Config) (bool, error) {
	menu := []string{"Yes", "No"}
	choice, err := InputOption(menu, config)
	if err != nil {
		return false, err
	}

	return choice.Index == 0, nil
}
