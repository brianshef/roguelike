package game

import tl "github.com/JoelOtter/termloop"

// GenerateLevel returns a new level
func GenerateLevel() (level *tl.BaseLevel) {
	level = tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: '.',
	})

	// Environment
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorWhite))
	return
}
