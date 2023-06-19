package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board [3][3]string
var currentPlayerSymbol string

const player1 = "X"
const player2 = "O"
const emptyField = "_"

func main() {
	boardInit()
	switchPlayer()
	run()
}

func run() {
	var x uint
	var y uint

	for {
		display()
		if !boardIsPlayable() {
			fmt.Println("!!!НІЧИЯ!!!!")
			break
		}

		fmt.Printf("Player %s зробіть хід (наприклад, 0 1): ", currentPlayerSymbol)
		_, err := fmt.Scanln(&x, &y)

		if err != nil {
			fmt.Printf("error: %s", err.Error())
		}

		if coordinatesIsValid(x, y) {
			board[x][y] = currentPlayerSymbol
			if isWin() {
				clearDisplay()
				fmt.Printf("!!!Перемога за гравцем %s!!!", currentPlayerSymbol)
				display()
				break
			} else {
				switchPlayer()
				clearDisplay()
			}
		} else {
			clearDisplay()
			fmt.Println("Неправильні координати. Спробуйте ще раз.")
		}
	}
}

func isWin() bool {
	var win bool

	for i := 0; i < len(board); i++ {
		if board[i][0] != emptyField && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			win = true
		}
		if board[0][i] != emptyField && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			win = true
		}
	}

	if board[0][0] != emptyField && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		win = true
	}
	if board[2][0] != emptyField && board[2][0] == board[1][1] && board[2][0] == board[0][2] {
		win = true
	}

	return win
}

func boardIsPlayable() bool {
	var isPlayable bool

	for _, v := range board {
		for _, j := range v {
			if j == emptyField {
				isPlayable = true
				break
			}
		}
	}

	return isPlayable
}

func switchPlayer() {

	if currentPlayerSymbol == player1 {
		currentPlayerSymbol = player2
	} else if currentPlayerSymbol == player2 {
		currentPlayerSymbol = player1
	} else {
		currentPlayerSymbol = player1
	}
}

func coordinatesIsValid(x uint, y uint) bool {
	var check bool

	if (x < 3 && y < 3) && board[x][y] == emptyField {
		check = true
	}

	return check
}

func boardInit() {
	for i := 0; i < len(board); i++ {
		board[i] = [3]string{emptyField, emptyField, emptyField}
	}
}

func display() {
	fmt.Println("   0 1 2")
	for i, val := range board {
		fmt.Println(i, val)
	}
}

func clearDisplay() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
