package lvl1

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WinHeight = 600
	WinWidth  = 1440
)

var (
	GrassTile  rl.Texture2D
	PlayerIdle rl.Texture2D

	Grass_Src  rl.Rectangle
	Grass_Dest rl.Rectangle

	Player_Src  rl.Rectangle
	Player_Dest rl.Rectangle
)

func Init() {
	rl.InitWindow(WinWidth, WinHeight, "Adventure of Gales")
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)

	GrassTile = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Mossy_assets/Mossy - FloatingPlatforms.png")
	PlayerIdle = rl.LoadTexture("/home/isaac/github.com/JIsaacSamuel/gale-adv/Assets/Hero Knight/Sprites/Idle.png")

	Player_Src = rl.NewRectangle(0, 0, 180, 180)
	Player_Dest = rl.NewRectangle(90, 90, 360, 360)

	Grass_Src = rl.NewRectangle(0, 1536, 2048, 512)
	Grass_Dest = rl.NewRectangle(10, 500, 512, 128)
}
