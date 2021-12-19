package piece

import (
	"Tesla_Vehicle_Software/board"
	"testing"
)


var err error
var mes message
type pieceTest struct {
	x,y int
}
var rookTests = []pieceTest{
	{1,1},
	{3,3},
	{4,4},
	{5,5},
	{6,6},
	{7,7},
	{8,8},
}


func TestSetPiece(t *testing.T)  {
	interfacelayout := board.InterfaceBoardLayout(board.New())
	interfacelayout.SetBoard(8)
	interfacePiece := InterfacePiece(Rook{})



	for _, v := range rookTests {
		err, mes = interfacePiece.SetPiece(v.x, v.y, interfacelayout)
		//checking for error assigning rooks to specified positions with the given rules
		if err != nil{
			t.Errorf("Error Assiging Cordinates with Messages : %s %s, ", err, mes)
		}
	}

	//algorithm check
	err, mes = interfacePiece.SetPiece(3,2, interfacelayout)
	if err == nil {
		t.Errorf("Error with algorithm, Message : %s ", mes)
	}
	err, mes  = interfacePiece.SetPiece(2,5, interfacelayout)
	if err == nil {
		t.Errorf("Error with algorithm, Message : %s", mes)
	}


}

func TestSetCoordinatesN(t *testing.T) {
	interfacelayout := board.InterfaceBoardLayout(board.New())
	err = interfacelayout.SetBoard(8)
	interfacePiece := InterfacePiece(Rook{})
	if err != nil {
		t.Errorf(err.Error())
	}
	N := 4
	count := 0
	err := interfacePiece.SetCoordinatesN(N, interfacelayout)
	if err != nil {
		t.Errorf(err.Error())
	}
	//check if the numbers of assigned rooks is equal to requested rooks (N)
	for _, v := range interfacelayout.GetBoard(){
		for _, k := range v{
			if k == 1{
				count++
			}
		}
	}
	if count != N {
		t.Errorf("Numbers of rooks found on board is greater or less than requested N. Request rooks N : %d Found rooks = %d", N,count)
	}
}

