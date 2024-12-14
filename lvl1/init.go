package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WinHeight = 600
	WinWidth  = 1440
)

var (
	GrassTile    rl.Texture2D
	PlayerIdle   rl.Texture2D
	Player_right rl.Texture2D
	Player_left  rl.Texture2D
	Player_Jump  rl.Texture2D

	Grass_Src  rl.Rectangle
	Grass_Dest rl.Rectangle

	Player_Src  rl.Rectangle
	Player_Dest rl.Rectangle

	Player_right_src  rl.Rectangle
	Player_right_dest rl.Rectangle

	Player_left_src  rl.Rectangle
	Player_left_dest rl.Rectangle

	Player_Jump_src rl.Rectangle
)

func Init() {
	rl.InitWindow(WinWidth, WinHeight, "Adventure of Gales")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	Initialise_player()

	GrassTile = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Texture/TX Tileset Ground.png")
	PlayerIdle = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Hero Knight/Sprites/Idle.png")
	Player_right = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Hero Knight/Sprites/Run.png")
	Player_left = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Hero Knight/Sprites/Run-inverted.png")
	Player_Jump = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Hero Knight/Sprites/Jump.png")

	// Player rectangles
	Player_Src = rl.NewRectangle(60, 60, 60, 60)
	Player_right_src = rl.NewRectangle(60, 60, 60, 60)
	Player_left_src = rl.NewRectangle(60, 60, 60, 60)
	Player_Jump_src = rl.NewRectangle(60, 60, 60, 60)

	Player_Dest = rl.NewRectangle(0, 478, 60, 60)

	// Ground tiles
	Grass_Src = rl.NewRectangle(0, 384, 60, 30)
	Grass_Dest = rl.NewRectangle(0, 536, 96, 64)
}
