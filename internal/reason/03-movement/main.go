package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/supano/raylib/internal/resources"
)

const (
	screenWidth          = 1920
	screenHeight         = 1080
	playerSpeed  float32 = 4
)

var (
	playerTexture         rl.Texture2D
	playerSrc, playerDest rl.Rectangle
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Movement")
	// rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	playerTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", resources.PlayerTextureFile, int32(len(resources.PlayerTextureFile))))
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(0, 0, 200, 200)
}

func close() {
	rl.UnloadTexture(playerTexture)

	rl.CloseWindow()
}

func input() {
	if rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawFPS(0, 0)

	origin := rl.NewVector2(0, 0)
	rl.DrawTexturePro(playerTexture, playerSrc, playerDest, origin, 0, rl.White)

	rl.EndDrawing()
}

func main() {
	defer close()

	for !rl.WindowShouldClose() {
		input()
		render()
	}
}
