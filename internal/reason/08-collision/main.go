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
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Collision")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawFPS(0, 0)

		mouse := rl.GetMousePosition()

		rec1 := rl.NewRectangle(500, 500, 100, 100)
		rec2 := rl.NewRectangle(float32(mouse.X-50), float32(mouse.Y-50), 100, 100)

		rl.DrawRectangleRec(rec1, rl.Blue)
		rl.DrawRectangleRec(rec2, rl.Red)

		if rl.CheckCollisionRecs(rec1, rec2) {
			text := "Collision Detected!"
			fontSize := float32(50)
			fontSpacing := float32(10)
			textVect := rl.MeasureTextEx(rl.GetFontDefault(), text, fontSize, fontSpacing)
			rl.DrawTextPro(rl.GetFontDefault(), text, rl.NewVector2(screenWidth/2, 100), rl.NewVector2(textVect.X/2, textVect.Y/2), 0, fontSize, fontSpacing, rl.Red)
		}

		rl.EndDrawing()
	}
}
