package main

import (
	"fmt"
	"os"
	"os/exec"
)

var gameBoard [3][3]string
var currentPlayerSymbol string
var resultsBoard map[string]int

const player1 = "X"
const player2 = "O"
const draw = "НІЧИЯ"
const emptyField = "_"

func main() {
	resultsBoardInit()
	switchPlayer()
	run()
}

func restart() bool {
	var choice int
	var result bool
	fmt.Println("\n\n Якщо бажаєте зіграти ще натисніть - 1 \n Для завершення натисніть - 0")

	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println(err.Error())
	}

	if choice == 1 {
		result = true
	} else if choice != 0 {
		fmt.Println("Неправильний вибір")
		restart()
	}

	return result
}

func run() {
	var x uint
	var y uint

	for {
		gameBoardInit()
		fmt.Println("\n\n НОВА ГРА \n ")

		for {
			displayGameBoard()
			if !gameBoardIsPlayable() {
				fmt.Println("!!!НІЧИЯ!!!!")
				addDraw()
				displayResults()
				break
			}

			fmt.Printf("Player %s зробіть хід (наприклад, 0 1): ", currentPlayerSymbol)
			_, err := fmt.Scanln(&x, &y)

			if err != nil {
				fmt.Printf("error: %s", err.Error())
			}

			if coordinatesIsValid(x, y) {
				gameBoard[x][y] = currentPlayerSymbol
				if isWin() {
					clearDisplay()
					fmt.Printf("!!!Перемога за гравцем %s!!!", currentPlayerSymbol)
					addWin()
					displayGameBoard()
					displayResults()
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

		if !restart() {
			fmt.Println("\n\n Нехай щастить!!!")
			break
		}
	}
}

func isWin() bool {
	var win bool

	for i := 0; i < len(gameBoard); i++ {
		if gameBoard[i][0] != emptyField && gameBoard[i][0] == gameBoard[i][1] && gameBoard[i][0] == gameBoard[i][2] {
			win = true
		}
		if gameBoard[0][i] != emptyField && gameBoard[0][i] == gameBoard[1][i] && gameBoard[0][i] == gameBoard[2][i] {
			win = true
		}
	}

	if gameBoard[0][0] != emptyField && gameBoard[0][0] == gameBoard[1][1] && gameBoard[0][0] == gameBoard[2][2] {
		win = true
	}
	if gameBoard[2][0] != emptyField && gameBoard[2][0] == gameBoard[1][1] && gameBoard[2][0] == gameBoard[0][2] {
		win = true
	}

	return win
}

func gameBoardIsPlayable() bool {
	var isPlayable bool

	for _, v := range gameBoard {
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

	if (x < 3 && y < 3) && gameBoard[x][y] == emptyField {
		check = true
	}

	return check
}

func gameBoardInit() {
	for i := 0; i < len(gameBoard); i++ {
		gameBoard[i] = [3]string{emptyField, emptyField, emptyField}
	}
}

func displayGameBoard() {
	fmt.Println("   0 1 2")
	for i, val := range gameBoard {
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

func resultsBoardInit() {
	resultsBoard = map[string]int{player1: 0, player2: 0, draw: 0}
}

func addWin() {
	resultsBoard[currentPlayerSymbol] = resultsBoard[currentPlayerSymbol] + 1
}

func addDraw() {
	resultsBoard[draw] = resultsBoard[draw] + 1
}

func displayResults() {
	fmt.Printf("Таблиця результатів: \n Перемоги player %s: %d \n "+
		"Перемоги player %s: %d \n"+
		" %s: %d", player1, resultsBoard[player1], player2, resultsBoard[player2], draw, resultsBoard[draw])
}
