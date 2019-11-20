package main

import (
	rl "github.com/brianshef/roguelike/game"
	// "fmt"
	"io"
	"log"
	"os"

	tl "github.com/JoelOtter/termloop"
)

const (
	logOpts = log.Ldate | log.Ltime | log.Lshortfile
	fps     = 30
)

var (
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger

	player *rl.Player
)

func initLogger(
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	debug = log.New(debugHandle, "[DEBUG] ", logOpts)
	info = log.New(infoHandle, "[INFO] ", logOpts)
	warning = log.New(warningHandle, "[WARN] ", logOpts)
	err = log.New(errorHandle, "[ERROR] ", logOpts)
}

// main entry point

func main() {
	initLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

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
			err.Println("failed to add player to game")
			panic(e)
		}
		player = p
	}

	game.Screen().SetLevel(level)
	game.Start()
}
