package ui

import n "github.com/noahdw/goui/node"

// H1 creates a heading level 1 node
func H1(children ...n.Node) n.Node {
	node := n.NewBaseNode("h1", n.NewStyles(n.StyleProps{
		FontSize:   &n.StyleValue{Type: n.PIXEL, Value: 24, Source: n.Default},
		FontWeight: &n.StyleValue{Type: n.PIXEL, Value: 700, Source: n.Default},
		Margin:     &n.EdgeInsets{Top: 16, Right: 0, Bottom: 8, Left: 0},
	}))
	node.AddChildren(children...)
	return &node
}

// H2 creates a heading level 2 node
func H2(children ...n.Node) n.Node {
	node := n.NewBaseNode("h2", n.NewStyles(n.StyleProps{
		FontSize:   &n.StyleValue{Type: n.PIXEL, Value: 20, Source: n.Default},
		FontWeight: &n.StyleValue{Type: n.PIXEL, Value: 700, Source: n.Default},
		Margin:     &n.EdgeInsets{Top: 14, Right: 0, Bottom: 7, Left: 0},
	}))
	node.AddChildren(children...)
	return &node
}

// Text creates a text node with the given content
func Text(text string) n.Node {
	node := n.NewBaseNode("text", n.NewStyles(n.StyleProps{
		FontSize:   &n.StyleValue{Type: n.PIXEL, Value: 20, Source: n.Default},
		FontWeight: &n.StyleValue{Type: n.PIXEL, Value: 700, Source: n.Default},
		Color:      &n.Black,
	}))
	return n.NewTextNode(node, text)
}

// Button creates a button node with the given children
func Button(children ...n.Node) n.Node {
	node := n.NewBaseNode("button", n.NewStyles(n.StyleProps{
		Padding:      &n.EdgeInsets{Top: 8, Right: 16, Bottom: 8, Left: 16},
		Background:   &n.Gray,
		Color:        &n.White,
		BorderRadius: &n.EdgeInsets{Top: 4, Right: 4, Bottom: 4, Left: 4},
	}))
	node.AddChildren(children...)
	return &node
}

// Layout creates a layout node with specified direction and children
func Layout(direction string, children ...n.Node) n.Node {
	node := n.NewBaseNode("layout", n.NewStyles(n.StyleProps{
		FlexDirection: &direction,
	}))
	node.AddChildren(children...)
	return &node
}

// Image creates an image node with the given source URL
func Image(sourceURL string) n.Node {
	node := n.NewBaseNode("image", n.NewStyles(n.StyleProps{}))
	return n.NewImageNode(node, sourceURL)
}

// Rect creates a rectangle node with the given children
func Rect(children ...n.Node) n.Node {
	flexDirection := "row"
	node := n.NewBaseNode("rect", n.NewStyles(n.StyleProps{
		Padding:       &n.EdgeInsets{Top: 8, Right: 8, Bottom: 8, Left: 8},
		Background:    &n.Gray,
		Color:         &n.White,
		FlexDirection: &flexDirection,
	}))
	node.AddChildren(children...)
	return &node
}
