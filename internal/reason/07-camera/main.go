package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/supano/raylib/internal/resources"
	"github.com/supano/raylib/internal/util"
)

const (
	screenWidth          = 1920
	screenHeight         = 1080
	playerSpeed  float32 = 4
)

var (
	playerTexture                                         rl.Texture2D
	playerSrc, playerDest                                 rl.Rectangle
	isPlayerMoving                                        = false
	playerGoUp, playerGoDown, playerGoLeft, playerGoRight bool
	frameCount, playerFrameCount                          int64
	camera                                                rl.Camera2D
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Camera")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	playerTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", resources.PlayerTextureFile, int32(len(resources.PlayerTextureFile))))
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(0, 0, 200, 200)

	camera = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2),
		rl.NewVector2(playerDest.X, playerDest.Y),
		0,
		2,
	)
}

func close() {
	rl.UnloadTexture(playerTexture)

	rl.CloseWindow()
}

func input() {
	isPlayerMoving = false
	playerGoUp = false
	playerGoDown = false
	playerGoLeft = false
	playerGoRight = false

	if rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
		isPlayerMoving = true
		playerGoDown = true
	}

	if rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
		isPlayerMoving = true
		playerGoUp = true
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
		isPlayerMoving = true
		playerGoLeft = true
	}

	if rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
		isPlayerMoving = true
		playerGoRight = true
	}
}

func update() {
	frameCount++

	if isPlayerMoving {
		if frameCount%10 == 0 {
			playerFrameCount++
		}

		if playerGoUp {
			playerSrc.X = playerSrc.Width * float32(playerFrameCount)
			playerSrc.Y = playerSrc.Height * 1
		}

		if playerGoDown {
			playerSrc.X = playerSrc.Width * float32(playerFrameCount)
			playerSrc.Y = playerSrc.Height * 0
		}

		if playerGoLeft {
			playerSrc.X = playerSrc.Width * float32(playerFrameCount)
			playerSrc.Y = playerSrc.Height * 2
		}

		if playerGoRight {
			playerSrc.X = playerSrc.Width * float32(playerFrameCount)
			playerSrc.Y = playerSrc.Height * 3
		}
	} else {
		playerSrc.X = 0
	}

	if playerFrameCount > 3 {
		playerFrameCount = 0
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawFPS(0, 0)
	rl.BeginMode2D(camera)

	util.DrawLineAtZeroZero()
	origin := rl.NewVector2(playerDest.Width/2, playerDest.Height/2)
	rl.DrawTexturePro(playerTexture, playerSrc, playerDest, origin, 0, rl.White)

	camera.Target = rl.NewVector2(playerDest.X, playerDest.Y)

	rl.EndMode2D()
	rl.EndDrawing()
}

func main() {
	defer close()

	for !rl.WindowShouldClose() {
		input()
		update()
		render()
	}
}
