package models

import "fmt"

type Board struct {
	Size         int
	PlayingBoard [][]PlayingPiece
}

func NewBoard(size int) *Board {
	playingBoard := make([][]PlayingPiece, size)
	for i := range playingBoard {
		playingBoard[i] = make([]PlayingPiece, size)
	}
	return &Board{
		Size:         size,
		PlayingBoard: playingBoard,
	}
}

func (b *Board) AddPiece(player Player, row, col int) {
	if b.PlayingBoard[row][col] != "" {
		fmt.Println("enter some other row and column")
		return
	}

	b.PlayingBoard[row][col] = player.Piece
}
