package component

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/node"
)

type Button struct {
	node.BaseNode
	label *Label
}

func (b *Button) Render() {
	x, y := b.GlobalPosition()
	rl.DrawRectangle(int32(x), int32(y), int32(b.Width()), int32(b.Height()), b.GetColor())
	b.label.SetPositionX(b.Width()/2 - b.label.Width()/2)
	b.label.SetPositionY(b.Height()/2 - b.label.Height()/2)
	b.label.Render()
}

func NewButton() *Button {
	button := &Button{
		BaseNode: node.BaseNode{
			Color:   rl.Black,
			Opacity: 255,
		},
	}
	label := NewLabel()
	button.label = label
	button.AddChild(label)
	return button
}

func (b *Button) SetText(text string) {
	b.label.text = text
}
