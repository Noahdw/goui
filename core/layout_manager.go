package core

import (
	. "github.com/noahdw/goui/node"
	"github.com/noahdw/goui/node/style"
)

// LayoutManager handles the layout calculation and management
type LayoutManager struct {
	rootNode      Node
	renderContext RenderContext
	windowWidth   float64
	windowHeight  float64
	needsLayout   bool
}

// NewLayoutManager creates a new layout manager
func NewLayoutManager(root Node, context RenderContext, width, height float64) *LayoutManager {
	return &LayoutManager{
		rootNode:      root,
		renderContext: context,
		windowWidth:   width,
		windowHeight:  height,
		needsLayout:   true,
	}
}

// UpdateLayout performs a complete layout pass if needed
func (l *LayoutManager) UpdateLayout() bool {
	if !l.needsLayout {
		return false
	}

	// Pass 1: Resolve styles
	l.rootNode.ResolveStyles(style.NewStyles(make(map[string]interface{})))

	// Pass 2: Measure preferred sizes
	l.rootNode.MeasurePreferred(l.renderContext)

	// Pass 3: Apply constraints and layout
	viewport := Constraints{
		MinWidth:  0,
		MaxWidth:  l.windowWidth,
		MinHeight: 0,
		MaxHeight: l.windowHeight,
	}
	finalSize := l.rootNode.Layout(l.renderContext, viewport)

	// Pass 4: Position elements
	bounds := style.Rect{
		Position: style.Point{X: 0, Y: 0},
		Size:     finalSize,
	}
	l.rootNode.ArrangeChildren(l.renderContext, bounds)

	l.needsLayout = false
	return true
}

// MarkDirty marks the layout as needing recalculation
func (l *LayoutManager) MarkDirty() {
	l.needsLayout = true
}

// UpdateWindowSize updates the window dimensions
func (l *LayoutManager) UpdateWindowSize(width, height float64) {
	if l.windowWidth != width || l.windowHeight != height {
		l.windowWidth = width
		l.windowHeight = height
		l.needsLayout = true
	}
}

// GetWindowSize returns the current window dimensions
func (l *LayoutManager) GetWindowSize() (width, height float64) {
	return l.windowWidth, l.windowHeight
}
