package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	timer float32 = 0.0
	frame int     = 0
)

func draw_ground() {
	ground_temp_dest := Grass_Dest
	var i float32
	for i = 0; i < 1376; i += 64 {
		rl.DrawTexturePro(GrassTile, Grass_Src, ground_temp_dest, rl.NewVector2(0, 0), 0, rl.White)
		ground_temp_dest.X += 92
	}
}

func idle_player() {
	timer += rl.GetFrameTime()

	if timer >= 0.1 {
		timer = 0.0
		frame += 1
	}

	frame %= 3
	Player_temp_Src := Player_Src
	Player_temp_Src.X += float32(frame * 180)

	rl.DrawTexturePro(PlayerIdle, Player_temp_Src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func DrawScene() {
	draw_ground()
	idle_player()
}
