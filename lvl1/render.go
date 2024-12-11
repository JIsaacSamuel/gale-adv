package lvl1

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.NewColor(1, 63, 69, 1))

	DrawScene()

	rl.EndDrawing()
}
