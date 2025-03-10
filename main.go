package main

import (
	"abhaygit6790/tic-tac-toe-lld/models"
)

func main() {
	board := models.NewBoard(3)
	game := models.NewGame(board)
	player1 := models.NewPlayer("abhay", models.PlayingPiece("X"))
	player2 := models.NewPlayer("abhay", models.PlayingPiece("0"))
	game.AddPlayer(player1)
	game.AddPlayer(player2)
	game.StartGame()
}
