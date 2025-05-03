package node

func H1(children ...Node) Node {
	node := &BaseNode{
		nodeType: "h1",
		styles: NewStyles(StyleProps{
			FontSize:   &StyleValue{Type: PIXEL, Value: 24, Source: Default},
			FontWeight: &StyleValue{Type: PIXEL, Value: 700, Source: Default},
			Margin:     &EdgeInsets{16, 0, 8, 0},
		}),
	}
	node.AddChildren(children...)
	return node
}

func H2(children ...Node) Node {
	node := &BaseNode{
		nodeType: "h2",
		styles: NewStyles(StyleProps{
			FontSize:   &StyleValue{Type: PIXEL, Value: 20, Source: Default},
			FontWeight: &StyleValue{Type: PIXEL, Value: 700, Source: Default},
			Margin:     &EdgeInsets{14, 0, 7, 0},
		}),
	}
	node.AddChildren(children...)
	return node
}

func Text(text string) Node {
	node := &TextNode{
		BaseNode: BaseNode{
			nodeType: "text",
			styles: NewStyles(StyleProps{
				FontSize:   &StyleValue{Type: PIXEL, Value: 20, Source: Default},
				FontWeight: &StyleValue{Type: PIXEL, Value: 700, Source: Default},
				Color:      &Black,
			}),
		},
		text: text,
	}
	return node
}

func Button(children ...Node) Node {
	node := &BaseNode{
		nodeType: "button",
		styles: NewStyles(StyleProps{
			Padding:      &EdgeInsets{8, 16, 8, 16},
			Background:   &Gray,
			Color:        &White,
			BorderRadius: &EdgeInsets{4, 4, 4, 4},
		}),
	}
	node.AddChildren(children...)
	return node
}

func Layout(direction string, children ...Node) Node {
	node := &BaseNode{
		nodeType: "layout",
		styles: NewStyles(StyleProps{
			FlexDirection: &direction,
		}),
	}
	node.AddChildren(children...)
	return node
}
