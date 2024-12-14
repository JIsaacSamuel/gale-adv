package lvl1

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	timer         float32 = 0.0
	frame         int     = 0
	lframe        int     = 7
	up            bool    = true
	jump_velocity int     = 25
	gravity       int     = -1

	// jump_y float32 = 0.0
)

func draw_ground() {
	ground_temp_dest := Grass_Dest
	var i float32
	for i = 0; i < 1376; i += 64 {
		rl.DrawTexturePro(GrassTile, Grass_Src, ground_temp_dest, rl.NewVector2(0, 0), 0, rl.White)
		ground_temp_dest.X += 92
	}
}

func jump() {
	var Player_temp_src rl.Rectangle

	if up {
		if jump_velocity > 0 {
			timer += rl.GetFrameTime()

			if timer >= 0.2 {
				timer = 0
				frame++
			}

			frame %= 3
			Player_temp_src = Player_Jump_src
			Player_temp_src.X += float32(frame * 180)
			Player_Dest.Y -= float32(jump_velocity + gravity)
			jump_velocity += gravity
			if User.right {
				Player_Dest.X += User.h_speed
			} else if User.left {
				Player_Dest.X -= User.h_speed
			}

			rl.DrawTexturePro(Player_Jump, Player_temp_src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
		} else {
			jump_velocity = 0
			up = false
		}
	} else {
		if jump_velocity > -24 {
			timer += rl.GetFrameTime()

			if timer >= 0.2 {
				timer = 0
				frame++
			}

			frame %= 3
			Player_temp_src = Player_Jump_src
			Player_temp_src.X += float32(frame * 180)
			Player_Dest.Y -= float32(jump_velocity + gravity)
			jump_velocity += gravity
			if User.right {
				Player_Dest.X += User.h_speed
			} else if User.left {
				Player_Dest.X -= User.h_speed
			}

			rl.DrawTexturePro(Player_Jump, Player_temp_src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
		} else {
			jump_velocity = 25
			up = true
		}
	}
}

func run_left() {
	jump_velocity = 25
	timer += rl.GetFrameTime()

	if timer >= 0.1 {
		timer = 0.0
		lframe--
	}

	if lframe < 0 {
		lframe = 7
	}
	Player_temp_Src := Player_left_src
	Player_temp_Src.X += float32(lframe * 180)

	Player_Dest.X -= User.h_speed

	rl.DrawTexturePro(Player_left, Player_temp_Src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func run_right() {
	jump_velocity = 25
	timer += rl.GetFrameTime()

	if timer >= 0.1 {
		timer = 0.0
		frame += 1
	}

	frame %= 8
	Player_temp_Src := Player_right_src
	Player_temp_Src.X += float32(frame * 180)

	Player_Dest.X += User.h_speed

	rl.DrawTexturePro(Player_right, Player_temp_Src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func idle_player() {
	jump_velocity = 25
	timer += rl.GetFrameTime()

	if timer >= 0.1 {
		timer = 0.0
		frame += 1
	}

	frame %= 3
	Player_temp_Src := Player_Src
	Player_temp_Src.X += float32(frame * 180)
	Player_Dest.Y = 478

	rl.DrawTexturePro(PlayerIdle, Player_temp_Src, Player_Dest, rl.NewVector2(0, 0), 0, rl.White)
}

func DrawScene() {
	draw_ground()
	if User.jump {
		jump()
	} else if User.left {
		run_left()
	} else if User.right {
		run_right()
	} else {
		idle_player()
	}
}
