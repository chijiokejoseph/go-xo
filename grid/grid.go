package grid

import (
	"fmt"

	"github.com/chijiokejoseph/go-xo/inputs"
	pl "github.com/chijiokejoseph/go-xo/players"
)

// Grid
//
//   - FreeSlots: holds the available positions on the board
//   - Slots: holds the starting positions on the board
//   - Values: a dynamic list whose entries can be replaced with
//     'X' and 'O' and is used to create the board display
type Grid struct {
	FreeSlots []int
	Slots     [3][3]int
	Values    [3][3]any
}

// Place takes the userPos argument and locates its
// corresponding position on Board.Values, and
// updates its corresponding handle to the passed in players.Handle argument
// which is stored in userHandle.
func (grid *Grid) Place(pos int, handle pl.Handle) {
	done := false
	for i, row := range grid.Values {
		for j := range row {
			if (3*i + (j + 1)) == pos {
				grid.Values[i][j] = handle
				done = true
			}
			if done {
				break
			}
		}
	}
}

// Display creates a formatted string based on the contents
// of the board, after which the formatted string is displayed.
func (grid *Grid) Display() {
	// instantiate final output string
	var output string

	// get board values / contents
	values := grid.Values

	// format string for every board layer
	boardCase :=
		"+-----+-----+-----+\n" +
			"|  %v  |  %v  |  %v  |\n"

	// string for board final bottom
	boardBottom := "+-----+-----+-----+\n"

	// for each row in values
	for i, row := range values {
		// convert row array to slice for string embedding
		rowSlice := row[:]

		// append formatted string to output string
		output += fmt.Sprintf(boardCase, rowSlice...)

		// on final row of values add the board bottom layer
		if i == len(values)-1 {
			output += boardBottom
		}
	}

	// Display final output string
	inputs.InfoPrint("The Game board ♟️: ")
	fmt.Println(output)
}
