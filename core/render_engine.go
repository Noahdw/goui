package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	. "github.com/noahdw/goui/node"
)

// RenderEngine handles the rendering process
type RenderEngine struct {
	rootNode      Node
	renderContext RenderContext
	windowWidth   float64
	windowHeight  float64
	needsLayout   bool
	camera        rl.Camera2D
	hasAnimations bool
	lastFoundObj  Node
	focusedNode   Node // Currently focused node for keyboard events
	pressedObj    Node // Node that was pressed, for click detection
	pressX        float64
	pressY        float64
}

// NewRenderEngine creates a new render engine
func NewRenderEngine(root Node, context RenderContext, width, height float64) *RenderEngine {
	return &RenderEngine{
		rootNode:      root,
		renderContext: context,
		windowWidth:   width,
		windowHeight:  height,
		needsLayout:   true,
		camera: rl.Camera2D{
			Zoom: 1,
		},
	}
}

// SetFocus sets the currently focused node
func (r *RenderEngine) SetFocus(node Node) {
	// Notify old focused node that it lost focus
	if r.focusedNode != nil {
		event := NewUIEvent(UIBlur, r.focusedNode)
		r.focusedNode.DispatchEvent(event)
	}

	// Set new focus
	r.focusedNode = node

	// Notify new focused node
	if node != nil {
		event := NewUIEvent(UIFocus, node)
		node.DispatchEvent(event)
	}
}

func (r *RenderEngine) RenderFrame() {
	mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), r.camera)
	cursor := Rect{
		Position: Point{
			X: float64(mouseWorldPos.X),
			Y: float64(mouseWorldPos.Y),
		}}
	var foundObj = r.getObjUnderCursor(r.rootNode, cursor)
	if foundObj != nil {
		if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			// Store the pressed object and position for click detection
			r.pressedObj = foundObj
			r.pressX = float64(mouseWorldPos.X)
			r.pressY = float64(mouseWorldPos.Y)

			// Dispatch press event
			event := NewUIMouseEvent(UIPress, foundObj, r.pressX, r.pressY)
			foundObj.DispatchEvent(event)

			// Set focus on press
			r.SetFocus(foundObj)
		}

		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			event := NewUIMouseEvent(UIMove, foundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
			foundObj.DispatchEvent(event)
		}

		if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			// Dispatch release event
			event := NewUIMouseEvent(UIRelease, foundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
			foundObj.DispatchEvent(event)

			// If we released on the same object we pressed, it's a click
			if r.pressedObj == foundObj {
				// Calculate if the mouse moved too far during press (drag threshold)
				dx := float64(mouseWorldPos.X) - r.pressX
				dy := float64(mouseWorldPos.Y) - r.pressY
				distance := dx*dx + dy*dy

				// If the mouse didn't move too far, it's a click
				if distance < 100 { // 10 pixels threshold
					clickEvent := NewUIMouseEvent(UIClick, foundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
					foundObj.DispatchEvent(clickEvent)
				}
			}

			// Clear pressed state
			r.pressedObj = nil
		}

		// Last frame we found a different object, so we know that the mouse both
		// entered the found object and exited the previous frames object.
		if foundObj != r.lastFoundObj {
			event := NewUIMouseEvent(UIEnter, foundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
			foundObj.DispatchEvent(event)
			if r.lastFoundObj != nil {
				event := NewUIMouseEvent(UILeave, r.lastFoundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
				r.lastFoundObj.DispatchEvent(event)
			}
			r.lastFoundObj = foundObj
		}
	} else if r.lastFoundObj != nil {
		event := NewUIMouseEvent(UILeave, r.lastFoundObj, float64(mouseWorldPos.X), float64(mouseWorldPos.Y))
		r.lastFoundObj.DispatchEvent(event)
		r.lastFoundObj = nil
	}

	// Handle keyboard events - only send to focused node
	key := rl.GetKeyPressed()
	if key != 0 && r.focusedNode != nil {
		event := NewUIKeyboardEvent(UIKeyPress, r.focusedNode, int(key), rune(key))
		r.focusedNode.DispatchEvent(event)
	}

	// Handle Tab key for focus navigation
	if rl.IsKeyPressed(rl.KeyTab) && r.focusedNode != nil {
		// TODO: Implement focus navigation logic
		// This would involve finding the next focusable node in the tree
		// For now, we'll just clear focus
		r.SetFocus(nil)
	}

	// Check if layout needs recalculation
	if r.needsRender() {
		// Pass 1: Resolve styles
		r.rootNode.ResolveStyles(NewStyles(StyleProps{}))

		// Pass 2: Measure preferred sizes
		r.rootNode.MeasurePreferred(r.renderContext)

		// Pass 3: Apply constraints and layout
		viewport := Constraints{
			MinWidth:  0,
			MaxWidth:  r.windowWidth,
			MinHeight: 0,
			MaxHeight: r.windowHeight,
		}
		finalSize := r.rootNode.Layout(r.renderContext, viewport)

		// Pass 4: Position elements
		bounds := Rect{
			Position: Point{X: 0, Y: 0},
			Size:     finalSize,
		}
		r.rootNode.ArrangeChildren(r.renderContext, bounds)

		r.needsLayout = false
	}

	// Clear the screen
	r.renderContext.Clear()

	// Pass 5: Actual rendering
	r.rootNode.Paint(r.renderContext)
}

// MarkLayoutDirty marks the layout as needing recalculation
func (r *RenderEngine) MarkLayoutDirty() {
	r.needsLayout = true
}

func (r *RenderEngine) GetCamera() rl.Camera2D {
	return r.camera
}

func (r *RenderEngine) needsRender() bool {
	if rl.IsWindowResized() {
		r.windowHeight = float64(rl.GetScreenHeight())
		r.windowWidth = float64(rl.GetScreenWidth())
		return true
	}
	return r.needsLayout
}

func (r *RenderEngine) getObjUnderCursor(node Node, cursor Rect) Node {
	nodebr := node.GetFinalBounds()
	if !nodebr.Intersects(cursor) {
		return nil
	}
	foundObj := node
	for _, child := range node.Children() {
		fo := r.getObjUnderCursor(child, cursor)
		if fo != nil {
			foundObj = fo
		}
	}
	return foundObj
}
