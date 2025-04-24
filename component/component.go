package component

import . "github.com/noahdw/goui/bounds"

type Relational interface {
	AddChild(Component)
	removeChild(Component)
	Children() []Component
}

type BaseRelation struct {
	children []Component
}

func (b *BaseRelation) AddChild(child Component) {
	b.children = append(b.children, child)
}

func (b *BaseRelation) removeChild(child Component) {

}

func (b *BaseRelation) Children() []Component {
	return b.children
}

type MouseHandler interface {
	HandleMouse(MouseEvent)
}

type BaseMouseHandler struct {
	mouseEventHandler func(MouseEvent)
}

func (b *BaseMouseHandler) HandleMouse(event MouseEvent) {
	if b.mouseEventHandler != nil {
		b.mouseEventHandler(event)
	}
}

func (b *BaseMouseHandler) OnMouseEvent(handler func(MouseEvent)) {
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
}

func (b *BaseComponent) Render() {
	for _, child := range b.Children() {
		child.Render()
	}
}

type Renderable interface {
	Render()
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
