package main

import (
	_ "embed"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Raylib Example - Timer")
	rl.SetTargetFPS(60)
	rl.SetWindowMonitor(0) // open on primary monitor

	timer := NewStopWatch()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawFPS(screenWidth-100, 0)

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			timer.Start()
		} else {
			timer.Stop()
		}

		timer.Update()

		rl.DrawText(fmt.Sprintf("%f", rl.GetTime()), 10, 10, 50, rl.Black)
		rl.DrawText(fmt.Sprintf("Mouse Pressed Time: %f", timer.GetTime()), 10, 80, 50, rl.Black)

		rl.EndDrawing()
	}
}

type StopWatch struct {
	start bool
	cur   float32
}

func NewStopWatch() *StopWatch {
	return &StopWatch{}
}

func (s *StopWatch) Start() {
	s.start = true
}

func (s *StopWatch) Stop() {
	s.start = false
}

func (s *StopWatch) Update() {
	if s.start {
		s.cur += rl.GetFrameTime()
	}
}

func (s *StopWatch) Reset() {
	s.cur = 0
}

func (s *StopWatch) GetTime() float32 {
	return s.cur
}
