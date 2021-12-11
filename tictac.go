package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Space struct {
	// SPACES CONTAIN AN ID (1-9) AND A SYMBOL (1-9, to be replaced with x or o when the space is chosen)
	id  int
	sym string
}

// USED ONLY FOR SINGLE PLAYER MODE
// computer randomly generates an id number and checks to make sure
// the space with that id has not already been chosen
func compChoice(chosen map[int]bool) int {
	id := rand.Intn(9)

	if chosen[id] {
		id = compChoice(chosen)

		return id
	}

	return id
}

func makeChoice(chosen map[int]bool) int {

	var input string
	fmt.Scanln(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	if chosen[id] {
		fmt.Println("MUST CHOOSE EMPTY SPACE")
		id = makeChoice(chosen)

		return id
	}

	return id
}

func printBoard(board [9]Space) {
	fmt.Printf("\n")
	for i, val := range board {

		fmt.Printf(val.sym)
		if (i+1)%3 != 0 {
			fmt.Printf("|")
		}
		if i == 2 || i == 5 {
			fmt.Printf("\n-----\n")
		}

	}
	fmt.Printf("\n\n********\n")
}

func chooseMode() string {
	var mode string
	fmt.Scanln(&mode)
	if mode != "M" && mode != "S" {
		fmt.Println("MUST ENTER 'M' OR 'S'")
		mode = chooseMode()

		return mode
	}

	return mode
}

func main() {

	fmt.Println("CHOOSE MODE\nFOR SINGLE PLAYER ENTER 'S'\nFOR MULTIPLAYER ENTER 'M'")
	mode := chooseMode()

	board := [9]Space{}
	chosen := make(map[int]bool)

	for i := 0; i < 9; i++ {
		id := i
		sym := strconv.Itoa(i)
		board[i] = Space{id, sym}
		chosen[i] = false
	}

	printBoard(board)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 9; i++ {
		var id int

		if i%2 == 0 {
			id = makeChoice(chosen)
		} else if mode == "S" {
			id = compChoice(chosen)
		} else if mode == "M" {
			id = makeChoice(chosen)
		}
		finished := false

		chosen[id] = true
		for m, val := range board {
			if val.id == id {
				if i%2 == 0 {
					board[m] = Space{id, "x"}
				} else {
					board[m] = Space{id, "o"}
				}
			}
		}

		printBoard(board)

		for i := 0; i < 3; i++ {
			if board[i*3].sym == board[i*3+1].sym && board[i*3].sym == board[i*3+2].sym {
				fmt.Println("GAME OVER")
				fmt.Printf(board[i*3].sym)
				fmt.Printf(" WINS\n")
				finished = true
				break
			}
			if board[i].sym == board[i+3].sym && board[i].sym == board[i+6].sym {
				fmt.Println("GAME OVER")
				fmt.Printf(board[i].sym)
				fmt.Printf(" WINS\n")
				finished = true
				break
			}

		}

		if board[0].sym == board[4].sym && board[0].sym == board[8].sym {
			fmt.Println("GAME OVER")
			fmt.Printf(board[0].sym)
			fmt.Printf(" WINS\n")
			finished = true
			break
		}
		if board[2].sym == board[4].sym && board[2].sym == board[6].sym {
			fmt.Println("GAME OVER")
			fmt.Printf(board[2].sym)
			fmt.Printf(" WINS\n")
			finished = true
			break
		}

		if finished == true {
			break
		}

		if i == 8 && !finished {
			fmt.Println("GAME OVER: TIE")
		}
	}

}
