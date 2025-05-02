package node

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	. "github.com/noahdw/goui/bounds"
)

func (b *BaseNode) removeChild(child Node) {

}

func (b *BaseNode) Children() []Node {
	return b.children
}

func (b *BaseNode) Parent() Node {
	return b.parent
}

func (b *BaseNode) SetParent(parent Node) {
	b.parent = parent
}

type Node interface {
	MouseHandler
	AddChild(Node)
	removeChild(Node)
	Children() []Node
	Parent() Node
	SetParent(Node)
	ID() string
	Render()
	SetColor(color.RGBA)
	GetColor() color.RGBA
	SetOpacity(float32)
	GetOpacity() uint8
	BoundingRect() Bounds
	SetGlobalPositionY(y float64)
	SetGlobalPositionX(x float64)
	SetPositionY(y float64)
	SetPositionX(x float64)
	CheckAndClearDirtyPosition() bool
	GlobalPosition() (float64, float64)
	SetGlobalPosition(x, y float64)
	PositionX() float64
	PositionY() float64
	GlobalPositionX() float64
	GlobalPositionY() float64
	Width() float64
	Height() float64
	SetSize(width, height float64)
	markDirty()
	UpdateLayout()
}

type BaseNode struct {
	BaseMouseHandler
	Color         color.RGBA
	Opacity       uint8
	DesiredBounds Bounds
	DrawBounds    bool

	bounds        Bounds
	dirtyPosition bool
	localXOffset  float64
	localYOffset  float64
	children      []Node
	parent        Node
	id            string
	dirty         bool
}

func (b *BaseNode) Render() {
	for _, child := range b.Children() {
		child.Render()
	}
	if b.DrawBounds {
		rl.DrawBoundingBox(rl.BoundingBox{
			Min: rl.Vector3{
				X: float32(b.bounds.X),
				Y: float32(b.bounds.Y),
				Z: 0,
			},
			Max: rl.Vector3{
				X: float32(b.bounds.X) + float32(b.bounds.Width),
				Y: float32(b.bounds.Y) + float32(b.bounds.Height),
				Z: 0,
			},
		}, rl.Black)
	}
}

func (b *BaseNode) AddChild(child Node) {
	b.children = append(b.children, child)
	child.SetParent(b)
}

func (b *BaseNode) BoundingRect() Bounds {
	return *b.bounds.BoundingRect()
}

func (b *BaseNode) Width() float64 {
	return b.bounds.Width
}

func (b *BaseNode) Height() float64 {
	return b.bounds.Height
}

func (b *BaseNode) SetSize(width, height float64) {
	b.bounds.Width = width
	b.bounds.Height = height
}

func (b *BaseNode) Intersects(otherNode Node) bool {
	return b.bounds.Intersects(otherNode.BoundingRect())
}

func (b *BaseNode) SetGlobalPosition(x, y float64) {
	b.SetGlobalPositionX(x)
	b.SetGlobalPositionY(y)
}

func (b *BaseNode) SetGlobalPositionX(x float64) {
	if b.bounds.X == x {
		return
	}
	b.bounds.X = x // should actually adjust local pos
	b.markDirty()
}

func (b *BaseNode) SetGlobalPositionY(y float64) {
	if b.bounds.Y == y {
		return
	}
	b.bounds.Y = y // should actually adjust local pos
	b.markDirty()
}

func (b *BaseNode) SetPositionX(x float64) {
	if b.localXOffset == x {
		return
	}
	b.localXOffset = x
	b.markDirty()
}

func (b *BaseNode) SetPositionY(y float64) {
	if b.localYOffset == y {
		return
	}
	b.localYOffset = y
	b.markDirty()
}

func (b *BaseNode) SetWidth(width float64) {
	if b.bounds.Width == width {
		return
	}
	b.bounds.Width = width
	b.markDirty()
}

func (b *BaseNode) Position() (float64, float64) {
	return b.localXOffset, b.localYOffset
}

func (b *BaseNode) PositionX() float64 {
	return b.localXOffset
}

func (b *BaseNode) PositionY() float64 {
	return b.localYOffset
}

func (b *BaseNode) GlobalPositionX() float64 {
	if b.parent != nil {
		return b.Parent().GlobalPositionX() + b.localXOffset
	}
	return b.localXOffset
}

func (b *BaseNode) GlobalPositionY() float64 {
	if b.parent != nil {
		return b.Parent().GlobalPositionY() + b.localYOffset
	}
	return b.localYOffset
}

func (b *BaseNode) GlobalPosition() (float64, float64) {
	if b.parent != nil {
		if b.dirtyPosition {
			x, y := b.parent.GlobalPosition()
			b.bounds.X = x
			b.bounds.Y = y
			return x + b.localXOffset, y + b.localYOffset
		} else {
			return b.bounds.X, b.bounds.Y
		}

	}
	return b.Position()
}

func (b *BaseNode) SetHeight(height float64) {
	if b.bounds.Height == height {
		return
	}
	b.bounds.Height = height
	b.markDirty()
}

func (b *BaseNode) CheckAndClearDirtyPosition() bool { // NEED (internal)
	isDirty := b.dirtyPosition
	if !isDirty {
		return false
	}
	b.bounds.X = b.GlobalPositionX()
	b.bounds.Y = b.GlobalPositionY()
	b.dirtyPosition = false

	return true
}

func (b *BaseNode) MaxChildHeight() float64 { // CONV
	maxHeight := 0.0
	for _, child := range b.Children() {
		maxHeight = max(maxHeight, child.BoundingRect().Height)
	}
	return maxHeight
}

// Bubbles the dirty marker up the tree
func (b *BaseNode) markDirty() {
	if b.dirtyPosition == true {
		return
	}
	b.dirtyPosition = true
	if b.parent != nil {
		b.parent.markDirty()
	}
}

func (b *BaseNode) MaxChildWidth() float64 { // CONV
	maxWidth := 0.0
	for _, child := range b.Children() {
		maxWidth = max(maxWidth, child.BoundingRect().Width)
	}
	return maxWidth
}

func (b *BaseNode) ID() string {
	if b.id == "" {
		b.id = uuid.New().String()
	}
	return b.id
}

func (b *BaseNode) SetColor(color color.RGBA) {
	b.Color = color
}

func (b *BaseNode) SetOpacity(opacity float32) {
	opacity = min(opacity, 1)
	b.Opacity = MapRangeFloat32ToUint8(opacity, 0, 1, 0, 255)
}

func (b *BaseNode) GetColor() color.RGBA {
	color := b.Color
	color.A = uint8(b.Opacity)
	return color
}

func (b *BaseNode) GetOpacity() uint8 {
	return b.Opacity
}

func (b *BaseNode) UpdateLayout() {

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
