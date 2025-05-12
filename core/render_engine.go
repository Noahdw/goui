package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	. "github.com/noahdw/goui/node"
	"github.com/noahdw/goui/node/style"
)

// RenderEngine handles the rendering process
type RenderEngine struct {
	rootNode      Node
	renderContext RenderContext
	layoutManager *LayoutManager
	eventManager  *EventManager
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
	lastMouseX    float64 // Track last mouse position
	lastMouseY    float64
}

// NewRenderEngine creates a new render engine
func NewRenderEngine(root Node, context RenderContext, width, height float64) *RenderEngine {
	engine := &RenderEngine{
		rootNode:      root,
		renderContext: context,
		windowWidth:   width,
		windowHeight:  height,
		needsLayout:   true,
		camera: rl.Camera2D{
			Zoom: 1,
		},
	}

	engine.layoutManager = NewLayoutManager(root, context, width, height)
	engine.eventManager = NewEventManager(engine)

	return engine
}

// RenderFrame handles a single frame of rendering
func (r *RenderEngine) RenderFrame() {
	// Update window size if needed
	if rl.IsWindowResized() {
		r.layoutManager.UpdateWindowSize(
			float64(rl.GetScreenWidth()),
			float64(rl.GetScreenHeight()),
		)
	}

	// Update layout if needed
	r.layoutManager.UpdateLayout()

	// Get mouse position in world coordinates
	mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), r.camera)
	mouseX := float64(mouseWorldPos.X)
	mouseY := float64(mouseWorldPos.Y)

	// Find object under cursor
	cursor := style.Rect{
		Position: style.Point{X: mouseX, Y: mouseY},
		Size:     style.Size{Width: 1, Height: 1},
	}
	foundObj := r.getObjUnderCursor(r.rootNode, cursor)

	// Handle all events
	r.eventManager.HandleMouseEvents(mouseX, mouseY, foundObj)
	r.eventManager.HandleKeyboardEvents()

	// Clear the screen and render
	r.renderContext.Clear()
	r.rootNode.Paint(r.renderContext)
}

// MarkLayoutDirty marks the layout as needing recalculation
func (r *RenderEngine) MarkLayoutDirty() {
	r.layoutManager.MarkDirty()
}

// GetCamera returns the current camera
func (r *RenderEngine) GetCamera() rl.Camera2D {
	return r.camera
}

// SetFocus sets the currently focused node
func (r *RenderEngine) SetFocus(node Node) {
	r.eventManager.SetFocus(node)
}

func (r *RenderEngine) getObjUnderCursor(node Node, cursor style.Rect) Node {
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
