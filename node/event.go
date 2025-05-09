package node

import (
	"time"
)

// UIEventType represents the type of UI event
type UIEventType string

const (
	// Mouse events
	UIClick   UIEventType = "click"
	UIPress   UIEventType = "mousedown"
	UIRelease UIEventType = "mouseup"
	UIEnter   UIEventType = "mouseenter"
	UILeave   UIEventType = "mouseleave"
	UIMove    UIEventType = "mousemove"

	// Keyboard events
	UIKeyPress   UIEventType = "keydown"
	UIKeyRelease UIEventType = "keyup"
	UIKeyChar    UIEventType = "keychar"

	// Focus events
	UIFocus UIEventType = "focus"
	UIBlur  UIEventType = "blur"
)

// UIEvent represents a UI event
type UIEvent struct {
	Type      UIEventType
	Target    Node
	Timestamp int64
	// Event specific data
	MouseX  float64
	MouseY  float64
	KeyCode int
	KeyChar rune
	// Prevent default behavior
	PreventDefault bool
	// Stop propagation
	StopPropagation bool
}

// NewUIEvent creates a new UI event with the given type
func NewUIEvent(eventType UIEventType, target Node) UIEvent {
	return UIEvent{
		Type:      eventType,
		Target:    target,
		Timestamp: time.Now().UnixNano(),
	}
}

// NewUIMouseEvent creates a new mouse UI event
func NewUIMouseEvent(eventType UIEventType, target Node, x, y float64) UIEvent {
	return UIEvent{
		Type:      eventType,
		Target:    target,
		Timestamp: time.Now().UnixNano(),
		MouseX:    x,
		MouseY:    y,
	}
}

// NewUIKeyboardEvent creates a new keyboard UI event
func NewUIKeyboardEvent(eventType UIEventType, target Node, keyCode int, keyChar rune) UIEvent {
	return UIEvent{
		Type:      eventType,
		Target:    target,
		Timestamp: time.Now().UnixNano(),
		KeyCode:   keyCode,
		KeyChar:   keyChar,
	}
}
