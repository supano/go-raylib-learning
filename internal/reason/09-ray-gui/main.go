package main

import (
	_ "embed"

	rgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Ray GUI")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	// rgui.SetStyle(rgui.DEFAULT, rgui.TEXT_SIZE, 10)

	todos := []string{}

	text := ""

	rgui.SetStyle(rgui.DEFAULT, rgui.TEXT_SIZE, 50)
	rgui.SetStyle(rgui.DEFAULT, rgui.TEXT_SPACING, 3)

	padding := 100

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawFPS(0, 0)

		rgui.Label(rl.NewRectangle(100, 100, 500, 100), "Hello, Ray GUI!")
		rgui.MessageBox(rl.NewRectangle(100, 300, 500, 100), "Message Box", "body", "ok")

		_ = rgui.TextBox(rl.NewRectangle(100, 500, 500, 100), &text, 30, true)
		isButtonClick := rgui.Button(rl.NewRectangle(650, 500, 150, 100), "Add")
		if isButtonClick {
			todos = append(todos, text)
			text = ""
		}

		for i, todo := range todos {
			rl.DrawText(todo, screenWidth/2, int32((100*i)+padding), 50, rl.White)
			onClick := rgui.Button(rl.NewRectangle(screenWidth-250, float32((100*i)+padding), 200, 100), "Remove")
			if onClick {
				todos = append(todos[:i], todos[i+1:]...)
			}
		}

		rl.EndDrawing()
	}
}
