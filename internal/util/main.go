package util

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawLineAtZeroZero() {
	rl.DrawLine(-100, 0, 100, 0, rl.Red)
	rl.DrawLine(0, -100, 0, 100, rl.Red)
}

func DrawLineAtCenterScreen(screenWidth, screenHeight int32) {
	rl.DrawLine(screenWidth/2, 0, screenWidth/2, screenHeight, rl.Red)
	rl.DrawLine(0, screenHeight/2, screenWidth, screenHeight/2, rl.Red)
}
