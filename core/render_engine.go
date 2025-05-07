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
		if foundObj != r.lastFoundObj {
			event := NewMouseEvent(Entered)
			r.bubbleMouseEvent(event, foundObj)
			if r.lastFoundObj != nil {
				event := NewMouseEvent(Exited)
				r.bubbleMouseEvent(event, r.lastFoundObj)
			}
			r.lastFoundObj = foundObj
		}
	} else if r.lastFoundObj != nil {
		event := NewMouseEvent(Exited)
		r.bubbleMouseEvent(event, r.lastFoundObj)
		r.lastFoundObj = nil
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

func (r *RenderEngine) bubbleMouseEvent(event MouseEvent, node Node) {
	handleState := node.HandleMouse(event)
	if handleState == Propogate {
		if parent := node.Parent(); parent != nil {
			r.bubbleMouseEvent(event, parent)
		}
	}
}
