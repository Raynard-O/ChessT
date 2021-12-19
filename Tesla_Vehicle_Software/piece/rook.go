package piece

import (
	"Tesla_Vehicle_Software/board"
	error2 "Tesla_Vehicle_Software/error"
	"math/rand"
	"time"
)

type message string
type Rook struct {

}


func (r Rook) SetPiece(x,y int, b board.InterfaceBoardLayout) (error, message) {

	_, found := b.GetBoardLayoutData().XCoordinates[x]
	_, found2 := b.GetBoardLayoutData().YCoordinates[y]

	if found && found2{
		return error2.InvalidPosition, " Tile already occupied"
	} else if  found{
		return error2.InvalidPosition, " Row already occupied"
	} else if  found2{
		return error2.InvalidPosition, " Column already occupied"
	}
	b.SetCoordinates(x,y)

	return nil, ""
}

func (r Rook) SetCoordinatesN(x int, b board.InterfaceBoardLayout) error {
	if x > b.GetBoardLayoutData().Height{
		return error2.InvalidNoRooks
	}
	s1 := rand.NewSource(time.Now().UnixNano())
	r3 := rand.New(s1)

	for i := 0; i < x; i++ {
		found := true
		var xCor, yCor int
		for found {
			xCor = r3.Intn(b.GetBoardLayoutData().Height+1)
			_, found = b.GetBoardLayoutData().XCoordinates[xCor]
			if xCor == 0 {
				found = true
			}
		}
		found = true
		for found {
			yCor = r3.Intn(b.GetBoardLayoutData().Height+1)
			_, found = b.GetBoardLayoutData().YCoordinates[yCor]
			if yCor == 0 {
				found = true
			}
		}
		b.SetCoordinates(xCor,yCor)
	}

	return nil
}
