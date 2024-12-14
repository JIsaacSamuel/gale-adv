package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

func Exit() {
	rl.UnloadTexture(GrassTile)
	rl.UnloadTexture(PlayerIdle)
	rl.UnloadTexture(Player_right)
	rl.UnloadTexture(Player_left)
	rl.UnloadTexture(Player_Jump)

	rl.CloseWindow()
}
