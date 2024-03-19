package main

import (
	"encoding/json"
	"math"
	"os"

	"github.com/crazy3lf/colorconv"
	ray "github.com/gen2brain/raylib-go/raylib"
)

type Turtle struct {
  X, Y int32
  angle float64
}

func (t *Turtle) branch(length float64, depth int) {
  if depth > state.TREE_DEPTH { return }
  dir := 1.0
  if depth % 2 == 1 { dir = -1.0 }
  t.forward(state.MAIN_LEN*length, depth)
  t.angle += state.BRANCH_ANGLE * dir
  t.forward(state.BRANCH_LEN_FACTOR*state.MAIN_LEN*length, depth)
  t.branch(state.LEN_FACTOR*length, depth+1)
  t.back(state.BRANCH_LEN_FACTOR*state.MAIN_LEN*length)
  t.angle -= 2*state.BRANCH_ANGLE * dir
  t.forward(state.BRANCH_LEN_FACTOR*state.MAIN_LEN*length, depth)
  t.branch(state.LEN_FACTOR*length, depth+1)
  t.back(state.BRANCH_LEN_FACTOR*state.MAIN_LEN*length)
  t.angle += state.BRANCH_ANGLE * dir
  t.back(state.MAIN_LEN*length)
}

func (t *Turtle) back(length float64) {
  t.X -= int32(length*math.Cos(t.angle))
  t.Y -= int32(length*math.Sin(t.angle))
}

func (t *Turtle) forward(length float64, depth int) {
  nextX := t.X + int32(length*math.Cos(t.angle))
  nextY := t.Y + int32(length*math.Sin(t.angle))
  r, g, b, _ := colorconv.HSLToRGB(float64(depth)/float64(state.TREE_DEPTH)*255, 1, .5)
  clr := ray.NewColor(r,g,b,255)
  ray.DrawLine(
    t.X, t.Y,
    nextX, nextY,
    clr,
    // ray.White,
  )
  t.X, t.Y = nextX, nextY
}

func main() {
  ray.SetTraceLogLevel(ray.LogError)
  ray.InitWindow(2560, 1440, "Tree")
  ray.ToggleFullscreen()
  defer ray.CloseWindow()
  ray.SetTargetFPS(120)
  for !ray.WindowShouldClose() {
    if ray.IsKeyPressed(ray.KeyEscape) { break }
    if ray.IsKeyDown(ray.KeyH) { state.BRANCH_ANGLE += .02 }
    if ray.IsKeyDown(ray.KeyL) { state.BRANCH_ANGLE -= .02 }
    if ray.IsKeyDown(ray.KeyK) { state.LEN_FACTOR += .01 }
    if ray.IsKeyDown(ray.KeyJ) { state.LEN_FACTOR -= .01 }
    if ray.IsKeyDown(ray.KeyPeriod) { state.BRANCH_LEN_FACTOR += .01 }
    if ray.IsKeyDown(ray.KeyComma) { state.BRANCH_LEN_FACTOR -= .01 }
    if ray.IsKeyDown(ray.KeyRightBracket) { state.MAIN_LEN += 1 }
    if ray.IsKeyDown(ray.KeyLeftBracket) { state.MAIN_LEN -= 1 }
    if ray.IsKeyDown(ray.KeyZero) { state.XOFFSET += 3 }
    if ray.IsKeyDown(ray.KeyNine) { state.XOFFSET -= 3 }
    if ray.IsKeyDown(ray.KeyTwo) { state.YOFFSET += 3 }
    if ray.IsKeyDown(ray.KeyOne) { state.YOFFSET -= 3 }
    if ray.IsKeyDown(ray.KeyW) { state.LEN_OFFSET += 1 }
    if ray.IsKeyDown(ray.KeyQ) { state.LEN_OFFSET -= 1 }
    if ray.IsKeyPressed(ray.KeyEqual) { state.TREE_DEPTH ++ }
    if ray.IsKeyPressed(ray.KeyMinus) { state.TREE_DEPTH -- }
    if ray.IsKeyPressed(ray.KeyM) { state.COUNT ++ }
    if ray.IsKeyPressed(ray.KeyN) { state.COUNT -- }
    if ray.IsKeyPressed(ray.KeyO) {
      data, err := os.ReadFile("config.json")
      if err != nil { panic(err) }
      err = json.Unmarshal(data, &state)
      if err != nil { panic(err) }
    }
    if ray.IsKeyPressed(ray.KeyP) {
      data, err := json.Marshal(state)
      if err != nil { panic(err) }
      err = os.WriteFile("config.json", data, 0644)
      if err != nil { panic(err) }
    }
    if ray.IsKeyPressed(ray.KeyS) {
      ray.TakeScreenshot("screen.png")
    }
    turtles := []Turtle{}
    for i := range state.COUNT {
      turtle := Turtle{
        X: state.XOFFSET, Y: state.YOFFSET,
        angle: 2*math.Pi/float64(state.COUNT)*float64(i),
      }
      turtle.back(-state.LEN_OFFSET)
      turtles = append(turtles, turtle)
    }
    ray.BeginDrawing()
    ray.ClearBackground(ray.Black)
    for _, t := range turtles {
      t.branch(1, 0)
    }
    ray.EndDrawing()
  }
}

type State struct {
  MAIN_LEN float64
  BRANCH_LEN_FACTOR float64
  BRANCH_ANGLE float64
  TREE_DEPTH int
  LEN_FACTOR float64
  XOFFSET int32
  YOFFSET int32
  COUNT int
  LEN_OFFSET float64
}

var state = State{}
