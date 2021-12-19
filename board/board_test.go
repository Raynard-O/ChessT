package board


import "testing"

type boardTest struct {
	width int
}
type rookTest struct {
	x,y int
}


var boardTests = []boardTest{
	{8},
	{4},
	{16},
}



var rookTests = []rookTest{
	{1,1},
	{2,2},
	{3, 3},
	{4,4},
	{5,5},
	{6,6},
	{7,7},
	{8,8},
}



func TestNew(t *testing.T)  {
	newBoard := New()
	if newBoard == nil{
		t.Errorf("New board not instantiated, Values: %q ", newBoard)
	}
}



func TestSetBoard(t *testing.T) {
	Intlayout := InterfaceBoardLayout(New())

	for _, v := range boardTests {
		err := Intlayout.SetBoard(v.width)
		if err != nil {
			t.Errorf("Board details are not assigend correctly. Board width %d Explected Width %d\n, Board Height %d Explected Height %d", Intlayout.GetBoardLayoutData().Width, v.width, Intlayout.GetBoardLayoutData().Height, v.width)
		}
	}
}

func TestGetBoardLayoutData(t *testing.T) {
	Intlayout := InterfaceBoardLayout(New())

	_ = Intlayout.SetBoard(8)
	board := Intlayout.GetBoardLayoutData()

	if board == nil{
		t.Errorf("Board object is empty")
	}
}

func TestSetCoordinates(t *testing.T) {
	Layout := InterfaceBoardLayout(New())

	_ = Layout.SetBoard(8)

	for _, v := range rookTests {
		Layout.SetCoordinates(v.x, v.y)
		if Layout.GetBoard()[v.x-1][v.y-1] != 1 {
			t.Errorf("Rooks not placed in the right cordinates, cordinates X : %d, Y: %d, Expected : 1, Got : %d", v.x, v.y, Layout.GetBoard()[v.x][v.y])
		}
	}
}


