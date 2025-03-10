package models

import "fmt"

type Game struct {
	PlayingBoard *Board
	Players      []Player
}

func NewGame(b *Board) *Game {
	var player []Player
	return &Game{
		PlayingBoard: b,
		Players:      player,
	}
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}

func (g *Game) StartGame() {
	var noWinner bool = true
	playerTurnIndex := 0
	for noWinner {
		var row, column int
		fmt.Printf("Enter row and column")
		fmt.Scan(&row, &column)
		g.PlayingBoard.AddPiece(g.Players[playerTurnIndex], row, column)
		fmt.Println(g.PlayingBoard.PlayingBoard)
		isWinner := g.isThereWinner(row, column, g.Players[playerTurnIndex].Piece)
		if isWinner {
			noWinner = false
			fmt.Printf("game won by %s", g.Players[playerTurnIndex].Name)
		}
		playerTurnIndex = 1 - playerTurnIndex
	}
}
func (g *Game) isThereWinner(row, col int, p PlayingPiece) bool {
	for i := 0; i < g.PlayingBoard.Size; i++ {
		if g.PlayingBoard.PlayingBoard[row][i] != p {
			return false
		}
	}

	// for i := 0; i < g.PlayingBoard.Size; i++ {
	// 	if g.PlayingBoard.PlayingBoard[i][col] != p {
	// 		return false
	// 	}
	// }

	return true

}
