package component

import "github.com/noahdw/goui/node"

type LayoutAlignment int

const (
	AlignLeft LayoutAlignment = iota
	AlignRight
	AlignCenter
)

type VerticalBoxLayout struct {
	node.BaseNode
	spacing   float64
	alignment LayoutAlignment
}

func NewVerticalBoxLayout() VerticalBoxLayout {
	return VerticalBoxLayout{
		spacing:   1,
		alignment: AlignLeft,
	}
}

func (v *VerticalBoxLayout) Render() {
	currentVerticalPosition := v.BoundingRect().Y
	for _, child := range v.Children() {
		child.SetPositionY(currentVerticalPosition)
		if v.alignment == AlignLeft {
			child.SetPositionX(v.BoundingRect().X)
		}
		childRect := child.BoundingRect()
		currentVerticalPosition += childRect.Height + v.spacing

		child.Render()
	}
}

func (v *VerticalBoxLayout) SetSpacing(spacing float64) {
	v.spacing = spacing
}

func (v *VerticalBoxLayout) SetAlignment(alignment LayoutAlignment) {
	v.alignment = alignment
}
