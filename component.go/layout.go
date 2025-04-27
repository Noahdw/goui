package component

import "github.com/noahdw/goui/node"

type VLayoutAlignment int

const (
	AlignLeft VLayoutAlignment = iota
	AlignRight
	AlignVCenter
)

type HLayoutAlignment int

const (
	AlignTop HLayoutAlignment = iota
	AlignBottom
	AlignHCenter
)

type VerticalBoxLayout struct {
	node.BaseNode
	spacing   float64
	alignment VLayoutAlignment
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
		} else if v.alignment == AlignRight {
			child.SetPositionX(v.BoundingRect().X + v.BoundingRect().Width) // wrong
		}
		childRect := child.BoundingRect()
		currentVerticalPosition += childRect.Height + v.spacing

		child.Render()
	}
}

func (v *VerticalBoxLayout) SetSpacing(spacing float64) {
	v.spacing = spacing
}

func (v *VerticalBoxLayout) SetAlignment(alignment VLayoutAlignment) {
	v.alignment = alignment
}

type HorizontalBoxLayout struct {
	node.BaseNode
	spacing   float64
	alignment HLayoutAlignment
}

func NewHorizontalBoxLayout() HorizontalBoxLayout {
	return HorizontalBoxLayout{
		spacing:   1,
		alignment: AlignTop,
	}
}

func (h *HorizontalBoxLayout) Render() {
	maxChildHeight := 0.0
	if h.alignment == AlignBottom {
		maxChildHeight = h.maxChildHeight()
	}
	currentHorizPosition := h.BoundingRect().Y
	for _, child := range h.Children() {
		child.SetPositionX(currentHorizPosition)
		if h.alignment == AlignTop {
			child.SetPositionY(h.BoundingRect().Y)
		} else if h.alignment == AlignBottom {
			child.SetPositionY(maxChildHeight - child.BoundingRect().Height)
		}
		childRect := child.BoundingRect()
		currentHorizPosition += childRect.Width + h.spacing

		child.Render()
	}
}

func (h *HorizontalBoxLayout) SetSpacing(spacing float64) {
	h.spacing = spacing
}

func (h *HorizontalBoxLayout) SetAlignment(alignment HLayoutAlignment) {
	h.alignment = alignment
}

func (h *HorizontalBoxLayout) maxChildHeight() float64 {
	maxHeight := 0.0
	for _, child := range h.Children() {
		maxHeight = max(maxHeight, child.BoundingRect().Height)
	}
	return maxHeight
}
