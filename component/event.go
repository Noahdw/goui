package component

type mouseEventState int

const (
	Pressed mouseEventState = iota
	Released
	Down
)

type EventHandleState int

const (
	Handled EventHandleState = iota
	Propogate
)

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

func NewMouseEvent(state mouseEventState) MouseEvent {
	return MouseEvent{
		state: state,
	}
}
