package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Game struct {
	gameBoard           [3][3]string
	currentPlayerSymbol string
	resultsBoard        map[string]int
	player1             string
	player2             string
	draw                string
	emptyField          string
}

func main() {
	game := Game{
		gameBoard:           [3][3]string{},
		currentPlayerSymbol: "X",
		resultsBoard:        map[string]int{},
		player1:             "X",
		player2:             "O",
		draw:                "НІЧИЯ",
		emptyField:          "_",
	}
	game.resultsBoardInit()
	game.switchPlayer()
	game.run()
}

func (g *Game) restart() bool {
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
		g.restart()
	}

	return result
}

func (g *Game) run() {
	var x uint
	var y uint

	for {
		g.gameBoardInit()
		fmt.Println("\n\n НОВА ГРА \n ")

		for {
			g.displayGameBoard()
			if !g.gameBoardIsPlayable() {
				fmt.Println("!!!НІЧИЯ!!!!")
				g.addDraw()
				g.displayResults()
				break
			}

			fmt.Printf("Player %s зробіть хід (наприклад, 0 1): ", g.currentPlayerSymbol)
			_, err := fmt.Scanln(&x, &y)

			if err != nil {
				fmt.Printf("error: %s", err.Error())
			}

			if g.coordinatesIsValid(x, y) {
				g.gameBoard[x][y] = g.currentPlayerSymbol
				if g.isWin() {
					g.clearDisplay()
					fmt.Printf("!!!Перемога за гравцем %s!!!", g.currentPlayerSymbol)
					g.addWin()
					g.displayGameBoard()
					g.displayResults()
					break
				} else {
					g.switchPlayer()
					g.clearDisplay()
				}
			} else {
				g.clearDisplay()
				fmt.Println("Неправильні координати. Спробуйте ще раз.")
			}
		}

		if !g.restart() {
			fmt.Println("\n\n Нехай щастить!!!")
			break
		}
	}
}

func (g *Game) isWin() bool {
	var win bool

	for i := 0; i < len(g.gameBoard); i++ {
		if g.gameBoard[i][0] != g.emptyField && g.gameBoard[i][0] == g.gameBoard[i][1] && g.gameBoard[i][0] == g.gameBoard[i][2] {
			win = true
		}
		if g.gameBoard[0][i] != g.emptyField && g.gameBoard[0][i] == g.gameBoard[1][i] && g.gameBoard[0][i] == g.gameBoard[2][i] {
			win = true
		}
	}

	if g.gameBoard[0][0] != g.emptyField && g.gameBoard[0][0] == g.gameBoard[1][1] && g.gameBoard[0][0] == g.gameBoard[2][2] {
		win = true
	}
	if g.gameBoard[2][0] != g.emptyField && g.gameBoard[2][0] == g.gameBoard[1][1] && g.gameBoard[2][0] == g.gameBoard[0][2] {
		win = true
	}

	return win
}

func (g *Game) gameBoardIsPlayable() bool {
	var isPlayable bool

	for _, v := range g.gameBoard {
		for _, j := range v {
			if j == g.emptyField {
				isPlayable = true
				break
			}
		}
	}

	return isPlayable
}

func (g *Game) switchPlayer() {

	if g.currentPlayerSymbol == g.player1 {
		g.currentPlayerSymbol = g.player2
	} else if g.currentPlayerSymbol == g.player2 {
		g.currentPlayerSymbol = g.player1
	} else {
		g.currentPlayerSymbol = g.player1
	}
}

func (g *Game) coordinatesIsValid(x uint, y uint) bool {
	var check bool

	if (x < 3 && y < 3) && g.gameBoard[x][y] == g.emptyField {
		check = true
	}

	return check
}

func (g *Game) gameBoardInit() {
	for i := 0; i < len(g.gameBoard); i++ {
		g.gameBoard[i] = [3]string{g.emptyField, g.emptyField, g.emptyField}
	}
}

func (g *Game) displayGameBoard() {
	fmt.Println("   0 1 2")
	for i, val := range g.gameBoard {
		fmt.Println(i, val)
	}
}

func (g *Game) clearDisplay() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}

func (g *Game) resultsBoardInit() {
	g.resultsBoard = map[string]int{g.player1: 0, g.player2: 0, g.draw: 0}
}

func (g *Game) addWin() {
	g.resultsBoard[g.currentPlayerSymbol] = g.resultsBoard[g.currentPlayerSymbol] + 1
}

func (g *Game) addDraw() {
	g.resultsBoard[g.draw] = g.resultsBoard[g.draw] + 1
}

func (g *Game) displayResults() {
	fmt.Printf("Таблиця результатів: \n Перемоги player %s: %d \n "+
		"Перемоги player %s: %d \n"+
		" %s: %d", g.player1, g.resultsBoard[g.player1], g.player2, g.resultsBoard[g.player2], g.draw, g.resultsBoard[g.draw])
}
