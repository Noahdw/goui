package component

import (
	"image/color"

	. "github.com/noahdw/goui/bounds"
)

type Relational interface {
	AddChild(Component)
	removeChild(Component)
	Children() []Component
	Parent() Component
	SetParent(Component)
}

type BaseRelation struct {
	children []Component
	parent   Component
}

func (b *BaseRelation) AddChild(child Component) {
	b.children = append(b.children, child)
}

func (b *BaseRelation) removeChild(child Component) {

}

func (b *BaseRelation) Children() []Component {
	return b.children
}

func (b *BaseRelation) Parent() Component {
	return b.parent
}

func (b *BaseRelation) SetParent(parent Component) {
	b.parent = parent
}

type MouseHandler interface {
	HandleMouse(MouseEvent) EventHandleState
}

type BaseMouseHandler struct {
	mouseEventHandler func(MouseEvent) EventHandleState
}

func (b *BaseMouseHandler) HandleMouse(event MouseEvent) EventHandleState {
	if b.mouseEventHandler != nil {
		return b.mouseEventHandler(event)
	}
	return Propogate
}

func (b *BaseMouseHandler) OnMouseEvent(handler func(MouseEvent) EventHandleState) {
	b.mouseEventHandler = handler
}

type Component interface {
	Renderable
	MouseHandler
	Boundable
	Relational
}

type BaseComponent struct {
	BaseBounds
	BaseMouseHandler
	BaseRelation
	BaseRender
}

func (b *BaseComponent) Render() {
	for _, child := range b.Children() {
		child.Render()
	}
}

func (b *BaseComponent) SetParent(parent Component) {
	b.BaseRelation.SetParent(parent)
}

func (b *BaseComponent) AddChild(child Component) {
	b.BaseRelation.AddChild(child)
	child.SetParent(b)
}

type Renderable interface {
	Render()
	SetColor(color.RGBA)
	GetColor() color.RGBA
	SetOpacity(float32)
	GetOpacity() uint8
}

type BaseRender struct {
	Color   color.RGBA
	Opacity uint8
}

func (b *BaseRender) SetColor(color color.RGBA) {
	b.Color = color
}

func (b *BaseRender) SetOpacity(opacity float32) {
	opacity = min(opacity, 1)
	b.Opacity = MapRangeFloat32ToUint8(opacity, 0, 1, 0, 255)
}

func (b *BaseRender) GetColor() color.RGBA {
	color := b.Color
	color.A = uint8(b.Opacity)
	return color
}

func (b *BaseRender) GetOpacity() uint8 {
	return b.Opacity
}

type Boundable interface {
	BoundingRect() Bounds
}

type BaseBounds struct {
	Bounds
}

func (b *BaseBounds) BoundingRect() Bounds {
	return b.Bounds
}

func MapRangeFloat32ToUint8(value, fromLow, fromHigh float32, toLow, toHigh uint8) uint8 {
	ratio := (value - fromLow) / (fromHigh - fromLow)
	result := float32(toLow) + ratio*(float32(toHigh)-float32(toLow))

	if result < 0 {
		result = 0
	} else if result > 255 {
		result = 255
	}

	return uint8(result)
}
