package models

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
