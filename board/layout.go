package board

import err "Tesla_Vehicle_Software/error"

/*InterfaceBoardLayout
an interface that implements the board methods
*/
type InterfaceBoardLayout interface {
	SetBoard(width int) error
	GetBoardLayoutData() *Board
	GetBoard() [][]int
	SetCoordinates(x, y int)
	PrintBoard()
}

/*Board
A data structure that contains the properties of a board object
*/
type Board struct {
	Width int
	Height int
	BoardValue   [][]int
	XCoordinates map[int]int
	YCoordinates map[int]int
}

// New
//returns new instance of A board object
func New() *Board {
	//return new instance of board
	return &Board{
		XCoordinates: make(map[int]int),
		YCoordinates: make(map[int]int),
	}
}

//SetBoard
// takes in one input and sets the square dimension of the board
func (b *Board) SetBoard(width int) error {
	b.Width, b.Height = width, width
	// check if board dimension is valid
	if b.Width == 0 || b.Height == 0{
		return err.WrongFormat
	}
	//populate board
	for i := 0; i < b.Height; i++ {
		var myBoard []int
		for j := 0; j < b.Width; j++ {
			myBoard = append(myBoard, 0)
		}
		b.BoardValue = append(b.BoardValue, myBoard)
		myBoard = []int{}
	}
	return nil
}



//GetBoardLayoutData
//Get board properties
func (b *Board) GetBoardLayoutData() *Board {
	return b
}











