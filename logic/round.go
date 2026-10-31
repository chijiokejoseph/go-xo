package logic

import (
	"github.com/chijiokejoseph/go-xo/grid"
	pl "github.com/chijiokejoseph/go-xo/players"
)

func Round(gm *grid.GridManager, user, computer pl.Player) (pl.Status, error) {
	// get game board
	board := gm.Grid
	// display game board
	board.Display()

	for {
		// if player is X, player should start first
		if user.Handle == pl.X {
			// player turn
			userPos, err := UserTurn(board, user)
			if err != nil {
				return pl.Cancel, err
			}

			// grade player status
			status := gm.GradeUser(userPos)
			// if status is not pending, the game has concluded
			if status != pl.Pending {
				return status, nil
			}

			// computer turn
			computerPos := ComputerTurn(gm, computer)
			// grade computer status
			status = gm.GradeComputer(computerPos)
			// if status is non-nil, the game has concluded
			if status != pl.Pending {
				return status, nil
			}

			continue
		}

		// computer starts first this time
		computerPos := ComputerTurn(gm, computer)
		status := gm.GradeComputer(computerPos)
		if status != pl.Pending {
			return status, nil
		}

		playerPos, err := UserTurn(board, user)
		if err != nil {
			return pl.Cancel, err
		}
		status = gm.GradeUser(playerPos)
		if status != pl.Pending {
			return status, nil
		}
	}
}
