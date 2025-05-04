package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Game Loop")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawFPS(0, 0)

		rl.ClearBackground(rl.White)
		rl.DrawText("Hello, Raylib!", 100, 100, 100, rl.Black)

		rl.EndDrawing()
	}
}
