package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	lvl "github.com/JIsaacSamuel/gale-adv/lvl1"
)

func main() {
	lvl.Init()
	defer lvl.Exit()

	for !rl.WindowShouldClose() {
		lvl.Render()
	}
}
