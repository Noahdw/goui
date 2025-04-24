package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/bounds"
	. "github.com/noahdw/goui/component"
	"github.com/noahdw/goui/core"
)

type Button struct {
	BaseComponent
}

func (b *Button) Render() {
	rl.DrawRectangle(int32(b.X), int32(b.Y), int32(b.Width), int32(b.Height), rl.Red)

}

func NewButton(bounds bounds.Bounds) *Button {
	return &Button{
		BaseComponent: BaseComponent{
			BaseBounds: BaseBounds{
				Bounds: bounds,
			},
		},
	}
}

func main() {
	rl.InitWindow(1000, 700, "Vode")
	rl.SetTargetFPS(120)
	defer rl.CloseWindow()

	root := &BaseComponent{}
	ui := core.NewUI(root)

	panel := &BaseComponent{
		BaseBounds: BaseBounds{
			Bounds: Bounds{
				X:      0,
				Y:      0,
				Width:  10000,
				Height: 10000,
			},
		},
		BaseMouseHandler: BaseMouseHandler{},
		BaseRelation:     BaseRelation{},
	}
	panel.OnMouseEvent(func(event MouseEvent) {
		if event.IsMouseButtonDown() {
			camera := ui.GetCamera()
			delta := rl.GetMouseDelta()
			delta = rl.Vector2Scale(delta, -1.0/camera.Zoom)
			camera.Target = rl.Vector2Add(camera.Target, delta)
		}
	})
	root.AddChild(panel)

	button := NewButton(Bounds{
		X:      100,
		Y:      100,
		Width:  200,
		Height: 200,
	})
	button.OnMouseEvent(func(event MouseEvent) {
		if event.IsMouseButtonPressed() {
			println("Wow it works")
		}
	})

	button2 := NewButton(Bounds{
		X:      400,
		Y:      400,
		Width:  100,
		Height: 200,
	})
	button2.OnMouseEvent(func(event MouseEvent) {
		if event.IsMouseButtonPressed() {
			println("Wow it works2")
		}
	})

	panel.AddChild(button)
	panel.AddChild(button2)

	ui.RenderLoop()
}
