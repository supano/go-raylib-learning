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
	tileTexture rl.Texture2D
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Tile map")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor
	tileTexture = rl.LoadTextureFromImage(rl.LoadImageFromMemory(".png", resources.GrassTextureFile, int32(len(resources.GrassTextureFile))))
}

func close() {
	rl.UnloadTexture(tileTexture)

	rl.CloseWindow()
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
		/*   1   2   3   4   5   */
		01, 02, 02, 02, 02, 03, // 0
		12, 13, 13, 13, 13, 14, // 1
		12, 13, 13, 13, 13, 14, // 2
		12, 13, 13, 13, 13, 14, // 3
		12, 13, 13, 13, 13, 14, // 4
		12, 13, 13, 13, 13, 14, // 5
		12, 13, 13, 13, 13, 14, // 6
		23, 24, 24, 24, 24, 25, // 7
	}
)

func getTile(mapData []int, row, col int) int {
	return mapData[row*mapColumn+col]
}

func render() {
	rl.BeginDrawing()
	rl.DrawFPS(0, 0)
	rl.ClearBackground(rl.White)

	for row := range mapRow {
		for col := range mapColumn {
			tile := getTile(defaultLevel, row, col)

			srcX := ((tile - 1) % srcTileColumn) * srcTileWidth
			srcY := ((tile - 1) / srcTileColumn) * srcTileHeight

			tileSrc := rl.NewRectangle(
				float32(srcX),
				float32(srcY),
				float32(srcTileWidth),
				float32(srcTileHeight),
			)

			destX := col * destTileWidth
			destY := row * destTileHeight

			tileDest := rl.NewRectangle(
				float32(destX),
				float32(destY),
				float32(destTileWidth),
				float32(destTileHeight),
			)

			rl.DrawTexturePro(tileTexture, tileSrc, tileDest, rl.NewVector2(0, 0), 0, rl.White)
		}
	}

	rl.EndDrawing()
}

func main() {
	defer close()

	for !rl.WindowShouldClose() {
		render()
	}
}
