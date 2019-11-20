package main

import (
	tl "github.com/JoelOtter/termloop"
	rl "github.com/brianshef/roguelike/game"
	"github.com/brianshef/roguelike/logger"
)

const (
	fps = 30
)

var (
	player *rl.Player
	log    *logger.Log
)

// main entry point

func main() {

	log = logger.NewLoggers()

	game := tl.NewGame()
	game.Screen().SetFps(fps)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))

	if player == nil {
		p, e := rl.InitPlayer(level, tl.ColorRed)
		if e != nil {
			log.Error.Println("failed to add player to game")
			panic(e)
		}
		player = p
	}

	game.Screen().SetLevel(level)
	game.Start()
}
