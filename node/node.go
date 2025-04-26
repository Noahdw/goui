package node

import (
	"image/color"

	"github.com/google/uuid"
	. "github.com/noahdw/goui/bounds"
)

type Relational interface {
	AddChild(Node)
	removeChild(Node)
	Children() []Node
	Parent() Node
	SetParent(Node)
	ID() string
}

type BaseRelation struct {
	children []Node
	parent   Node
	id       string
}

func (b *BaseRelation) AddChild(child Node) {
	b.children = append(b.children, child)
}

func (b *BaseRelation) removeChild(child Node) {

}

func (b *BaseRelation) Children() []Node {
	return b.children
}

func (b *BaseRelation) Parent() Node {
	return b.parent
}

func (b *BaseRelation) SetParent(parent Node) {
	b.parent = parent
}

type Node interface {
	Renderable
	MouseHandler
	Boundable
	Relational
}

type BaseNode struct {
	BaseBounds
	BaseMouseHandler
	BaseRelation
	BaseRender
}

func (b *BaseNode) Render() {
	for _, child := range b.Children() {
		child.Render()
	}
}

func (b *BaseNode) SetParent(parent Node) {
	b.BaseRelation.SetParent(parent)
}

func (b *BaseNode) AddChild(child Node) {
	b.BaseRelation.AddChild(child)
	child.SetParent(b)
}

func (b *BaseNode) ID() string {
	if b.id == "" {
		b.id = uuid.New().String()
	}
	return b.id
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
	dirty   bool
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
	SetPositionY(float64)
	SetPositionX(x float64)
	CheckAndClearDirtyPosition() bool
}

type BaseBounds struct {
	Bounds
	dirtyPosition bool
}

func (b *BaseBounds) BoundingRect() Bounds {
	return b.Bounds
}

func (b *BaseBounds) SetPositionY(y float64) {
	if b.Y == y {
		return
	}
	b.Y = y
	b.dirtyPosition = true
}

func (b *BaseBounds) SetPositionX(x float64) {
	if b.X == x {
		return
	}
	b.X = x
	b.dirtyPosition = true
}

func (b *BaseBounds) CheckAndClearDirtyPosition() bool {
	isDirty := b.dirtyPosition
	b.dirtyPosition = false
	return isDirty
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
