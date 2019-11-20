package game

import (
	"errors"

	tl "github.com/JoelOtter/termloop"
)

const fps = 30

var (
	player *Player
	// Game is the TermLoop game instance
	Game *tl.Game
)

// NewGame is a factory function that constructs a new game instance
func NewGame() (g *tl.Game, err error) {
	if Game != nil {
		return Game, errors.New("an instance of game already exists")

	}

	Game = tl.NewGame()
	Game.Screen().SetFps(fps)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))

	if player == nil {
		p, e := InitPlayer(level, tl.ColorRed)
		if e != nil {
			return Game, errors.New("failed to add player to game")
		}
		player = p
	}

	Game.Screen().SetLevel(level)

	return Game, nil
}
