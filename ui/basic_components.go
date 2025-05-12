package ui

import (
	"fmt"

	n "github.com/noahdw/goui/node"
)

// H1 creates a heading level 1 node
func H1(children ...n.Node) n.Node {
	textAlign := "center"
	alignItems := "center"
	props := map[string]interface{}{
		"fontSize":   &n.StyleValue{Type: n.PIXEL, Value: 24, Source: n.Default},
		"fontWeight": &n.StyleValue{Type: n.PIXEL, Value: 700, Source: n.Default},
		"margin":     &n.EdgeInsets{Top: 16, Right: 0, Bottom: 8, Left: 0},
		"textAlign":  &textAlign,
		"alignItems": &alignItems,
	}
	node := n.NewBaseNode("h1", n.NewStyles(props))
	node.AddChildren(children...)
	return &node
}

// H2 creates a heading level 2 node
func H2(children ...n.Node) n.Node {
	textAlign := "center"
	alignItems := "center"
	props := map[string]interface{}{
		"fontSize":   &n.StyleValue{Type: n.PIXEL, Value: 20, Source: n.Default},
		"fontWeight": &n.StyleValue{Type: n.PIXEL, Value: 700, Source: n.Default},
		"margin":     &n.EdgeInsets{Top: 14, Right: 0, Bottom: 7, Left: 0},
		"textAlign":  &textAlign,
		"alignItems": &alignItems,
	}
	node := n.NewBaseNode("h2", n.NewStyles(props))
	node.AddChildren(children...)
	return &node
}

// Text creates a text node with the given content
func Text(text string) n.Node {
	textAlign := "left"
	alignItems := "center"
	props := map[string]interface{}{
		"fontSize":   &n.StyleValue{Type: n.PIXEL, Value: 16, Source: n.Default},
		"fontWeight": &n.StyleValue{Type: n.PIXEL, Value: 400, Source: n.Default},
		"color":      &n.Black,
		"textAlign":  &textAlign,
		"alignItems": &alignItems,
	}
	node := n.NewBaseNode("text", n.NewStyles(props))
	return n.NewTextNode(node, text)
}

// Button creates a button node with the given children
func Button(children ...n.Node) n.Node {
	props := map[string]interface{}{
		"padding":      n.EdgeInsets{Top: 8, Right: 16, Bottom: 8, Left: 16},
		"background":   n.Gray,
		"color":        n.Black,
		"borderRadius": n.EdgeInsets{Top: 4, Right: 4, Bottom: 4, Left: 4},
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
		"padding":       n.EdgeInsets{Top: 4, Right: 4, Bottom: 4, Left: 4},
		"background":    n.Gray,
		"color":         n.White,
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
//	  StyleOnEvent("hover", &n.StyleProps{
//	    Background: &n.Color{240, 240, 240, 255},
//	    Scale: &1.1,
//	  }),
//	  Text("Hover me"),
//	)
func StyleOnEvent(state string, style *n.StyleProps) n.Node {
	// Create a base node that will be absorbed by its parent
	props := map[string]interface{}{}
	baseNode := n.NewBaseNode("style_handler", n.NewStyles(props))
	node := &baseNode

	// Convert StyleProps to map[string]interface{}
	stateProps := make(map[string]interface{})
	if style.Background != nil {
		stateProps["background"] = style.Background
	}
	if style.Color != nil {
		stateProps["color"] = style.Color
	}
	if style.Scale != nil {
		stateProps["scale"] = style.Scale
	}

	// Create and add the state style
	stateStyles := n.NewStyles(stateProps)

	// Ensure background color is marked as explicitly set if provided
	if style.Background != nil {
		stateStyles.MarkPropertyExplicit(n.BackgroundProp)
	}
	if style.Color != nil {
		stateStyles.MarkPropertyExplicit(n.ColorProp)
	}

	node.GetStyles().AddStateStyle(state, &stateStyles)

	// Debug logging
	fmt.Printf("[DEBUG] Added style for state %s with properties: %+v\n", state, style)

	return node
}
