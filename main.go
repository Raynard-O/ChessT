package main

import (
	"Tesla_Vehicle_Software/board"
	"Tesla_Vehicle_Software/piece"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	interfacelayout := board.InterfaceBoardLayout(board.New())
	interfacePiece := piece.InterfacePiece(piece.Rook{})
	reader := bufio.NewReader(os.Stdin)

loop:
	fmt.Print("Enter Board Size: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		goto loop
	}
	in, err := strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		fmt.Println(err.Error())
		goto loop
	}
	err = interfacelayout.SetBoard(in)
	if err != nil {
		fmt.Println(err.Error())
		goto loop
	}
loop2:
	fmt.Print("Enter Amount of Preset Rooks: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	in, err = strconv.Atoi(strings.Trim(input, "\n"))
	if err != nil {
		fmt.Println(err.Error())
		goto loop2
	}
	err = interfacePiece.SetCoordinatesN(in, interfacelayout)
	if err != nil {
		fmt.Println(err.Error())
		goto loop2
	}
	interfacelayout.PrintBoard()

	// normal play
	scanner := bufio.NewScanner(os.Stdin)
	input = ""
	rooksPlaced := len(interfacelayout.GetBoardLayoutData().XCoordinates)
	freePlacement := interfacelayout.GetBoardLayoutData().Width - rooksPlaced

	for input != "End" && freePlacement != 0 {
		fmt.Print("Enter Command.. (i.e Add to add Rooks, Print to print board, End to End game): ")
		scanner.Scan()
		input = scanner.Text()

		switch input {
		//Add
		//Add rook to board
		case "Add":
			fmt.Print("Enter Rook's Coordinates: ")
			scanner.Scan()
			input = scanner.Text()
			splitInput := strings.Split(input, ",")
			x, _ := strconv.Atoi(splitInput[0])
			y, _ := strconv.Atoi(splitInput[1])

			err, mes  := interfacePiece.SetPiece(x, y, interfacelayout)

			if err != nil {
				fmt.Println( err.Error(), mes)
				continue
			}
			fmt.Printf("Rook placed at Point at X: %d and Y: %d\n", x, y)
			interfacelayout.PrintBoard()

		//Print
		// print board layout
		case "Print":
			fmt.Println(" ")
			interfacelayout.PrintBoard()
			fmt.Println(" ")
			//End Program
		case "END":

		}
		rooksPlaced = len(interfacelayout.GetBoardLayoutData().XCoordinates)
		freePlacement = interfacelayout.GetBoardLayoutData().Width - rooksPlaced
	}

	fmt.Println("Game End ....")
	interfacelayout.PrintBoard()

}
