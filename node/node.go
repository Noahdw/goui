package node

import (
	"fmt"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node interface {
	// Style methods
	Width(value interface{}) Node
	Height(value interface{}) Node
	Margin(value interface{}) Node
	Padding(value interface{}) Node
	FontSize(value interface{}) Node
	Color(value string) Node
	Opacity(value float64) Node
	Background(value string) Node
	Border(value string) Node
	Flex(value string) Node
	BorderWidth(value interface{}) Node
	BorderRadius(value interface{}) Node

	// Structure methods
	AddChildren(children ...Node)
	Children() []Node
	GetType() string
	GetStyles() Styles
	Intersects(node Node) bool
	Parent() Node
	SetParent(parent Node)
	DispatchEvent(event UIEvent)

	// Layout methods
	ResolveStyles(parentStyles Styles) Styles
	MeasurePreferred(ctx RenderContext) Size
	Layout(ctx RenderContext, constraints Constraints) Size
	ArrangeChildren(ctx RenderContext, bounds Rect)
	Paint(ctx RenderContext)
	GetFinalSize() Size
	GetFinalBounds() Rect
}

type BaseNode struct {
	nodeType       string
	parent         Node
	children       []Node
	styles         Styles
	finalSize      Size
	finalBounds    Rect
	preferredSize  Size
	finalOpacity   float64
	eventCallbacks map[UIEventType]func(UIEvent)
}

type Event struct {
}

func NewBaseNode(nodeType string, styles Styles) BaseNode {
	return BaseNode{
		nodeType:       nodeType,
		styles:         styles,
		eventCallbacks: make(map[UIEventType]func(UIEvent)),
	}
}

// Size represents width and height
type Size struct {
	Width, Height float64
}

// Point represents x and y coordinates
type Point struct {
	X, Y float64
}

// Rect represents a rectangle with position and size
type Rect struct {
	Position Point
	Size     Size
}

// Constraints represents size constraints
type Constraints struct {
	MinWidth, MaxWidth, MinHeight, MaxHeight float64
}

// RenderContext provides context for rendering
type RenderContext interface {
	// Drawing methods would go here
	Clear()
	DrawBackground(bounds Rect, styles Styles, opacity float64)
	DrawBorders(bounds Rect, styles Styles, opacity float64)
	DrawText(text string, bounds Rect, styles Styles, opacity float64)
	DrawTexture(sourceURL string, bounds Rect, styles Styles, opacity float64)
	ClipRect() Rect
	LoadTexture(sourceURL string) rl.Texture2D
	Present()
}

// Style method implementations for BaseNode
func (n *BaseNode) Width(value interface{}) Node {
	n.styles.Width = parseStyleValue(value)
	n.styles.setProperties[string(WidthProp)] = Explicit
	return n
}

func (n *BaseNode) Height(value interface{}) Node {
	n.styles.Height = parseStyleValue(value)
	n.styles.setProperties[string(HeightProp)] = Explicit
	return n
}

func (n *BaseNode) Margin(value interface{}) Node {
	m := parseMarginPadding(value)
	n.styles.Margin = m
	n.styles.setProperties[string(MarginProp)] = Explicit
	return n
}

func (n *BaseNode) Padding(value interface{}) Node {
	p := parseMarginPadding(value)
	n.styles.Padding = p
	n.styles.setProperties[string(PaddingProp)] = Explicit
	return n
}

func (n *BaseNode) Border(value string) Node {
	n.styles.Border.Style = value
	n.styles.setProperties[string(BorderProp)] = Explicit
	return n
}

func (n *BaseNode) BorderWidth(value interface{}) Node {
	n.styles.Border.Width = parseMarginPadding(value)
	n.styles.setProperties[string(BorderProp)] = Explicit
	return n
}

func (n *BaseNode) BorderRadius(value interface{}) Node {
	m := parseMarginPadding(value)
	n.styles.BorderRadius = m
	n.styles.setProperties[string(BorderRadiusProp)] = Explicit
	return n
}

func (n *BaseNode) Flex(value string) Node {
	n.styles.FlexDirection = value
	n.styles.setProperties[string(FlexDirectionProp)] = Explicit
	return n
}

func (n *BaseNode) FontSize(value interface{}) Node {
	n.styles.FontSize = parseStyleValue(value)
	n.styles.setProperties[string(FontSizeProp)] = Explicit
	return n
}

func (n *BaseNode) Color(value string) Node {
	n.styles.Color = parseColor(value)
	n.styles.setProperties[string(ColorProp)] = Explicit
	return n
}

func (n *BaseNode) Background(value string) Node {
	n.styles.Background = parseColor(value)
	n.styles.setProperties[string(BackgroundProp)] = Explicit
	return n
}

func (n *BaseNode) Opacity(value float64) Node {
	n.styles.Opacity = value
	n.styles.setProperties[string(OpacityProp)] = Explicit
	return n
}

// Helper to parse margin/padding values
func parseMarginPadding(value interface{}) EdgeInsets {
	switch v := value.(type) {
	case int:
		return EdgeInsets{float64(v), float64(v), float64(v), float64(v)}
	case float64:
		return EdgeInsets{v, v, v, v}
	case [4]int:
		return EdgeInsets{float64(v[0]), float64(v[1]), float64(v[2]), float64(v[3])}
	case [4]float64:
		return EdgeInsets{v[0], v[1], v[2], v[3]}
	case [2]int:
		return EdgeInsets{float64(v[0]), float64(v[1]), float64(v[0]), float64(v[1])}
	case [2]float64:
		return EdgeInsets{v[0], v[1], v[0], v[1]}
	}
	// Default
	return EdgeInsets{0, 0, 0, 0}
}

// Helper to parse color values
func parseColor(value string) Color {
	// Handle named colors
	switch value {
	case "black":
		return Black
	case "white":
		return White
	case "red":
		return Red
	case "green":
		return Green
	case "blue":
		return Blue
	case "yellow":
		return Yellow
	case "cyan":
		return Cyan
	case "magenta":
		return Magenta
	case "gray":
		return Gray
	case "transparent":
		return Transparent
	}

	// Handle hex colors
	if strings.HasPrefix(value, "#") {
		hex := value[1:]
		if len(hex) == 3 {
			// #RGB format
			r, _ := strconv.ParseUint(string(hex[0])+string(hex[0]), 16, 8)
			g, _ := strconv.ParseUint(string(hex[1])+string(hex[1]), 16, 8)
			b, _ := strconv.ParseUint(string(hex[2])+string(hex[2]), 16, 8)
			return Color{uint8(r), uint8(g), uint8(b), 255}
		} else if len(hex) == 6 {
			// #RRGGBB format
			r, _ := strconv.ParseUint(hex[0:2], 16, 8)
			g, _ := strconv.ParseUint(hex[2:4], 16, 8)
			b, _ := strconv.ParseUint(hex[4:6], 16, 8)
			return Color{uint8(r), uint8(g), uint8(b), 255}
		}
	}

	// Default to black for invalid colors
	return Black
}

func (n *BaseNode) Parent() Node {
	return n.parent
}

func (n *BaseNode) SetParent(parent Node) {
	n.parent = parent
}

// Structure methods
func (n *BaseNode) AddChildren(children ...Node) {
	// We catch the event nodes being added and rather than add them as children to this node,
	// we filter them out and incorporate their callbacks into this node.
	// This allows for a more declarative UI construction where events are treated as first-class nodes.
	var finalChildren []Node
	for _, child := range children {
		if eventNode, ok := child.(*EventNode); ok {
			// Register the event callback with this node
			n.eventCallbacks[eventNode.eventType] = eventNode.callback
		} else {
			child.SetParent(n)
			finalChildren = append(finalChildren, child)
		}
	}
	n.children = append(n.children, finalChildren...)
}

func (n *BaseNode) Children() []Node {
	return n.children
}

func (n *BaseNode) GetType() string {
	return n.nodeType
}

func (n *BaseNode) GetStyles() Styles {
	return n.styles
}

// Layout methods
func (n *BaseNode) ResolveStyles(parentStyles Styles) Styles {
	// Start with this node's styles
	resolvedStyles := n.styles

	// For inheritable properties, check if they're set in this node
	// If not, inherit from parent
	inheritableProps := []StyleProperty{
		FontFamilyProp, FontSizeProp, ColorProp, LineHeightProp, BackgroundProp, OpacityProp,
	}

	for _, prop := range inheritableProps {
		// Only inherit if:
		// 1. Parent has the property set (explicitly or inherited)
		// 2. This node doesn't have it explicitly set
		parentSource, parentHasIt := parentStyles.setProperties[string(prop)]
		parentHasIt = (parentSource == Explicit || parentSource == Inherited)

		selfSource, selfHasIt := resolvedStyles.setProperties[string(prop)]
		selfHasIt = (selfSource == Explicit)

		if parentHasIt && !selfHasIt {

			switch prop {
			case FontFamilyProp:
				resolvedStyles.FontFamily = parentStyles.FontFamily
				resolvedStyles.setProperties[string(prop)] = Inherited
			case FontSizeProp:
				resolvedStyles.FontSize = parentStyles.FontSize
				resolvedStyles.FontSize.Source = Inherited
				resolvedStyles.setProperties[string(prop)] = Inherited
			case ColorProp:
				resolvedStyles.Color = parentStyles.Color
				resolvedStyles.setProperties[string(prop)] = Inherited
			case OpacityProp:
				resolvedStyles.Opacity = parentStyles.Opacity
				resolvedStyles.setProperties[string(prop)] = Inherited
			case BackgroundProp:
				resolvedStyles.Background = parentStyles.Background
				resolvedStyles.setProperties[string(prop)] = Inherited
			case LineHeightProp:
				resolvedStyles.LineHeight = parentStyles.LineHeight
				resolvedStyles.LineHeight.Source = Inherited
				resolvedStyles.setProperties[string(prop)] = Inherited
			}

		}
	}

	resolvedStyles.finalOpacity = resolvedStyles.Opacity * parentStyles.finalOpacity
	n.finalOpacity = resolvedStyles.finalOpacity

	// Apply the same process to all children
	for _, child := range n.children {
		child.ResolveStyles(resolvedStyles)
	}
	n.styles = resolvedStyles
	return resolvedStyles
}

func (n *BaseNode) MeasurePreferred(ctx RenderContext) Size {
	// For leaf nodes, calculate intrinsic size
	if len(n.children) == 0 {
		return Size{0, 0} // Override in specific node types
	}

	if n.styles.Opacity == 0 {
		return Size{0, 0} // Dont render hidden objects
	}

	// For container nodes, measure all children first
	childSizes := make([]Size, len(n.children))
	for i, child := range n.children {
		childSizes[i] = child.MeasurePreferred(ctx)
	}

	// Based on layout type, calculate how big this node needs to be
	var totalSize Size
	switch n.styles.FlexDirection {
	case "row":
		// Sum width, max height
		for i, size := range childSizes {
			totalSize.Width += size.Width
			if size.Height > totalSize.Height {
				totalSize.Height = size.Height
			}
			// Add margin between items
			if i < len(childSizes)-1 {
				child1 := n.children[i]
				child2 := n.children[i+1]
				totalSize.Width += child1.GetStyles().Margin.Right + child2.GetStyles().Margin.Left
			}
		}
	case "column":
		// Sum height, max width
		for i, size := range childSizes {
			totalSize.Height += size.Height
			if size.Width > totalSize.Width {
				totalSize.Width = size.Width
			}
			// Add margin between items
			if i < len(childSizes)-1 {
				child1 := n.children[i]
				child2 := n.children[i+1]
				totalSize.Height += child1.GetStyles().Margin.Bottom + child2.GetStyles().Margin.Top
			}
		}
	}

	// Add padding
	totalSize.Width += n.styles.Padding.Left + n.styles.Padding.Right
	totalSize.Height += n.styles.Padding.Top + n.styles.Padding.Bottom
	n.preferredSize = totalSize
	return totalSize
}

func (n *BaseNode) Layout(ctx RenderContext, constraints Constraints) Size {
	// Start with preferred size
	preferredSize := n.preferredSize
	fmt.Println(preferredSize.Width)
	// Apply constraints to determine final size
	finalSize := Size{
		Width:  clamp(preferredSize.Width, constraints.MinWidth, constraints.MaxWidth),
		Height: clamp(preferredSize.Height, constraints.MinHeight, constraints.MaxHeight),
	}

	// Handle special cases like Fill Parent
	if n.styles.Width.Type == PERCENTAGE {
		finalSize.Width = constraints.MaxWidth * n.styles.Width.Value / 100
	}

	if n.styles.Height.Type == PERCENTAGE {
		finalSize.Height = constraints.MaxHeight * n.styles.Height.Value / 100
	}

	// If we have children, layout them now
	if len(n.children) > 0 {
		// Calculate available space for children (minus padding)
		availableWidth := finalSize.Width - n.styles.Padding.Left - n.styles.Padding.Right
		availableHeight := finalSize.Height - n.styles.Padding.Top - n.styles.Padding.Bottom

		// Create child constraints based on layout direction
		childConstraints := Constraints{
			MinWidth:  0,
			MaxWidth:  availableWidth,
			MinHeight: 0,
			MaxHeight: availableHeight,
		}

		// Layout each child
		for _, child := range n.children {
			_ = child.Layout(ctx, childConstraints)
		}
	}

	// Store final size for use in arrangement pass
	n.finalSize = finalSize
	return finalSize
}

func (n *BaseNode) ArrangeChildren(ctx RenderContext, bounds Rect) {
	// Set this node's final bounds
	n.finalBounds = bounds

	// For leaf nodes, we're done
	if len(n.children) == 0 {
		return
	}

	// Calculate content area (bounds minus padding)
	contentArea := Rect{
		Position: Point{
			X: bounds.Position.X + n.styles.Padding.Left,
			Y: bounds.Position.Y + n.styles.Padding.Top,
		},
		Size: Size{
			Width:  bounds.Size.Width - n.styles.Padding.Left - n.styles.Padding.Right,
			Height: bounds.Size.Height - n.styles.Padding.Top - n.styles.Padding.Bottom,
		},
	}

	// Position each child based on layout direction
	var currentX = contentArea.Position.X
	var currentY = contentArea.Position.Y

	for i, child := range n.children {
		childSize := child.GetFinalSize()
		childStyles := child.GetStyles()

		if n.styles.FlexDirection == "row" {
			childBounds := Rect{
				Position: Point{X: currentX, Y: currentY},
				Size:     childSize,
			}
			child.ArrangeChildren(ctx, childBounds)
			// Add child's right margin and next child's left margin
			if i < len(n.children)-1 {
				nextChild := n.children[i+1]
				currentX += childSize.Width + childStyles.Margin.Right + nextChild.GetStyles().Margin.Left
			}
		} else { // column
			childBounds := Rect{
				Position: Point{X: currentX, Y: currentY},
				Size:     childSize,
			}
			child.ArrangeChildren(ctx, childBounds)
			// Add child's bottom margin and next child's top margin
			if i < len(n.children)-1 {
				nextChild := n.children[i+1]
				currentY += childSize.Height + childStyles.Margin.Bottom + nextChild.GetStyles().Margin.Top
			}
		}
	}
}

func (n *BaseNode) GetFinalSize() Size {
	return n.finalSize
}

func (n *BaseNode) GetFinalBounds() Rect {
	return n.finalBounds
}

func (n *BaseNode) Paint(ctx RenderContext) {
	// Draw this node's background
	ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

	// Draw borders if needed
	if n.styles.Border.CanDisplay() {
		ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
	}

	// Draw all children
	for _, child := range n.children {
		child.Paint(ctx)
	}
}

// TextNode is a specialized node for text content
type TextNode struct {
	BaseNode
	text string
}

func NewTextNode(baseNode BaseNode, text string) Node {
	return &TextNode{
		BaseNode: baseNode,
		text:     text,
	}
}

// Specialized implementation for TextNode
func (n *TextNode) MeasurePreferred(ctx RenderContext) Size {
	fontSize := int32(n.styles.FontSize.Value)
	textWidth := float64(rl.MeasureText(n.text, fontSize))
	textHeight := float64(fontSize) * 1.2 // Add some line height

	// Add padding to the text size
	return Size{
		Width:  textWidth + n.styles.Padding.Left + n.styles.Padding.Right,
		Height: textHeight + n.styles.Padding.Top + n.styles.Padding.Bottom,
	}
}

func (n *TextNode) Paint(ctx RenderContext) {
	// Draw background and border first
	ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

	// Draw borders if needed
	if n.styles.Border.CanDisplay() {
		ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
	}

	// Then draw the text
	ctx.DrawText(n.text, n.finalBounds, n.styles, n.finalOpacity)
}

type ImageNode struct {
	BaseNode
	sourceURL string
}

func NewImageNode(baseNode BaseNode, sourceURL string) Node {
	imageNode := ImageNode{
		BaseNode:  baseNode,
		sourceURL: sourceURL,
	}
	return &imageNode
}

type EventNode struct {
	BaseNode
	callback  func(UIEvent)
	eventType UIEventType
}

func NewEventNode(baseNode BaseNode, eventType UIEventType, callback func(UIEvent)) Node {
	eventNode := EventNode{
		BaseNode:  baseNode,
		callback:  callback,
		eventType: eventType,
	}
	return &eventNode
}

func (n *ImageNode) Paint(ctx RenderContext) {
	// Draw background and border first
	ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

	// Draw borders if needed
	if n.styles.Border.CanDisplay() {
		ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
	}

	// Then draw the text
	ctx.DrawTexture(n.sourceURL, n.finalBounds, n.styles, n.finalOpacity)
}

func (n *ImageNode) MeasurePreferred(ctx RenderContext) Size {
	texture := ctx.LoadTexture(n.sourceURL)
	n.preferredSize = Size{
		Width:  float64(texture.Width),
		Height: float64(texture.Height),
	}
	return n.preferredSize
}

// Helper function to clamp a value between min and max
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func (b *BaseNode) Intersects(node Node) bool {
	return b.finalBounds.Intersects(node.GetFinalBounds())
}

// Intersects - Checks if a Bounds object intersects with another Bounds
func (r *Rect) Intersects(node Rect) bool {

	aMaxX := r.Position.X + r.Size.Width
	aMaxY := r.Position.Y + r.Size.Height
	bMaxX := node.Position.X + node.Size.Width
	bMaxY := node.Position.Y + node.Size.Height

	// a is left of b
	if aMaxX < node.Position.X {
		return false
	}

	// a is right of b
	if r.Position.X > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < node.Position.Y {
		return false
	}

	// a is below b
	if r.Position.Y > bMaxY {
		return false
	}

	// The two overlap
	return true
}

// RemoveEventHandler removes an event handler for the given event type
func (n *BaseNode) RemoveEventHandler(eventType UIEventType) {
	delete(n.eventCallbacks, eventType)
}

// DispatchEvent dispatches an event to this node and its parent chain
func (n *BaseNode) DispatchEvent(event UIEvent) {
	// Set the target if not already set
	if event.Target == nil {
		event.Target = n
	}

	// Call the handler if one exists
	if handler, has := n.eventCallbacks[event.Type]; has {
		handler(event)
		if event.StopPropagation {
			return
		}
	}

	// Propagate to parent
	if parent := n.parent; parent != nil {
		parent.DispatchEvent(event)
	}
}
