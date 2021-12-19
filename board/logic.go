package board

import (
"fmt"
)

//PrintBoard
// draws the board on std out
func (b *Board) PrintBoard(){
	for _, v := range b.BoardValue {
		fmt.Println(v)
	}
}


//SetCoordinates
// takes in the user entered coordinates (x,y) and populates the board
func (b *Board) SetCoordinates(x, y int)  {
	b.XCoordinates[x] = 1
	b.YCoordinates[y] = 1
	b.BoardValue[x-1][y-1] = 1
}


//GetBoard
//gives a double array for board
func (b *Board) GetBoard() [][]int {
	return b.BoardValue
}

