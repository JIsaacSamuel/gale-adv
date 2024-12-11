package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawScene() {
	rl.DrawTexturePro(GrassTile, Grass_Src, Grass_Dest, rl.NewVector2(0, 0), 0, rl.White)
	// rl.DrawTexture(GrassTile, 90, 90, rl.White)
	rl.DrawTexturePro(PlayerIdle, Player_Src, Player_Dest, rl.NewVector2(Player_Dest.Height/2, Player_Dest.Width/2), 0, rl.White)
}
