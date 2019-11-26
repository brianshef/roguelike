package ui

import (
	"fmt"

	"github.com/brianshef/roguelike/player"

	tl "github.com/JoelOtter/termloop"
)

// PlayerInfo is a simple data struct to hold text info about the player
type PlayerInfo struct {
	*tl.Text
	player *player.Player
}

// NewPlayerInfo is a factory function to construct a new PlayerInfo instance
// along with its panel ( background rectangle )
func NewPlayerInfo(p *player.Player) (info *PlayerInfo, panel *tl.Rectangle) {
	msg := fmt.Sprintf("Roguelike by Brian Shef")
	panel = tl.NewRectangle(0, 0, 100, 3, tl.ColorBlack)
	info = &PlayerInfo{tl.NewText(2, 1, msg, tl.ColorBlue, tl.ColorBlack), p}
	return
}

// Tick defines the game logic of the info
func (info *PlayerInfo) Tick(ev tl.Event) {
	info.SetText(info.player.Status())
}
