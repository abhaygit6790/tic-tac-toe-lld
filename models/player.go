package models

type Player struct {
	Name  string
	Piece PlayingPiece
}

func NewPlayer(name string, piece PlayingPiece) *Player {
	return &Player{
		Name:  name,
		Piece: piece,
	}
}
