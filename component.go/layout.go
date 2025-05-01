package component

import (
	"github.com/noahdw/goui/bounds"
	"github.com/noahdw/goui/node"
)

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
	currentVerticalPosition := v.PositionY()
	for _, child := range v.Children() {
		child.SetPositionY(currentVerticalPosition)

		if v.alignment == AlignLeft {
			child.SetPositionX(v.PositionX())
		} else if v.alignment == AlignRight {
			child.SetPositionX(v.PositionX() + v.Width()) // wrong
		}

		currentVerticalPosition += child.Height() + v.spacing
	}
	v.BaseNode.Render()

}

func (v *VerticalBoxLayout) SetSpacing(spacing float64) {
	v.spacing = spacing
}

func (v *VerticalBoxLayout) SetAlignment(alignment VLayoutAlignment) {
	v.alignment = alignment
}

func (v *VerticalBoxLayout) BoundingRect() bounds.Bounds {
	totalWidth := 0.
	totalHeight := 0.
	for _, child := range v.Children() {
		totalWidth += child.Width()
		totalHeight += child.Height()
	}
	v.SetWidth(v.MaxChildWidth())
	v.SetHeight(totalHeight + float64(len(v.Children())-1)*v.spacing)
	return v.BaseNode.BoundingRect()
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
		maxChildHeight = h.MaxChildHeight()
	}
	y := h.PositionY()
	currentHorizPosition := y
	for _, child := range h.Children() {
		child.SetPositionX(currentHorizPosition)
		if h.alignment == AlignTop {
			child.SetPositionY(y)
		} else if h.alignment == AlignBottom {
			child.SetPositionY(maxChildHeight - child.Height())
		}
		currentHorizPosition += child.Width() + h.spacing
	}
	h.BaseNode.Render()
}

func (h *HorizontalBoxLayout) SetSpacing(spacing float64) {
	h.spacing = spacing
}

func (h *HorizontalBoxLayout) SetAlignment(alignment HLayoutAlignment) {
	h.alignment = alignment
}

func (h *HorizontalBoxLayout) BoundingRect() bounds.Bounds {
	totalHeight := 0.
	totalWidth := 0.

	for _, child := range h.Children() {
		totalHeight += child.BoundingRect().Height
		totalWidth += child.BoundingRect().Width
	}

	h.SetWidth(totalWidth + float64(len(h.Children())-1)*h.spacing)
	h.SetHeight(h.MaxChildHeight())
	return h.BaseNode.BoundingRect()
}
