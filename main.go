package main

import (
	roguelike "github.com/brianshef/roguelike/game"
	"github.com/brianshef/roguelike/logger"
)

var logs *logger.Log

func main() {
	logs = logger.NewLoggers()

	game, err := roguelike.NewGame()
	if err != nil {
		logs.Error.Panicln(err)
	}

	game.Start()
}
