package main

import (
	"github.com/brianshef/roguelike/game"
	"github.com/brianshef/roguelike/logger"
)

var logs *logger.Log

func main() {
	logs = logger.NewLoggers()

	roguelike, err := game.NewGame()
	if err != nil {
		logs.Error.Panicln(err)
	}

	roguelike.Start()
}
