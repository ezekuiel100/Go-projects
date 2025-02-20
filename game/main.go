package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	posX := 50
	posY := 50

	rl.InitWindow(0, 0, "Meu jogo Raylib em Go")
	rl.ToggleFullscreen() // Alterna para modo fullscreen

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyDown(rl.KeyRight) {
			posX += 5
		} else if rl.IsKeyDown(rl.KeyLeft) {
			posX -= 5
		} else if rl.IsKeyDown(rl.KeyUp) {
			posY -= 5
		} else if rl.IsKeyDown(rl.KeyDown) {
			posY += 5
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		fpsText := fmt.Sprintf("FPS: %d", rl.GetFPS())
		rl.DrawText(fpsText, 20, 20, 20, rl.Yellow)

		rl.DrawRectangle(int32(posX), int32(posY), 50, 50, rl.Green)

		rl.EndDrawing()
	}
}
