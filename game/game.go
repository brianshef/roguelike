package game

import (
	"errors"

	tl "github.com/JoelOtter/termloop"
)

const fps = 30

var (
	// Game is the TermLoop game instance
	Game   *tl.Game
	player *Player
)

// NewGame is a factory function that constructs a new game instance
func NewGame() (g *tl.Game, err error) {
	if Game != nil {
		return Game, errors.New("an instance of game already exists")

	}

	g = tl.NewGame()
	g.Screen().SetFps(fps)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))

	if player == nil {
		p, err := InitPlayer(level, tl.ColorRed)
		if err != nil {
			return g, errors.New("failed to add player to game")
		}
		player = p
	}

	g.Screen().AddEntity(tl.NewRectangle(0, 0, 100, 3, tl.ColorBlack))
	playerInfo := NewPlayerInfo(player)
	g.Screen().AddEntity(playerInfo)

	g.Screen().SetLevel(level)

	return
}
