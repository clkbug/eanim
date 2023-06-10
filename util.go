package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func isKeyLongPressed(k ebiten.Key) bool {
	duration := inpututil.KeyPressDuration(k)
	return inpututil.IsKeyJustPressed(k) || (duration >= 30 && duration%30 == 0) || (duration >= 60 && duration%10 == 0)
}
