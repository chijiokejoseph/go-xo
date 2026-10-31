package grid

var slots = [3][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

var winSlots = [][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{1, 4, 7},
	{2, 5, 8},
	{3, 6, 9},
	{1, 5, 9},
	{3, 5, 7},
}

func newGrid() *Grid {

	// create 1D slice to hold 2D array numbers `slots`
	freeSlots := make([]int, 9)

	// create 2D array to hold board contents
	var values [3][3]any

	for i, row := range slots {
		for j, cell := range row {
			// append 2D array `slots` cell to 1D slice `freeSlots`
			freeSlots[3*i+j] = cell
			// convert `cell` to any to store values
			values[i][j] = any(cell)
		}
	}

	// instantiate Board struct
	board := Grid{
		FreeSlots: freeSlots,
		Slots:     slots,
		Values:    values,
	}

	return &board
}

func newGridManager(board *Grid) *GridManager {
	ps := GridManager{
		Player:   make([]int, 1),
		Computer: make([]int, 1),
		Grid:     board,
		WinSlots: winSlots,
	}

	return &ps
}

func NewGridManager() *GridManager {
	board := newGrid()
	return newGridManager(board)
}
