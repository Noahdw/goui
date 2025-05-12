package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	. "github.com/noahdw/goui/node"
)

// EventManager handles all UI events including mouse, keyboard, and focus events
type EventManager struct {
	focusedNode  Node
	pressedObj   Node
	pressX       float64
	pressY       float64
	lastMouseX   float64
	lastMouseY   float64
	lastFoundObj Node
	renderEngine *RenderEngine
}

// NewEventManager creates a new event manager
func NewEventManager(renderEngine *RenderEngine) *EventManager {
	return &EventManager{
		renderEngine: renderEngine,
	}
}

// SetFocus sets the currently focused node
func (e *EventManager) SetFocus(node Node) {
	if e.focusedNode != nil {
		event := NewUIEvent(UIBlur, e.focusedNode)
		e.focusedNode.DispatchEvent(event)
	}

	e.focusedNode = node

	if node != nil {
		event := NewUIEvent(UIFocus, node)
		node.DispatchEvent(event)
	}
}

// HandleMouseEvents processes all mouse-related events
func (e *EventManager) HandleMouseEvents(mouseX, mouseY float64, foundObj Node) {
	mouseMoved := mouseX != e.lastMouseX || mouseY != e.lastMouseY
	e.lastMouseX = mouseX
	e.lastMouseY = mouseY

	if foundObj != nil {
		e.handleMouseButtonEvents(foundObj, mouseX, mouseY)
		e.handleMouseMoveEvents(foundObj, mouseX, mouseY, mouseMoved)
		e.handleMouseEnterLeaveEvents(foundObj, mouseX, mouseY)
	} else if e.lastFoundObj != nil {
		e.handleMouseLeaveEvent(e.lastFoundObj, mouseX, mouseY)
	}
}

// HandleKeyboardEvents processes all keyboard-related events
func (e *EventManager) HandleKeyboardEvents() {
	key := rl.GetKeyPressed()
	if key != 0 && e.focusedNode != nil {
		event := NewUIKeyboardEvent(UIKeyPress, e.focusedNode, int(key), rune(key))
		e.focusedNode.DispatchEvent(event)
	}

	if rl.IsKeyPressed(rl.KeyTab) && e.focusedNode != nil {
		// TODO: Implement focus navigation logic
		e.SetFocus(nil)
	}
}

func (e *EventManager) handleMouseButtonEvents(foundObj Node, mouseX, mouseY float64) {
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		e.pressedObj = foundObj
		e.pressX = mouseX
		e.pressY = mouseY

		event := NewUIMouseEvent(UIPress, foundObj, e.pressX, e.pressY)
		foundObj.DispatchEvent(event)

		e.SetFocus(foundObj)
		e.renderEngine.MarkLayoutDirty()
	}

	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		event := NewUIMouseEvent(UIRelease, foundObj, mouseX, mouseY)
		foundObj.DispatchEvent(event)

		if e.pressedObj == foundObj {
			dx := mouseX - e.pressX
			dy := mouseY - e.pressY
			distance := dx*dx + dy*dy

			if distance < 100 {
				clickEvent := NewUIMouseEvent(UIClick, foundObj, mouseX, mouseY)
				foundObj.DispatchEvent(clickEvent)
			}
		}

		e.pressedObj = nil
		e.renderEngine.MarkLayoutDirty()
	}
}

func (e *EventManager) handleMouseMoveEvents(foundObj Node, mouseX, mouseY float64, mouseMoved bool) {
	if mouseMoved {
		event := NewUIMouseEvent(UIMove, foundObj, mouseX, mouseY)
		foundObj.DispatchEvent(event)
	}
}

func (e *EventManager) handleMouseEnterLeaveEvents(foundObj Node, mouseX, mouseY float64) {
	if foundObj != e.lastFoundObj {
		event := NewUIMouseEvent(UIEnter, foundObj, mouseX, mouseY)
		foundObj.DispatchEvent(event)

		if e.lastFoundObj != nil {
			e.handleMouseLeaveEvent(e.lastFoundObj, mouseX, mouseY)
		}
		e.lastFoundObj = foundObj
		e.renderEngine.MarkLayoutDirty()
	}
}

func (e *EventManager) handleMouseLeaveEvent(node Node, mouseX, mouseY float64) {
	event := NewUIMouseEvent(UILeave, node, mouseX, mouseY)
	node.DispatchEvent(event)
	e.lastFoundObj = nil
	e.renderEngine.MarkLayoutDirty()
}
