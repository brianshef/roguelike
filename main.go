package main

import (
	// "fmt"
	tl "github.com/JoelOtter/termloop"
	"io"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func InitLogger(
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Debug = log.New(debugHandle,
		"[DEBUG] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"[WARN] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// Player

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		player.prevX, player.prevY = player.Position()
		switch event.Key {
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
		// Debug.Println(fmt.Printf("player position: %v,%v", player.x, player.y))
	}
}

func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

// game / level / map

func addPlayer(level *tl.BaseLevel, fg tl.Attr) (player *Player, err error) {
	p := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	p.SetCell(0, 0, &tl.Cell{Fg: fg, Ch: '옷'})
	level.AddEntity(&p)
	return &p, nil
}

// main entry point

func main() {
	InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	game := tl.NewGame()
	game.Screen().SetFps(30)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))

	_, err := addPlayer(level, tl.ColorRed)
	if err != nil {
		Error.Println("failed to add player to game")
	}

	game.Screen().SetLevel(level)
	game.Start()
	// Info.Println("game started")
}
