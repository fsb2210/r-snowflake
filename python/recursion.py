"""
Snowflake drawing

Requirements:
    - raylib (python3 -m pip install raylib)
"""

from math import pi, cos, sin
from raylib import ffi, rl, colors

screen_factor: int = 80
width: int = 16 * screen_factor
heigth: int = 9 * screen_factor

branches: int = 5
branch_len: int = 2 * screen_factor
branch_thickness: float = 10.0
branch_angle: float = 2 * pi / branches
branch_level: int = 6

def draw_snowflake(center: list, level: int, length: int, thickness: float, hue: float) -> None:
    """Recursively draw a snowflake
    """
    if level <= 0: return
    col = rl.ColorFromHSV(hue, 1.0, 1.0)
    for i in range(branches):
        branch = [center[0]+cos(branch_angle*i)*length, center[1]+sin(branch_angle*i)*length]
        rl.DrawLineEx(center, branch, thickness, col)
        draw_snowflake(branch, level-1, length*0.5, thickness*0.5, hue+70.0)

if __name__ == "__main__":

    rl.InitWindow(width, heigth, b"Copo de nieve")
    rl.SetTargetFPS(60)

    while not rl.WindowShouldClose():
        rl.BeginDrawing()
        rl.ClearBackground(colors.BLACK)
        center = [rl.GetScreenWidth()*0.5, rl.GetScreenHeight()*0.5]
        draw_snowflake(center, branch_level, branch_len, branch_thickness, 0.0)
        rl.EndDrawing()

    rl.CloseWindow()
