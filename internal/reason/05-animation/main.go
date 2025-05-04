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

	playerFramePerAnimation = 10
	// playerMsPerPerAnimation float32 = 1.0 / 4
)

var (
	playerTexture                                         rl.Texture2D
	playerSrc, playerDest                                 rl.Rectangle
	isPlayerMoving                                        = false
	playerGoUp, playerGoDown, playerGoLeft, playerGoRight bool
	frameCount, playerAnimationFrame                      int64

	// timePassed float32 = 0
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Animation")
	rl.SetTargetFPS(144)
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
	// timePassed += rl.GetFrameTime()

	if isPlayerMoving {
		if frameCount%playerFramePerAnimation == 0 {
			playerAnimationFrame++
		}

		// if timePassed > playerMsPerPerAnimation {
		// 	playerAnimationFrame++
		// 	timePassed = 0
		// }

		if playerGoUp {
			playerSrc.X = playerSrc.Width * float32(playerAnimationFrame)
			playerSrc.Y = playerSrc.Height * 1
		}

		if playerGoDown {
			playerSrc.X = playerSrc.Width * float32(playerAnimationFrame)
			playerSrc.Y = playerSrc.Height * 0
		}

		if playerGoLeft {
			playerSrc.X = playerSrc.Width * float32(playerAnimationFrame)
			playerSrc.Y = playerSrc.Height * 2
		}

		if playerGoRight {
			playerSrc.X = playerSrc.Width * float32(playerAnimationFrame)
			playerSrc.Y = playerSrc.Height * 3
		}
	} else {
		playerSrc.X = 0
	}

	if playerAnimationFrame > 3 {
		playerAnimationFrame = 0
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
		update()
		render()
	}
}
