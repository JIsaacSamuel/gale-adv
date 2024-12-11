package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

func Input() {
	if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)) && Player_Dest.Y < WinWidth-5 {
		Player_Dest.X += User.h_speed
	}
	if (rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)) && Player_Dest.Y > 5 {
		Player_Dest.X -= User.h_speed
	}
}
