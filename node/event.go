package node

type mouseEventState int

const (
	Pressed mouseEventState = iota
	Released
	Down
	Entered
	Exited
)

type EventHandleState int

const (
	Handled EventHandleState = iota
	Propogate
)

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

type MouseEvent struct {
	state mouseEventState
}

func (m *MouseEvent) IsMouseButtonDown() bool {
	return m.state == Down
}

func (m *MouseEvent) IsMouseButtonPressed() bool {
	return m.state == Pressed
}

func (m *MouseEvent) IsMouseButtonReleased() bool {
	return m.state == Released
}

func (m *MouseEvent) IsMouseEntered() bool {
	return m.state == Entered
}

func (m *MouseEvent) IsMouseExited() bool {
	return m.state == Exited
}

func NewMouseEvent(state mouseEventState) MouseEvent {
	return MouseEvent{
		state: state,
	}
}
