package ui

import (
	"fmt"

	n "github.com/noahdw/goui/node"
	"github.com/noahdw/goui/node/style"
)

// H1 creates a heading level 1 node
func H1(children ...n.Node) n.Node {
	props := map[string]interface{}{
		"fontSize":   style.StyleValue{Type: style.PIXEL, Value: 24, Source: style.Default},
		"fontWeight": style.StyleValue{Type: style.PIXEL, Value: 700, Source: style.Default},
		"margin":     style.EdgeInsets{Top: 16, Right: 0, Bottom: 8, Left: 0},
		"textAlign":  "center",
		"alignItems": "center",
	}
	node := n.NewBaseNode("h1", style.NewStyles(props))
	node.AddChildren(children...)
	return &node
}

// H2 creates a heading level 2 node
func H2(children ...n.Node) n.Node {
	props := map[string]interface{}{
		"fontSize":   style.StyleValue{Type: style.PIXEL, Value: 20, Source: style.Default},
		"fontWeight": style.StyleValue{Type: style.PIXEL, Value: 700, Source: style.Default},
		"margin":     style.EdgeInsets{Top: 14, Right: 0, Bottom: 7, Left: 0},
		"textAlign":  "center",
		"alignItems": "center",
	}
	node := n.NewBaseNode("h2", style.NewStyles(props))
	node.AddChildren(children...)
	return &node
}

// Text creates a text node with the given content
func Text(text string) n.Node {
	props := map[string]interface{}{
		"fontSize":   style.StyleValue{Type: style.PIXEL, Value: 16, Source: style.Default},
		"fontWeight": style.StyleValue{Type: style.PIXEL, Value: 400, Source: style.Default},
		"color":      style.Black,
		"textAlign":  "left",
		"alignItems": "center",
	}
	node := n.NewBaseNode("text", style.NewStyles(props))
	return n.NewTextNode(node, text)
}

// Button creates a button node with the given children
func Button(children ...n.Node) n.Node {
	props := map[string]interface{}{
		"padding":      style.EdgeInsets{Top: 8, Right: 16, Bottom: 8, Left: 16},
		"background":   style.Gray,
		"color":        style.Black,
		"borderRadius": style.EdgeInsets{Top: 4, Right: 4, Bottom: 4, Left: 4},
	}
	node := n.NewBaseNodeWithProps("button", props)
	node.AddChildren(children...)
	return node
}

// Layout creates a layout node with specified direction and children
func Layout(direction string, children ...n.Node) n.Node {
	props := map[string]interface{}{
		"flexDirection": direction,
	}
	node := n.NewBaseNodeWithProps("layout", props)
	node.AddChildren(children...)
	return node
}

// Image creates an image node with the given source URL
func Image(sourceURL string) n.Node {
	props := map[string]interface{}{}
	node := n.NewBaseNodeWithProps("image", props)
	return n.NewImageNode(node, sourceURL)
}

// Rect creates a rectangle node with the given children
func Rect(children ...n.Node) n.Node {
	props := map[string]interface{}{
		"padding":       style.EdgeInsets{Top: 4, Right: 4, Bottom: 4, Left: 4},
		"background":    style.Gray,
		"color":         style.White,
		"flexDirection": "row",
	}
	node := n.NewBaseNodeWithProps("rect", props)
	node.AddChildren(children...)
	return node
}

// OnEvent creates an event handler node
func OnEvent(eventType n.UIEventType, callback func(n.UIEvent)) n.Node {
	props := map[string]interface{}{}
	node := n.NewBaseNodeWithProps("event_handler", props)
	return n.NewEventNode(node, eventType, callback)
}

// StyleOnEvent creates a style change handler that will be absorbed by its parent.
// The style handler will be registered with the parent node and will not appear in the node tree.
// This allows for a declarative approach to style management where style changes are treated as first-class nodes.
//
// Example:
//
//	Rect(
//	  StyleOnEvent("hover", style.StyleProps{
//	    Background: style.Color{240, 240, 240, 255},
//	    Scale: 1.1,
//	  }),
//	  Text("Hover me"),
//	)
func StyleOnEvent(state string, styleProps style.StyleProps) n.Node {
	// Create a base node that will be absorbed by its parent
	props := map[string]interface{}{}
	baseNode := n.NewBaseNode("style_handler", style.NewStyles(props))
	node := &baseNode

	// Convert StyleProps to map[string]interface{}
	stateProps := make(map[string]interface{})
	if styleProps.Background != nil {
		stateProps["background"] = *styleProps.Background
	}
	if styleProps.Color != nil {
		stateProps["color"] = *styleProps.Color
	}
	if styleProps.Scale != nil {
		stateProps["scale"] = *styleProps.Scale
	}

	// Create and add the state style
	stateStyles := style.NewStyles(stateProps)

	// Ensure background color is marked as explicitly set if provided
	if styleProps.Background != nil {
		stateStyles.MarkPropertyExplicit(style.BackgroundProp)
	}
	if styleProps.Color != nil {
		stateStyles.MarkPropertyExplicit(style.ColorProp)
	}

	// Pass a pointer to stateStyles
	node.GetStyles().AddStateStyle(state, &stateStyles)

	// Debug logging
	fmt.Printf("[DEBUG] Added style for state %s with properties: %+v\n", state, styleProps)

	return node
}
