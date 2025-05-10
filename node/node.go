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
	AlignItems(value string) Node
	FontWeight(value interface{}) Node

	// Structure methods
	AddChildren(children ...Node)
	Children() []Node
	GetType() string
	GetStyles() *Styles
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

	// State management methods
	ID() string
	SetID(id string) Node
	GetState() NodeState
	SetState(state string, value bool) Node
	OnStateChange(state string, callback func(StateChange)) Node
	NotifyStateChange(state string, value bool)
}

// NodeState represents the current state of a node
type NodeState struct {
	IsHovered    bool
	IsActive     bool
	IsFocused    bool
	IsDisabled   bool
	CustomStates map[string]bool
}

// StateChange represents a state change event
type StateChange struct {
	NodeID string
	State  string
	Value  bool
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
	id             string
	state          NodeState
	stateListeners map[string][]func(StateChange)
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

func (n *BaseNode) AlignItems(value string) Node {
	n.styles.AlignItems = value
	n.styles.setProperties[string(AlignItemsProp)] = Explicit
	return n
}

func (n *BaseNode) FontWeight(value interface{}) Node {
	n.styles.FontWeight = parseStyleValue(value)
	n.styles.setProperties[string(FontWeightProp)] = Explicit
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
	// We catch the event nodes and style event nodes being added and rather than add them as children to this node,
	// we filter them out and incorporate their callbacks/styles into this node.
	// This allows for a more declarative UI construction where events and styles are treated as first-class nodes.
	var finalChildren []Node
	for _, child := range children {
		if eventNode, ok := child.(*EventNode); ok {
			// Register the event callback with this node
			n.eventCallbacks[eventNode.eventType] = eventNode.callback
		} else if styleNode, ok := child.(*BaseNode); ok && styleNode.nodeType == "style_handler" {
			// Incorporate the style changes into this node's styles
			if styleNode.styles.StateStyles != nil {
				for state, style := range styleNode.styles.StateStyles {
					n.styles.AddStateStyle(state, style)
				}
			}
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

// GetStyles returns a pointer to the node's styles
func (n *BaseNode) GetStyles() *Styles {
	return &n.styles
}

// Layout methods
func (n *BaseNode) ResolveStyles(parentStyles Styles) Styles {
	// Start with this node's styles
	resolvedStyles := *n.GetStyles()

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

	// Restore original styles if no states are active
	if !n.state.IsHovered && !n.state.IsActive && !n.state.IsFocused && !n.state.IsDisabled {
		resolvedStyles.RestoreOriginalStyles()
	} else {
		// Apply state-based styles
		if n.state.IsHovered && resolvedStyles.StateStyles != nil {
			if hoverStyle := resolvedStyles.StateStyles["hover"]; hoverStyle != nil {
				applyStateStyle(&resolvedStyles, hoverStyle)
			}
		}
		if n.state.IsActive && resolvedStyles.StateStyles != nil {
			if activeStyle := resolvedStyles.StateStyles["active"]; activeStyle != nil {
				applyStateStyle(&resolvedStyles, activeStyle)
			}
		}
		if n.state.IsFocused && resolvedStyles.StateStyles != nil {
			if focusStyle := resolvedStyles.StateStyles["focus"]; focusStyle != nil {
				applyStateStyle(&resolvedStyles, focusStyle)
			}
		}
		if n.state.IsDisabled && resolvedStyles.StateStyles != nil {
			if disabledStyle := resolvedStyles.StateStyles["disabled"]; disabledStyle != nil {
				applyStateStyle(&resolvedStyles, disabledStyle)
			}
		}
	}

	resolvedStyles.finalOpacity = resolvedStyles.Opacity * parentStyles.finalOpacity
	n.finalOpacity = resolvedStyles.finalOpacity

	// Apply the same process to all children
	for _, child := range n.children {
		child.ResolveStyles(resolvedStyles)
	}
	*n.GetStyles() = resolvedStyles
	return resolvedStyles
}

// applyStateStyle applies a state style variation to the base styles
func applyStateStyle(base *Styles, state *Styles) {
	// Store original values before applying state styles
	if base.originalValues == nil {
		base.originalValues = make(map[string]interface{})
	}

	// Apply all explicitly set properties from the state style
	for prop, source := range state.setProperties {
		if source == Explicit {
			// Store original value if not already stored
			if _, exists := base.originalValues[prop]; !exists {
				switch StyleProperty(prop) {
				case WidthProp:
					base.originalValues[prop] = base.Width
				case HeightProp:
					base.originalValues[prop] = base.Height
				case MarginProp:
					base.originalValues[prop] = base.Margin
				case PaddingProp:
					base.originalValues[prop] = base.Padding
				case FontSizeProp:
					base.originalValues[prop] = base.FontSize
				case ColorProp:
					base.originalValues[prop] = base.Color
				case OpacityProp:
					base.originalValues[prop] = base.Opacity
				case BackgroundProp:
					base.originalValues[prop] = base.Background
				case BorderProp:
					base.originalValues[prop] = base.Border
				case BorderRadiusProp:
					base.originalValues[prop] = base.BorderRadius
				case FlexDirectionProp:
					base.originalValues[prop] = base.FlexDirection
				case AlignItemsProp:
					base.originalValues[prop] = base.AlignItems
				case FontWeightProp:
					base.originalValues[prop] = base.FontWeight
				case ScaleProp:
					base.originalValues[prop] = state.Scale
				}
			}

			// Apply new value
			switch StyleProperty(prop) {
			case WidthProp:
				base.Width = state.Width
			case HeightProp:
				base.Height = state.Height
			case MarginProp:
				base.Margin = state.Margin
			case PaddingProp:
				base.Padding = state.Padding
			case FontSizeProp:
				base.FontSize = state.FontSize
			case ColorProp:
				base.Color = state.Color
			case OpacityProp:
				base.Opacity = state.Opacity
			case BackgroundProp:
				base.Background = state.Background
			case BorderProp:
				base.Border = state.Border
			case BorderRadiusProp:
				base.BorderRadius = state.BorderRadius
			case FlexDirectionProp:
				base.FlexDirection = state.FlexDirection
			case AlignItemsProp:
				base.AlignItems = state.AlignItems
			case FontWeightProp:
				base.FontWeight = state.FontWeight
			case ScaleProp:
				base.Scale = state.Scale
			}
			base.setProperties[prop] = Explicit
		}
	}

	// Handle scale separately since it's not a StyleProperty
	if state.Scale != 1.0 {
		if _, exists := base.originalValues["scale"]; !exists {
			base.originalValues["scale"] = base.Scale
		}
		base.Scale = state.Scale
		base.setProperties["scale"] = Explicit
	}
}

// RestoreOriginalStyles restores the original styles when a state is removed
func (s *Styles) RestoreOriginalStyles() {
	if s.originalValues == nil {
		return
	}

	for prop, value := range s.originalValues {
		switch prop {
		case string(WidthProp):
			s.Width = value.(StyleValue)
		case string(HeightProp):
			s.Height = value.(StyleValue)
		case string(MarginProp):
			s.Margin = value.(EdgeInsets)
		case string(PaddingProp):
			s.Padding = value.(EdgeInsets)
		case string(FontSizeProp):
			s.FontSize = value.(StyleValue)
		case string(ColorProp):
			s.Color = value.(Color)
		case string(OpacityProp):
			s.Opacity = value.(float64)
		case string(BackgroundProp):
			s.Background = value.(Color)
		case string(BorderProp):
			s.Border = value.(BorderStyle)
		case string(BorderRadiusProp):
			s.BorderRadius = value.(EdgeInsets)
		case string(FlexDirectionProp):
			s.FlexDirection = value.(string)
		case string(AlignItemsProp):
			s.AlignItems = value.(string)
		case string(FontWeightProp):
			s.FontWeight = value.(StyleValue)
		case "scale":
			s.Scale = value.(float64)
		}
		delete(s.setProperties, prop)
	}
	s.originalValues = nil
}

func (n *BaseNode) MeasurePreferred(ctx RenderContext) Size {
	// For leaf nodes, calculate intrinsic size
	if len(n.children) == 0 {
		return Size{0, 0} // Override in specific node types
	}

	if n.styles.Opacity == 0 {
		return Size{0, 0} // Don't render hidden objects
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

	// Store preferred size
	n.preferredSize = totalSize
	return totalSize
}

func (n *BaseNode) Layout(ctx RenderContext, constraints Constraints) Size {
	// Start with preferred size
	preferredSize := n.preferredSize

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

	// Store final size
	n.finalSize = finalSize
	return finalSize
}

func (n *BaseNode) ArrangeChildren(ctx RenderContext, bounds Rect) {
	// Set this node's bounds
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
	// Debug paint operation
	//debugLog("Painting %s (scale: %v, opacity: %v)", n.nodeType, n.styles.Scale, n.finalOpacity)

	// Apply scale transformation if needed
	if n.styles.Scale != 1.0 {
		// Calculate the center point of the node
		centerX := n.finalBounds.Position.X + n.finalBounds.Size.Width/2
		centerY := n.finalBounds.Position.Y + n.finalBounds.Size.Height/2

		// Calculate scaled bounds
		scaledWidth := n.finalBounds.Size.Width * n.styles.Scale
		scaledHeight := n.finalBounds.Size.Height * n.styles.Scale

		// Calculate new position to keep the center point
		newX := centerX - scaledWidth/2
		newY := centerY - scaledHeight/2

		// Create scaled bounds
		scaledBounds := Rect{
			Position: Point{X: newX, Y: newY},
			Size:     Size{Width: scaledWidth, Height: scaledHeight},
		}

		// Draw this node's background with scaled bounds
		ctx.DrawBackground(scaledBounds, n.styles, n.finalOpacity)

		// Draw borders if needed
		if n.styles.Border.CanDisplay() {
			ctx.DrawBorders(scaledBounds, n.styles, n.finalOpacity)
		}
	} else {
		// Draw this node's background
		ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

		// Draw borders if needed
		if n.styles.Border.CanDisplay() {
			ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
		}
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
	fmt.Printf("[DEBUG] Dispatching event %s to node %s\n", event.Type, n.nodeType)

	// Set the target if not already set
	if event.Target == nil {
		event.Target = n
	}

	// Handle built-in events
	switch event.Type {
	case UIEnter:
		fmt.Printf("[DEBUG] Setting hover state to true for node %s\n", n.nodeType)
		n.SetState("hover", true)
	case UILeave:
		fmt.Printf("[DEBUG] Setting hover state to false for node %s\n", n.nodeType)
		n.SetState("hover", false)
	case UIPress:
		fmt.Printf("[DEBUG] Setting active state to true for node %s\n", n.nodeType)
		n.SetState("active", true)
	case UIRelease:
		fmt.Printf("[DEBUG] Setting active state to false for node %s\n", n.nodeType)
		n.SetState("active", false)
	case UIFocus:
		n.SetState("focus", true)
	case UIBlur:
		n.SetState("focus", false)
	}

	// Call the handler if one exists
	if handler, has := n.eventCallbacks[event.Type]; has {
		fmt.Printf("[DEBUG] Calling event handler for %s on node %s\n", event.Type, n.nodeType)
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

// State management implementations for BaseNode
func (n *BaseNode) ID() string {
	return n.id
}

func (n *BaseNode) SetID(id string) Node {
	n.id = id
	return n
}

func (n *BaseNode) GetState() NodeState {
	return n.state
}

func (n *BaseNode) SetState(state string, value bool) Node {
	fmt.Printf("[DEBUG] Setting state %s=%v for node %s\n", state, value, n.nodeType)

	// Handle built-in states
	switch state {
	case "hover":
		n.state.IsHovered = value
		fmt.Printf("[DEBUG] Hover state is now %v for node %s\n", n.state.IsHovered, n.nodeType)
	case "active":
		n.state.IsActive = value
		fmt.Printf("[DEBUG] Active state is now %v for node %s\n", n.state.IsActive, n.nodeType)
	case "focus":
		n.state.IsFocused = value
	case "disabled":
		n.state.IsDisabled = value
	default:
		// Handle custom states
		if n.state.CustomStates == nil {
			n.state.CustomStates = make(map[string]bool)
		}
		n.state.CustomStates[state] = value
	}

	// Notify listeners
	n.NotifyStateChange(state, value)

	// Mark layout as dirty to trigger re-render
	if parent := n.Parent(); parent != nil {
		if renderEngine, ok := parent.(interface{ MarkLayoutDirty() }); ok {
			renderEngine.MarkLayoutDirty()
		}
	}

	return n
}

func (n *BaseNode) OnStateChange(state string, callback func(StateChange)) Node {
	if n.stateListeners == nil {
		n.stateListeners = make(map[string][]func(StateChange))
	}
	n.stateListeners[state] = append(n.stateListeners[state], callback)
	return n
}

func (n *BaseNode) NotifyStateChange(state string, value bool) {
	if n.stateListeners == nil {
		return
	}

	change := StateChange{
		NodeID: n.id,
		State:  state,
		Value:  value,
	}

	// Notify all listeners for this state
	if listeners, ok := n.stateListeners[state]; ok {
		for _, listener := range listeners {
			listener(change)
		}
	}

	// Also notify listeners for "all" state changes
	if listeners, ok := n.stateListeners["all"]; ok {
		for _, listener := range listeners {
			listener(change)
		}
	}
}

// Debug logging function
func debugLog(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}

// Debug style changes
func debugStyleChange(nodeType string, state string, props map[string]StyleSource) {
	debugLog("Style change for %s (state: %s):", nodeType, state)
	for prop, source := range props {
		debugLog("  - %s: %v", prop, source)
	}
}

// Debug state changes
func debugStateChange(nodeType string, state string, value bool) {
	debugLog("State change for %s: %s = %v", nodeType, state, value)
}

// Debug event handling
func debugEventHandling(nodeType string, eventType UIEventType) {
	debugLog("Event handling for %s: %s", nodeType, eventType)
}

// Add a method to dump node state for debugging
func (n *BaseNode) DumpState() {
	debugLog("Node State Dump:")
	debugLog("  Type: %s", n.nodeType)
	debugLog("  ID: %s", n.id)
	debugLog("  States:")
	debugLog("    Hovered: %v", n.state.IsHovered)
	debugLog("    Active: %v", n.state.IsActive)
	debugLog("    Focused: %v", n.state.IsFocused)
	debugLog("    Disabled: %v", n.state.IsDisabled)
	debugLog("  Styles:")
	debugLog("    Scale: %v", n.styles.Scale)
	debugLog("    Opacity: %v", n.styles.Opacity)
	debugLog("    Background: %v", n.styles.Background)
	debugLog("  Bounds:")
	debugLog("    Position: (%v,%v)", n.finalBounds.Position.X, n.finalBounds.Position.Y)
	debugLog("    Size: (%v,%v)", n.finalBounds.Size.Width, n.finalBounds.Size.Height)
	debugLog("  Children: %d", len(n.children))
}
