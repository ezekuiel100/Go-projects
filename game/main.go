package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	posX := 0
	posY := 0

	rl.InitWindow(0, 0, "Meu jogo Raylib em Go")
	rl.ToggleFullscreen() // Alterna para modo fullscreen

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyRight) {
			posX += 5
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			posX -= 5
		}

		if rl.IsKeyDown(rl.KeyUp) {
			posY -= 5
		}
		if rl.IsKeyDown(rl.KeyDown) {
			posY += 5
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(int32(posX), int32(posY), 50, 50, rl.Green)

		rl.EndDrawing()
	}
}
