package models

type Game struct {
	PlayingBoard Board
	Players      []Player
}

func NewGame(b Board) *Game {
	var player []Player
	return &Game{
		PlayingBoard: b,
		Players:      player,
	}
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}
