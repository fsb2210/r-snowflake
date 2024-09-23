#include <math.h>
#include "raylib.h"

// window props
#define SCREEN_FACTOR 80
#define WIDTH 16*SCREEN_FACTOR
#define HEIGHT 9*SCREEN_FACTOR
// snowflake branches
#define BRANCHES 5
#define BRANCH_LEN SCREEN_FACTOR*2
#define BRANCH_THICKNESS 10.0
#define BRANCH_ANGLE 2.0*M_PI/BRANCHES
#define BRANCH_LEVEL 6

// recursive snowflake
void drawSnowflake(Vector2 center, int level, int length, float thick, float hue)
{
    if (level <= 0) return;
    Color color = ColorFromHSV(hue, 1.0, 1.0);
    for (int i = 0; i <= BRANCHES-1; i++) {
        Vector2 branch = {center.x + cos(BRANCH_ANGLE*i)*length, center.y + sin(BRANCH_ANGLE*i)*length};
        DrawLineEx(center, branch, thick, color);
        drawSnowflake(branch, level-1, length*0.5, thick*0.5, hue+70.0);
    }
}

int main(void)
{
    // SetConfigFlags(FLAG_WINDOW_RESIZABLE);
    InitWindow(WIDTH, HEIGHT, "Copo de nieve");
    SetTargetFPS(60);

    while (!WindowShouldClose())
    {
        BeginDrawing();
        ClearBackground(BLACK);
        Vector2 center = {GetScreenWidth()*0.5, GetScreenHeight()*0.5};
        drawSnowflake(center, BRANCH_LEVEL, BRANCH_LEN, BRANCH_THICKNESS, 0.0);
        EndDrawing();
    }
    CloseWindow();

    return 0;
}
