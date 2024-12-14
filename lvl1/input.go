package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

func Input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		User.jump = true
	} else if (rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight)) && Player_Dest.X < WinWidth-180 && !User.jump {
		User.right, User.left, User.jump, User.idle = true, false, false, false
	} else if (rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft)) && Player_Dest.X >= 0 && !User.jump {
		User.right, User.left, User.jump, User.idle = false, true, false, false
	} else {
		User.right, User.left, User.jump, User.idle = false, false, false, true
	}
}
