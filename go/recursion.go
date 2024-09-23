package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

const (
	screenFactor int32 = 80
	width        int32 = 16 * screenFactor
	height       int32 = 9 * screenFactor

	branches         int     = 5
	branch_len       int32   = 2 * screenFactor
	branch_thickness float32 = 10.0
	branch_angle     float64 = 2 * math.Pi / float64(branches)
	branch_level     int32   = 6
)

func drawSnowflake(center rl.Vector2, level int32, length int32, thickness float32, hue float32) {
	if level <= 0 {
		return
	}
	color := rl.ColorFromHSV(hue, 1.0, 1.0)
	for i := 0; i <= branches-1; i++ {
		branch := rl.Vector2{
			X: center.X + float32(length)*float32(math.Cos(branch_angle*float64(i))),
			Y: center.Y + float32(length)*float32(math.Sin(branch_angle*float64(i))),
		}
		rl.DrawLineEx(center, branch, thickness, color)
		drawSnowflake(branch, level-1, int32(float32(length)*0.5), thickness*0.5, hue+70)
	}
}

func main() {
	rl.InitWindow(width, height, "Copo de nieve")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		center := rl.Vector2{
			X: float32(rl.GetScreenWidth()) * 0.5,
			Y: float32(rl.GetScreenHeight()) * 0.5,
		}
		drawSnowflake(center, branch_level, branch_len, branch_thickness, 0.0)
		rl.EndDrawing()
	}
}
