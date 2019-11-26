package player

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

const char = '웃'

// Player defines the entity controlled by the user
type Player struct {
	*tl.Entity
	level       *tl.BaseLevel
	isColliding bool
	prevX       int
	prevY       int
	tick        int
}

// Draw defines the behavior for drawing the Player on the screen
func (player *Player) Draw(screen *tl.Screen) {
	w, h := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(w/2-x, h/2-y)
	player.Entity.Draw(screen)
}

// Tick defines the core logic of the Player
func (player *Player) Tick(event tl.Event) {
	player.tick++
	if event.Type == tl.EventKey {
		player.isColliding = false
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
	}
}

// Collide defines the behavior for the Player collisions
func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.isColliding = true
		player.SetPosition(player.prevX, player.prevY)
		return
	}
}

// Status returns a string with critical status information about the Player instance
func (player *Player) Status(a ...interface{}) (status string) {
	x, y := player.Position()

	colliding := ""
	if player.isColliding {
		colliding = "colliding "
	}

	extra := ""
	if len(a) > 0 {
		extra = fmt.Sprintf("%v", a)
	}

	status = fmt.Sprintf(
		"%v: [ player @ (%v,%v) %v] %v",
		player.tick,
		x, y, colliding,
		extra,
	)

	return
}

// InitPlayer adds the Player to the level
func InitPlayer(level *tl.BaseLevel, color tl.Attr) (player *Player, err error) {
	p := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}

	p.SetCell(0, 0, &tl.Cell{Fg: color, Ch: char})
	level.AddEntity(&p)
	return &p, nil
}
