package core

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"

	. "github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/node"
	. "github.com/noahdw/goui/quadtree"
)

type RenderThread struct {
	camera     rl.Camera2D
	root       Node
	boundsTree Quadtree
}

func NewRenderThread(root Node) RenderThread {
	return RenderThread{
		camera: rl.Camera2D{
			Zoom: 1,
		},
		root: root,
		boundsTree: Quadtree{
			Bounds: Bounds{
				X:      0,
				Y:      0,
				Width:  10000,
				Height: 10000,
			},
			MaxObjects: 4,
			MaxLevels:  8,
			Level:      0,
			Objects:    make([]Node, 0),
			Nodes:      make([]Quadtree, 0),
		},
	}
}

// Begins the render loop, blocking the caller until exit
func (r *RenderThread) StartRenderLoop() {
	// The scene has been set up, all child/parent relations made, so we can set up the quad tree
	// corresponding to the scene
	r.addToBoundstree(r.root)

	// Tracks the object which received the last render cycle's mouse focus event.
	// Useful for determing mouse enter / exit events
	var lastFoundObj Node

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(r.camera)
		rl.ClearBackground(rl.Beige)

		// At the start of every cycle, nodes with dirtied positions must be updated.
		// They get dirtied and updated at the same time since the alternate approach would
		// require in the moment recalcualations which ends up wasteful.
		r.updateNodesWithDirtyPositions(r.root)

		// An object representation of the mouse cursor
		cursor := &BaseNode{}
		cursor.SetSize(0, 0)
		mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), r.camera)
		cursor.SetGlobalPosition(float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
		// To determine what's underneath the cursor, we treat the cursor as a point (Size 0) and see
		// if any bounds of the nodes in the quadtree intersect it.
		var foundObj Node
		objs := r.boundsTree.Retrieve(cursor)
		for _, objUnderCursor := range objs {
			if cursor.Intersects(objUnderCursor) {
				foundObj = objUnderCursor
			}
		}
		// Handle event propogation for mouse
		// Events are sent to the object found under the cursor
		// and are either handled or propogated up the chain allowing
		// for parent nodes to make the same decision.
		if foundObj != nil {
			if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				event := NewMouseEvent(Pressed)
				r.bubbleMouseEvent(event, foundObj)
			}
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				event := NewMouseEvent(Down)
				r.bubbleMouseEvent(event, foundObj)
			}
			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				event := NewMouseEvent(Released)
				r.bubbleMouseEvent(event, foundObj)
			}
			// Last frame we found a different object, so we know that the mouse both
			// entered the found  object and exited the previous frames object.
			if foundObj != lastFoundObj {
				event := NewMouseEvent(Entered)
				r.bubbleMouseEvent(event, foundObj)
				if lastFoundObj != nil {
					event := NewMouseEvent(Exited)
					r.bubbleMouseEvent(event, lastFoundObj)
				}
				lastFoundObj = foundObj
			}
		} else if lastFoundObj != nil {
			event := NewMouseEvent(Exited)
			r.bubbleMouseEvent(event, lastFoundObj)
			lastFoundObj = nil
		}

		// Handle mouse wheel (this does not belong here)
		wheel := rl.GetMouseWheelMove()
		if wheel != 0 {
			// Get the world point that is under the mouse
			mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), r.camera)
			println(rl.GetMousePosition().X, rl.GetMousePosition().Y)

			// Set the offset to where the mouse is
			r.camera.Offset = rl.GetMousePosition()
			// Set the target to match, so that the camera maps the world space point
			// under the cursor to the screen space point under the cursor at any zoom
			r.camera.Target = mouseWorldPos

			// Zoom increment
			// Uses log scaling to provide consistent zoom speed
			scale := 0.2 * wheel
			r.camera.Zoom = rl.Clamp(float32(math.Exp(math.Log(float64(r.camera.Zoom))+float64(scale))), 0.125, 64.0)
		}

		// Render leads to a recursive descent down the tree starting at the root.
		// Parents are rendered before their children
		r.root.Render()

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

func (r *RenderThread) GetCamera() *rl.Camera2D {
	return &r.camera
}

func (r *RenderThread) addToBoundstree(node Node) {
	r.boundsTree.Insert(node)
	for _, child := range node.Children() {
		r.addToBoundstree(child)
	}
}

func (r *RenderThread) bubbleMouseEvent(event MouseEvent, node Node) {
	handleState := node.HandleMouse(event)
	if handleState == Propogate {
		if parent := node.Parent(); parent != nil {
			r.bubbleMouseEvent(event, parent)
		}
	}
}

func (r *RenderThread) updateNodesWithDirtyPositions(node Node) {
	if node.CheckAndClearDirtyPosition() {
		// Position is dirty, update in quadtree
		r.boundsTree.Remove(node.ID())
		r.boundsTree.Insert(node)
	}

	// Check children recursively
	for _, child := range node.Children() {
		r.updateNodesWithDirtyPositions(child)
	}
}
