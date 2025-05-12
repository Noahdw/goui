package node

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/node/style"
)

type Node interface {
	StyleBuilder // Include all style builder methods

	// Style methods
	GetStyle(key string) (interface{}, bool)
	GetStyleFloat(key string) (float64, bool)
	GetStyleString(key string) (string, bool)
	GetStyleColor(key string) (style.Color, bool)
	GetStyleEdgeInsets(key string) (style.EdgeInsets, bool)

	// Structure methods
	AddChildren(children ...Node)
	Children() []Node
	GetType() string
	GetStyles() *style.Styles
	Intersects(node Node) bool
	Parent() Node
	SetParent(parent Node)
	DispatchEvent(event UIEvent)

	// Layout methods
	ResolveStyles(parentStyles style.Styles) style.Styles
	MeasurePreferred(ctx RenderContext) style.Size
	Layout(ctx RenderContext, constraints Constraints) style.Size
	ArrangeChildren(ctx RenderContext, bounds style.Rect)
	Paint(ctx RenderContext)
	GetFinalSize() style.Size
	GetFinalBounds() style.Rect

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
	styles         style.Styles
	finalSize      style.Size
	finalBounds    style.Rect
	preferredSize  style.Size
	finalOpacity   float64
	eventCallbacks map[UIEventType]func(UIEvent)
	id             string
	state          NodeState
	stateListeners map[string][]func(StateChange)
}

type Event struct {
}

func NewBaseNode(nodeType string, styles style.Styles) BaseNode {
	return BaseNode{
		nodeType:       nodeType,
		styles:         styles,
		eventCallbacks: make(map[UIEventType]func(UIEvent)),
	}
}

// NewBaseNodeWithStyles creates a new base node with the given styles
func NewBaseNodeWithStyles(nodeType string, styles style.Styles) Node {
	return &BaseNode{
		nodeType:       nodeType,
		styles:         styles,
		eventCallbacks: make(map[UIEventType]func(UIEvent)),
	}
}

// NewBaseNodeWithProps creates a new base node with the given properties
func NewBaseNodeWithProps(nodeType string, props map[string]interface{}) Node {
	styles := style.NewStyles(props)

	// If there were any style errors, create an error node to display them
	if errorVal, ok := styles.Get("error"); ok {
		if errorMsg, ok := errorVal.(string); ok {
			errorStyles := style.NewStyles(map[string]interface{}{
				"background": style.Red,
				"color":      style.White,
				"padding":    style.EdgeInsets{10, 10, 10, 10},
				"width":      400,
				"height":     100,
			})
			errorBase := BaseNode{
				nodeType:       "error",
				styles:         errorStyles,
				eventCallbacks: make(map[UIEventType]func(UIEvent)),
			}
			errorText := TextNode{
				BaseNode: errorBase,
				text:     fmt.Sprintf("Style Error: %s", errorMsg),
			}
			return &errorText
		}
	}

	return &BaseNode{
		nodeType:       nodeType,
		styles:         styles,
		eventCallbacks: make(map[UIEventType]func(UIEvent)),
	}
}

// Constraints represents size constraints
type Constraints struct {
	MinWidth, MaxWidth, MinHeight, MaxHeight float64
}

// RenderContext provides context for rendering
type RenderContext interface {
	LoadTexture(sourceURL string) rl.Texture2D
	Present()
	Save()
	Restore()
	SetOpacity(opacity float64)
	SetFillColor(color style.Color)
	SetStrokeColor(color style.Color)
	SetLineWidth(width float64)
	StrokeLine(start style.Point, end style.Point)
	SetFontSize(size float64)
	DrawText(text string, bounds style.Rect, styles style.Styles, opacity float64)
	DrawBackground(bounds style.Rect, styles style.Styles, opacity float64)
	DrawBorders(bounds style.Rect, styles style.Styles, opacity float64)
	DrawTexture(sourceURL string, bounds style.Rect, styles style.Styles, opacity float64)
	FillRect(rect style.Rect)
	Scale(x, y float64)
	Clear()
	ClipRect() style.Rect
	SetClipRect(rect style.Rect)
}

func (n *BaseNode) GetStyle(key string) (interface{}, bool) {
	return n.styles.Get(key)
}

func (n *BaseNode) GetStyleFloat(key string) (float64, bool) {
	return n.styles.GetFloat(key)
}

func (n *BaseNode) GetStyleString(key string) (string, bool) {
	return n.styles.GetString(key)
}

func (n *BaseNode) GetStyleColor(key string) (style.Color, bool) {
	return n.styles.GetColor(key)
}

func (n *BaseNode) GetStyleEdgeInsets(key string) (style.EdgeInsets, bool) {
	return n.styles.GetEdgeInsets(key)
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
			if stateStyles := styleNode.styles.GetStateStyle("all"); stateStyles != nil {
				n.styles.AddStateStyle("all", stateStyles)
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
func (n *BaseNode) GetStyles() *style.Styles {
	return &n.styles
}

// Layout methods
func (n *BaseNode) ResolveStyles(parentStyles style.Styles) style.Styles {
	// Start with this node's styles
	resolvedStyles := *n.GetStyles()

	// For inheritable properties, check if they're set in this node
	// If not, inherit from parent
	inheritableProps := []string{
		"fontFamily", "fontSize", "color", "lineHeight", "background", "opacity",
	}

	for _, prop := range inheritableProps {
		// Only inherit if:
		// 1. Parent has the property set (explicitly or inherited)
		// 2. This node doesn't have it explicitly set
		_, parentHasIt := parentStyles.Get(prop)
		selfHasIt := resolvedStyles.IsExplicit(prop)

		if parentHasIt && !selfHasIt {
			if value, ok := parentStyles.Get(prop); ok {
				resolvedStyles.Set(prop, value)
			}
		}
	}

	// Restore original styles if no states are active
	if !n.state.IsHovered && !n.state.IsActive && !n.state.IsFocused && !n.state.IsDisabled {
		resolvedStyles.RestoreOriginalStyles()
	} else {
		// Apply state-based styles
		if n.state.IsHovered {
			if hoverStyle := resolvedStyles.GetStateStyle("hover"); hoverStyle != nil {
				applyStateStyle(&resolvedStyles, hoverStyle)
			}
		}
		if n.state.IsActive {
			if activeStyle := resolvedStyles.GetStateStyle("active"); activeStyle != nil {
				applyStateStyle(&resolvedStyles, activeStyle)
			}
		}
		if n.state.IsFocused {
			if focusStyle := resolvedStyles.GetStateStyle("focus"); focusStyle != nil {
				applyStateStyle(&resolvedStyles, focusStyle)
			}
		}
		if n.state.IsDisabled {
			if disabledStyle := resolvedStyles.GetStateStyle("disabled"); disabledStyle != nil {
				applyStateStyle(&resolvedStyles, disabledStyle)
			}
		}
	}

	// Calculate final opacity
	if opacity, ok := resolvedStyles.GetFloat("opacity"); ok {
		if parentOpacity, ok := parentStyles.GetFloat("opacity"); ok {
			resolvedStyles.SetFinalOpacity(opacity * parentOpacity)
		} else {
			resolvedStyles.SetFinalOpacity(opacity)
		}
	} else {
		resolvedStyles.SetFinalOpacity(parentStyles.GetFinalOpacity())
	}
	n.finalOpacity = resolvedStyles.GetFinalOpacity()

	// Apply the same process to all children
	for _, child := range n.children {
		child.ResolveStyles(resolvedStyles)
	}
	*n.GetStyles() = resolvedStyles
	return resolvedStyles
}

// applyStateStyle applies a state style variation to the base styles
func applyStateStyle(base *style.Styles, state *style.Styles) {
	// Apply all explicitly set properties from the state style
	inheritableProps := []string{
		"fontFamily", "fontSize", "color", "lineHeight", "background", "opacity",
	}
	for _, prop := range inheritableProps {
		if state.IsExplicit(prop) {
			// Store original value if not already stored
			if value, ok := base.Get(prop); ok {
				base.StoreOriginalValue(prop, value)
			}
			// Apply new value
			if value, ok := state.Get(prop); ok {
				base.Set(prop, value)
			}
		}
	}
}

func (n *BaseNode) Layout(ctx RenderContext, constraints Constraints) style.Size {
	// Start with preferred size
	preferredSize := n.preferredSize

	// Apply constraints to determine final size
	finalSize := style.Size{
		Width:  clamp(preferredSize.Width, constraints.MinWidth, constraints.MaxWidth),
		Height: clamp(preferredSize.Height, constraints.MinHeight, constraints.MaxHeight),
	}

	// Handle special cases like Fill Parent
	if width, ok := n.styles.Get("width"); ok {
		if widthValue, ok := width.(style.StyleValue); ok && widthValue.Type == style.PERCENTAGE {
			finalSize.Width = constraints.MaxWidth * widthValue.Value.(float64) / 100
		}
	}

	if height, ok := n.styles.Get("height"); ok {
		if heightValue, ok := height.(style.StyleValue); ok && heightValue.Type == style.PERCENTAGE {
			finalSize.Height = constraints.MaxHeight * heightValue.Value.(float64) / 100
		}
	}

	// If we have children, layout them now
	if len(n.children) > 0 {
		// Calculate available space for children (minus padding)
		availableWidth := finalSize.Width
		availableHeight := finalSize.Height

		if padding, ok := n.styles.GetEdgeInsets("padding"); ok {
			availableWidth -= padding.Left + padding.Right
			availableHeight -= padding.Top + padding.Bottom
		}

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

func (n *BaseNode) ArrangeChildren(ctx RenderContext, bounds style.Rect) {
	// Set this node's bounds
	n.finalBounds = bounds

	// For leaf nodes, we're done
	if len(n.children) == 0 {
		return
	}

	// Calculate content area (bounds minus padding)
	contentArea := style.Rect{
		Position: style.Point{
			X: bounds.Position.X,
			Y: bounds.Position.Y,
		},
		Size: style.Size{
			Width:  bounds.Size.Width,
			Height: bounds.Size.Height,
		},
	}

	// Apply padding if present
	if padding, ok := n.styles.GetEdgeInsets("padding"); ok {
		contentArea.Position.X += padding.Left
		contentArea.Position.Y += padding.Top
		contentArea.Size.Width -= padding.Left + padding.Right
		contentArea.Size.Height -= padding.Top + padding.Bottom
	}

	// Position each child based on layout direction
	var currentX = contentArea.Position.X
	var currentY = contentArea.Position.Y

	// Get flex direction
	flexDir := "row" // default
	if dir, ok := n.styles.GetString("flexDirection"); ok {
		flexDir = dir
	}

	for _, child := range n.children {
		childSize := child.GetFinalSize()
		childStyles := child.GetStyles()

		// Get margins for current child
		var margin style.EdgeInsets
		if m, ok := childStyles.GetEdgeInsets("margin"); ok {
			margin = m
		}

		if flexDir == "row" {
			// Position child horizontally
			childBounds := style.Rect{
				Position: style.Point{
					X: currentX + margin.Left,
					Y: currentY + margin.Top,
				},
				Size: childSize,
			}
			child.ArrangeChildren(ctx, childBounds)

			// Update currentX for next child
			currentX += childSize.Width + margin.Left + margin.Right
		} else { // column
			// Position child vertically
			childBounds := style.Rect{
				Position: style.Point{
					X: currentX + margin.Left,
					Y: currentY + margin.Top,
				},
				Size: childSize,
			}
			child.ArrangeChildren(ctx, childBounds)

			// Update currentY for next child
			currentY += childSize.Height + margin.Top + margin.Bottom
		}
	}
}

func (n *BaseNode) GetFinalSize() style.Size {
	return n.finalSize
}

func (n *BaseNode) GetFinalBounds() style.Rect {
	return n.finalBounds
}

func (n *BaseNode) Paint(ctx RenderContext) {
	// Apply opacity if set
	opacity, _ := n.styles.GetFloat("opacity")
	if opacity < 1.0 {
		ctx.Save()
		ctx.SetOpacity(opacity)
		defer ctx.Restore()
	}

	// Apply scale if set
	scale, _ := n.styles.GetFloat("scale")
	if scale != 1.0 {
		ctx.Save()
		ctx.Scale(scale, scale)
		defer ctx.Restore()
	}

	// Draw background if set
	if bgColor, ok := n.styles.GetColor("background"); ok {
		ctx.SetFillColor(bgColor)
		ctx.DrawBackground(n.finalBounds, n.styles, opacity)
	}

	// Draw border if set
	if border, ok := n.styles.Get("border"); ok {
		if borderStyle, ok := border.(style.BorderStyle); ok && borderStyle.CanDisplay() {
			ctx.DrawBorders(n.finalBounds, n.styles, opacity)
		}
	}

	// Draw text if set
	if text, ok := n.styles.GetString("text"); ok {
		if fontSize, ok := n.styles.GetFloat("fontSize"); ok {
			if textColor, ok := n.styles.GetColor("color"); ok {
				ctx.SetFontSize(fontSize)
				ctx.SetFillColor(textColor)
				ctx.DrawText(text, n.finalBounds, n.styles, opacity)
			}
		}
	}

	// Paint children
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
func (n *TextNode) MeasurePreferred(ctx RenderContext) style.Size {
	fontSize, _ := n.styles.GetFloat("fontSize")
	padding, _ := n.styles.GetEdgeInsets("padding")
	textWidth := float64(rl.MeasureText(n.text, int32(fontSize)))
	textHeight := fontSize * 1.2 // Add some line height

	// Add padding to the text size
	return style.Size{
		Width:  textWidth + padding.Left + padding.Right,
		Height: textHeight + padding.Top + padding.Bottom,
	}
}

func (n *TextNode) Paint(ctx RenderContext) {
	// Draw background and border first
	ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

	// Draw borders if needed
	if border, ok := n.styles.Get("border"); ok {
		if borderStyle, ok := border.(style.BorderStyle); ok && borderStyle.CanDisplay() {
			ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
		}
	}

	// Then draw the text
	ctx.DrawText(n.text, n.finalBounds, n.styles, n.finalOpacity)
}

type ImageNode struct {
	BaseNode
	sourceURL string
}

func NewImageNode(baseNode Node, sourceURL string) Node {
	if baseNode == nil {
		return nil
	}
	// Convert the Node to BaseNode
	base, ok := baseNode.(*BaseNode)
	if !ok {
		// If we can't convert, create an error node
		errorNode := NewBaseNodeWithProps("error", map[string]interface{}{
			"background": style.Red,
			"color":      style.White,
			"padding":    style.EdgeInsets{10, 10, 10, 10},
			"width":      400,
			"height":     100,
		})
		if errorText, ok := errorNode.(*TextNode); ok {
			errorText.text = "Invalid node type for ImageNode"
		}
		return errorNode
	}
	imageNode := ImageNode{
		BaseNode:  *base,
		sourceURL: sourceURL,
	}
	return &imageNode
}

type EventNode struct {
	BaseNode
	callback  func(UIEvent)
	eventType UIEventType
}

func NewEventNode(baseNode Node, eventType UIEventType, callback func(UIEvent)) Node {
	if baseNode == nil {
		return nil
	}
	// Convert the Node to BaseNode
	base, ok := baseNode.(*BaseNode)
	if !ok {
		// If we can't convert, create an error node
		errorNode := NewBaseNodeWithProps("error", map[string]interface{}{
			"background": style.Red,
			"color":      style.White,
			"padding":    style.EdgeInsets{10, 10, 10, 10},
			"width":      400,
			"height":     100,
		})
		if errorText, ok := errorNode.(*TextNode); ok {
			errorText.text = "Invalid node type for EventNode"
		}
		return errorNode
	}
	eventNode := EventNode{
		BaseNode:  *base,
		callback:  callback,
		eventType: eventType,
	}
	return &eventNode
}

func (n *ImageNode) Paint(ctx RenderContext) {
	// Draw background and border first
	ctx.DrawBackground(n.finalBounds, n.styles, n.finalOpacity)

	// Draw borders if needed
	if border, ok := n.styles.Get("border"); ok {
		if borderStyle, ok := border.(style.BorderStyle); ok && borderStyle.CanDisplay() {
			ctx.DrawBorders(n.finalBounds, n.styles, n.finalOpacity)
		}
	}

	// Then draw the image
	ctx.DrawTexture(n.sourceURL, n.finalBounds, n.styles, n.finalOpacity)
}

func (n *ImageNode) MeasurePreferred(ctx RenderContext) style.Size {
	texture := ctx.LoadTexture(n.sourceURL)
	n.preferredSize = style.Size{
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

func (n *BaseNode) MeasurePreferred(ctx RenderContext) style.Size {
	// Default implementation returns zero size
	return style.Size{Width: 0, Height: 0}
}
