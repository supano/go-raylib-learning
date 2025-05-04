package main

import (
	_ "embed"
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth          = 1920
	screenHeight         = 1080
	playerSpeed  float32 = 4
)

var (
	//go:embed resources/bgm/relaxed-piano-bgm-242462.mp3
	bgmFile []byte

	//go:embed resources/tilesets/grass.png
	grassSpriteFile []byte

	//go:embed resources/characters/basic-charakter-spritesheet.png
	playerSpriteFile []byte
)

var (
	grassTileSheet, playerSprite rl.Texture2D
	playerSrc, playerDest        rl.Rectangle
	bgm                          rl.Music
	camera                       rl.Camera2D

	isWindowOpen                                          = true
	isPlayerMoving                                        = false
	playerGoUp, playerGoDown, playerGoLeft, playerGoRight bool
	backgroundColor                                       = rl.NewColor(161, 212, 195, 255)
	frameCount, playerFrameCount                          int32
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example")
	rl.SetTargetFPS(60)
	//rl.SetExitKey(0)       // disable esc to close window
	rl.SetWindowMonitor(0) // open on primary monitor
	rl.InitAudioDevice()

	loadResources()

	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(0, 0, 100, 100)

	camera = rl.NewCamera2D(
		rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))),
		0,
		2,
	)
}

func loadResources() {
	grassTileSheet = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", grassSpriteFile, int32(len(grassSpriteFile))))
	playerSprite = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", playerSpriteFile, int32(len(playerSpriteFile))))
	bgm = rl.LoadMusicStreamFromMemory(".mp3", bgmFile, int32(len(bgmFile)))
	rl.SetMusicVolume(bgm, 0.2)
}

func close() {
	rl.UnloadTexture(grassTileSheet)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(bgm)

	rl.CloseAudioDevice()
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
	isWindowOpen = !rl.WindowShouldClose()
	rl.UpdateMusicStream(bgm)
	camera.Target = rl.NewVector2(float32(playerDest.X), float32(playerDest.Y))
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

const (
	srcTileWidth       int = 16 // how many pixels in a tile from source image
	srcTileHeight      int = 16 // how many pixels in a tile from source image
	srcTileSheetWidth  int = 176
	srcTileSheetHeight int = 112
	srcTileColumn      int = srcTileSheetWidth / srcTileWidth

	mapColumn      int = 6  // how may tiles in a row
	mapRow         int = 8  // how many tiles in a column
	destTileWidth  int = 32 // how many pixels it will draw on screen per tile
	destTileHeight int = 32 // how many pixels it will draw on screen per tile
)

var (
	defaultLevel = []int{
		01, 02, 02, 02, 02, 03,
		12, 13, 13, 13, 13, 14,
		12, 13, 13, 13, 13, 14,
		12, 13, 13, 13, 13, 14,
		12, 13, 13, 13, 13, 14,
		12, 13, 13, 13, 13, 14,
		12, 13, 13, 13, 13, 14,
		23, 24, 24, 24, 24, 25,
	}
)

func getTile(mapData []int, row, col int) int {
	return mapData[row*mapColumn+col]
}

func drawScene() {
	for row := range mapRow {
		for col := range mapColumn {
			tile := getTile(defaultLevel, row, col)

			// srcX := float32(((tile - 1) * int(srcTileWidth)) % int(srcTileSheetWidth))
			// srcY := float32(math.Floor(float64(tile-1)/float64(srcTileColumn)) * float64(srcTileHeight))

			srcX := ((tile - 1) % srcTileColumn) * srcTileWidth
			srcY := ((tile - 1) / srcTileColumn) * srcTileHeight

			tileSrc := rl.NewRectangle(
				float32(srcX),
				float32(srcY),
				float32(srcTileWidth),
				float32(srcTileHeight),
			)

			destX := (-destTileWidth * mapColumn / 2) + (col * destTileWidth)
			destY := (-destTileHeight * mapRow / 2) + row*destTileHeight

			tileDest := rl.NewRectangle(
				float32(destX),
				float32(destY),
				float32(destTileWidth),
				float32(destTileHeight),
			)

			rl.DrawTexturePro(grassTileSheet, tileSrc, tileDest, rl.NewVector2(0, 0), 0, rl.White)
			// rl.DrawText(fmt.Sprintf("%d", tile), int32(col*int(destTileWidth)), int32(row*int(destTileWidth)), 10, rl.White)
			// rl.DrawText(fmt.Sprintf("x:%d\ny:%d", int(srcX), int(srcY)), int32(col*int(destTileWidth)), int32(row*int(destTileWidth))+10, 3, rl.Black)
			// rl.DrawRectangleLines(int32(destX), int32(destY), int32(destTileWidth), int32(destTileHeight), rl.Red)
		}
	}

	// rl.DrawLine(-screenHeight, 0, screenHeight, 0, rl.Red)
	// rl.DrawLine(0, -screenWidth, 0, screenWidth, rl.Red)

	// rl.DrawTexture(grassSprite, 0, 0, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width/2, playerDest.Height/2), 0, rl.White)

	// player outline
	// rl.DrawRectangleLines(int32(playerDest.X), int32(playerDest.Y), int32(playerDest.Width), int32(playerDest.Height), rl.Red)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	rl.DrawFPS(0, 0)
	rl.BeginMode2D(camera)

	drawScene()

	rl.EndMode2D()
	rl.EndDrawing()
}

func main() {
	defer close()

	rl.PlayMusicStream(bgm)

	go counter()

	for isWindowOpen {
		input()
		update()
		render()
	}
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println("counter:", i)
		time.Sleep(time.Second)
	}
}
