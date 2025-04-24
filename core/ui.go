package core

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"

	. "github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/component"
	. "github.com/noahdw/goui/quadtree"
)

type UI struct {
	camera   rl.Camera2D
	root     Component
	quadtree Quadtree
}

func NewUI(root Component) UI {
	ui := UI{
		camera: rl.Camera2D{
			Zoom: 1,
		},
		root: root,
		quadtree: Quadtree{
			Bounds: Bounds{
				X:      0,
				Y:      0,
				Width:  10000,
				Height: 10000,
			},
			MaxObjects: 4,
			MaxLevels:  8,
			Level:      0,
			Objects:    make([]Component, 0),
			Nodes:      make([]Quadtree, 0),
		},
	}
	return ui
}

func (u *UI) GetCamera() *rl.Camera2D {
	return &u.camera
}

func (u *UI) RenderLoop() {
	u.addToQuadtree(u.root)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(u.camera)
		rl.ClearBackground(rl.Beige)
		//mouseEventConsumed := false

		mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), u.camera)
		cursor := &BaseComponent{
			BaseBounds: BaseBounds{
				Bounds: Bounds{
					X:      float64(mouseWorldPos.X),
					Y:      float64(mouseWorldPos.Y),
					Width:  0,
					Height: 0,
				},
			},
		}

		var foundObj Component
		objs := u.quadtree.Retrieve(cursor)
		for _, objUnderCursor := range objs {
			if cursor.Intersects(objUnderCursor.BoundingRect()) {
				foundObj = objUnderCursor
				//mouseEventConsumed = true
			}
		}
		if foundObj != nil {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				event := NewMouseEvent(Pressed)
				foundObj.HandleMouse(event)
			}
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				event := NewMouseEvent(Down)
				foundObj.HandleMouse(event)
			}
			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				event := NewMouseEvent(Released)
				foundObj.HandleMouse(event)
			}
		}

		// Handle mouse wheel
		wheel := rl.GetMouseWheelMove()
		if wheel != 0 {
			// Get the world point that is under the mouse
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), u.camera)
			println(rl.GetMousePosition().X, rl.GetMousePosition().Y)

			// Set the offset to where the mouse is
			u.camera.Offset = rl.GetMousePosition()
			// Set the target to match, so that the camera maps the world space point
			// under the cursor to the screen space point under the cursor at any zoom
			u.camera.Target = mouseWorldPos

			// Zoom increment
			// Uses log scaling to provide consistent zoom speed
			scale := 0.2 * wheel
			u.camera.Zoom = rl.Clamp(float32(math.Exp(math.Log(float64(u.camera.Zoom))+float64(scale))), 0.125, 64.0)
		}

		// Render the nodes

		u.root.Render()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func (u *UI) addToQuadtree(node Component) {
	u.quadtree.Insert(node)
	for _, child := range node.Children() {
		u.addToQuadtree(child)
	}
}
