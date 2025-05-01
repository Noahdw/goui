package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/component.go"
	"github.com/noahdw/goui/core"
	. "github.com/noahdw/goui/node"
)

func main() {
	rl.InitWindow(1000, 700, "Vode")
	rl.SetTargetFPS(120)
	defer rl.CloseWindow()

	root := &BaseNode{}
	renderThread := core.NewRenderThread(root)

	panel := &BaseNode{}
	panel.SetSize(10000, 10000)
	panel.OnMouseEvent(func(event MouseEvent) EventHandleState {
		if event.IsMouseButtonDown() {
			// Move around canvas by dragging mouse
			camera := renderThread.GetCamera()
			delta := rl.GetMouseDelta()
			delta = rl.Vector2Scale(delta, -1.0/camera.Zoom)
			camera.Target = rl.Vector2Add(camera.Target, delta)
			return Handled
		}
		return Propogate
	})
	root.AddChild(panel)

	hBox := component.NewHorizontalBoxLayout()
	hBox.DrawBounds = true
	hBox.SetSpacing(5)

	panel.AddChild(&hBox)
	hBox.SetAlignment(component.AlignTop)

	button2 := component.NewButton()
	button2.SetSize(100, 200)
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

	for i := range 10 {
		button := component.NewButton()
		button.SetSize(float64((i+1)*rand.Intn(40)+1), float64((i+1)*20))
		button.SetText("Test text")
		//button.Label.SetFontSize(32)
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
		hBox.AddChild(button)
	}
	//hBox.AddChild(vLayoutTest())

	renderThread.StartRenderLoop()
}

func vLayoutTest() *component.VerticalBoxLayout {
	vBox := component.NewVerticalBoxLayout()
	vBox.SetSpacing(15)
	vBox.SetAlignment(component.AlignLeft)
	vBox.DrawBounds = true
	for i := range 10 {
		button := component.NewButton()
		button.SetSize(float64((i+1)*rand.Intn(40)+1), float64((i+1)*20))
		button.SetText("Test text")
		//button.Label.SetFontSize(64)
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
