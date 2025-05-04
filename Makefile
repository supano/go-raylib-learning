.ONESHELL:

.DEFAULT_GOAL:
	run

run: 
	go run internal/main.go

build: 
	go build -o dist/my-first-game.exe internal/main.go
	