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

func (engine *RenderEngine) RenderFrame() {
	// Check if layout needs recalculation
	if engine.needsLayout {
		// Pass 1: Resolve styles
		defaultStyles := NewStyles(StyleProps{})
		engine.rootNode.ResolveStyles(defaultStyles)

		// Pass 2: Measure preferred sizes
		engine.rootNode.MeasurePreferred(engine.renderContext)

		// Pass 3: Apply constraints and layout
		viewport := Constraints{
			MinWidth:  0,
			MaxWidth:  engine.windowWidth,
			MinHeight: 0,
			MaxHeight: engine.windowHeight,
		}
		finalSize := engine.rootNode.Layout(engine.renderContext, viewport)

		// Pass 4: Position elements
		bounds := Rect{
			Position: Point{X: 0, Y: 0},
			Size:     finalSize,
		}
		engine.rootNode.ArrangeChildren(engine.renderContext, bounds)

		engine.needsLayout = false
	}

	// Clear the screen
	engine.renderContext.Clear()

	// Pass 5: Actual rendering
	engine.rootNode.Paint(engine.renderContext)
}

// MarkLayoutDirty marks the layout as needing recalculation
func (engine *RenderEngine) MarkLayoutDirty() {
	engine.needsLayout = true
}

func (r *RenderEngine) GetCamera() rl.Camera2D {
	return r.camera
}
