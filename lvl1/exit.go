package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

func Exit() {
	rl.UnloadTexture(GrassTile)
	rl.UnloadTexture(PlayerIdle)
	rl.CloseWindow()
}
