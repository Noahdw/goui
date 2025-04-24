package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/bounds"
	"github.com/noahdw/goui/component"
	. "github.com/noahdw/goui/component"
	. "github.com/noahdw/goui/quadtree"
)

// Render()
// Bounds() rect
// children()
// addChild()
// removeChild()

type Button struct {
	BaseComponent
}

func (b *Button) Render() {
	rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), rl.Red)

}

func NewButton(bounds bounds.Bounds) *Button {
	return &Button{
		BaseComponent: BaseComponent{
			BaseBounds: BaseBounds{
				Bounds: bounds,
			},
		},
	}
}

type Node struct {
	pos rl.Vector2
}

func (n *Node) Draw() {

}

func main() {
	rl.InitWindow(1000, 700, "Vode")
	rl.SetTargetFPS(120)
	defer rl.CloseWindow()

	root := &BaseComponent{}
	button := NewButton(Bounds{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
	})
	button.OnMouseEvent(func() {
		println("Wow it works")
	})

	button2 := NewButton(Bounds{
		X:      400,
		Y:      400,
		Width:  100,
		Height: 200,
	})
	button2.OnMouseEvent(func() {
		println("Wow it works2")
	})

	root.AddChild(button)
	root.AddChild(button2)

	qt := Quadtree{
		Bounds: bounds.Bounds{
			X:      0,
			Y:      0,
			Width:  10000,
			Height: 10000,
		},
		MaxObjects: 4,
		MaxLevels:  8,
		Level:      0,
		Objects:    make([]component.Component, 0),
		Nodes:      make([]Quadtree, 0),
	}

	camera := rl.Camera2D{
		Zoom: 1.0,
	}

	nodes := []component.Component{}
	nodes = append(nodes, root)
	addToQuadtree(root, &qt)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.Beige)
		mouseEventConsumed := false

		mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
		cursor := &component.BaseComponent{
			BaseBounds: BaseBounds{
				Bounds: Bounds{
					X:      float64(mouseWorldPos.X),
					Y:      float64(mouseWorldPos.Y),
					Width:  0,
					Height: 0,
				},
			},
		}

		var foundObj component.Component
		objs := qt.Retrieve(cursor)
		//println(len(objs))
		for _, objUnderCursor := range objs {
			if cursor.Intersects(objUnderCursor.BoundingRect()) {
				foundObj = objUnderCursor
				mouseEventConsumed = true
				println("hmm")
			}
		}

		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) && mouseEventConsumed {
			foundObj.HandleMouse()
		}

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) && !mouseEventConsumed {
			delta := rl.GetMouseDelta()
			delta = rl.Vector2Scale(delta, -1.0/camera.Zoom)
			camera.Target = rl.Vector2Add(camera.Target, delta)

		}

		// Handle mouse wheel
		wheel := rl.GetMouseWheelMove()
		if wheel != 0 {
			// Get the world point that is under the mouse
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
			println(rl.GetMousePosition().X, rl.GetMousePosition().Y)

			// Set the offset to where the mouse is
			camera.Offset = rl.GetMousePosition()
			// Set the target to match, so that the camera maps the world space point
			// under the cursor to the screen space point under the cursor at any zoom
			camera.Target = mouseWorldPos

			// Zoom increment
			// Uses log scaling to provide consistent zoom speed
			scale := 0.2 * wheel
			camera.Zoom = rl.Clamp(float32(math.Exp(math.Log(float64(camera.Zoom))+float64(scale))), 0.125, 64.0)
		}

		// Render the nodes
		for _, node := range nodes {
			node.Render()
		}

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func addToQuadtree(node Component, qt *Quadtree) {
	qt.Insert(node)
	for _, child := range node.Children() {
		addToQuadtree(child, qt)
	}
}
