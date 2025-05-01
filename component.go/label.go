package component

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/bounds"
	"github.com/noahdw/goui/node"
)

type Label struct {
	node.BaseNode
	text     string
	fontSize int
}

func NewLabel() *Label {
	return &Label{
		fontSize: 12,
		BaseNode: node.BaseNode{
			Color:   rl.Black,
			Opacity: 1,
		},
	}
}

func (l *Label) SetText(text string) {
	l.text = text
}
func (l *Label) SetFontSize(fontSize int) {
	l.fontSize = fontSize
}

func (l *Label) Render() {
	x, y := l.GlobalPosition()
	rl.DrawText(l.text, int32(x), int32(y), int32(l.fontSize), l.Color)

}

func (l *Label) BoundingRect() bounds.Bounds {
	textWidth := rl.MeasureText(l.text, int32(l.fontSize))
	bounds := bounds.Bounds{
		Width:  float64(textWidth),
		Height: float64(l.fontSize),
	}
	return bounds
}
