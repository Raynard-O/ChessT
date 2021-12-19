package piece

import "Tesla_Vehicle_Software/board"

type InterfacePiece interface {
	SetPiece(x,y int, b board.InterfaceBoardLayout) (error, message )
	SetCoordinatesN(x int, b board.InterfaceBoardLayout) error
}
