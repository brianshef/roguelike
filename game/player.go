package game

import tl "github.com/JoelOtter/termloop"

const char = '@'

// Player defines the entity controlled by the user
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

// Draw defines the behavior for drawing the Player on the screen
func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

// Tick defines the core logic of the Player
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
	}
}

// Collide defines the behavior for the Player collisions
func (player *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
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
