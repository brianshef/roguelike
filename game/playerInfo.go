package game

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// PlayerInfo is a simple data struct to hold text info about the player
type PlayerInfo struct {
	*tl.Text
	player *Player
}

// NewPlayerInfo is a factory function to construct a new PlayerInfo instance
func NewPlayerInfo(p *Player) *PlayerInfo {
	info := fmt.Sprintf("Roguelike by Brian Shef")
	return &PlayerInfo{tl.NewText(2, 1, info, tl.ColorBlue, tl.ColorBlack), p}
}

// Tick defines the game logic of the info
func (info *PlayerInfo) Tick(ev tl.Event) {
	info.SetText(info.player.Status())
}
