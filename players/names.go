package players

const (
	NameO = iota
	NameX
)

type Status int

const (
	Lose Status = iota
	Win
	Draw
	Pending
	Cancel
)

type Color string

const RED Color = "\033[31m"
const GREEN Color = "\033[32m"
const RESET Color = "\033[0m"

type Handle string

var Handles = []Handle{X, O}

var (
	X Handle = Handle(RED + "X" + RESET)
	O Handle = Handle(GREEN + "O" + RESET)
)

type Player struct {
	Name   string
	Handle Handle
}
