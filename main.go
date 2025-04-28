package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/bounds"
	"github.com/noahdw/goui/component.go"
	"github.com/noahdw/goui/core"
	. "github.com/noahdw/goui/node"
)

type Button struct {
	BaseNode
}

func (b *Button) Render() {
	rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), b.GetColor())
	b.BaseNode.Render()
}

func NewButton(bounds bounds.Bounds) *Button {
	return &Button{
		BaseNode: BaseNode{
			Bounds: bounds,
			BaseRender: BaseRender{
				Color:   rl.Black,
				Opacity: 255,
			},
		},
	}
}

func main() {
	rl.InitWindow(1000, 700, "Vode")
	rl.SetTargetFPS(120)
	defer rl.CloseWindow()

	root := &BaseNode{}
	ui := core.NewUI(root)

	panel := &BaseNode{
		Bounds: Bounds{
			X:      0,
			Y:      0,
			Width:  10000,
			Height: 10000,
		},
	}
	panel.OnMouseEvent(func(event MouseEvent) EventHandleState {
		if event.IsMouseButtonDown() {
			// Move around canvas by dragging mouse
			camera := ui.GetCamera()
			delta := rl.GetMouseDelta()
			delta = rl.Vector2Scale(delta, -1.0/camera.Zoom)
			camera.Target = rl.Vector2Add(camera.Target, delta)
			return Handled
		}
		return Propogate
	})
	root.AddChild(panel)

	hBox := component.NewHorizontalBoxLayout()
	hBox.SetSpacing(5)
	hBox.DrawBounds = true
	panel.AddChild(&hBox)
	hBox.SetAlignment(component.AlignTop)

	button2 := NewButton(Bounds{
		Width:  100,
		Height: 200,
	})
	button2.OnMouseEvent(func(event MouseEvent) EventHandleState {
		if event.IsMouseButtonPressed() {
			println("Wow it works2")
			return Propogate
		}
		return Propogate
	})
	button2.SetColor(rl.Blue)
	button2.SetOpacity(0.2)

	hBox.AddChild(button2)

	for range 10 {
		button := vLayoutTest()
		hBox.AddChild(button)
	}
	hBox.AddChild(vLayoutTest())

	ui.RenderLoop()
}

func vLayoutTest() *component.VerticalBoxLayout {
	vBox := component.NewVerticalBoxLayout()
	vBox.SetSpacing(15)
	vBox.SetAlignment(component.AlignLeft)

	for i := range 10 {
		button := NewButton(Bounds{
			Width:  float64((i + 1) * 20),
			Height: float64((i + 1) * 20),
		})
		button.SetColor(rl.Green)
		button.SetOpacity(0.1 * float32((i + 1)))
		button.OnMouseEvent(func(event MouseEvent) EventHandleState {
			if event.IsMouseButtonPressed() {
				println("Wow it works4")
			} else if event.IsMouseButtonDown() {
				button.SetColor(rl.Red)
			} else if event.IsMouseButtonReleased() {
				button.SetColor(rl.DarkGreen)
			} else if event.IsMouseEntered() {
				button.SetColor(rl.DarkGreen)
			} else if event.IsMouseExited() {
				button.SetColor(rl.Green)
			}
			return Handled
		})
		vBox.AddChild(button)
	}
	return &vBox
}
