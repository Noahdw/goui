package main

import (
	"github.com/noahdw/goui/core"
	. "github.com/noahdw/goui/ui"
)

func main() {

	root :=
		Rect(
			Rect().Background("green").Padding(50).Margin(20).Width(100).Height("200").AlignItems("center"),
			Rect().Background("blue").Padding(50).Margin(20).Width(100).Height("200").AlignItems("center"),
			Rect().Background("red").Padding(50).Margin(20).Width(100).Height("200").AlignItems("center"),
		).Background("black").Padding(50).Margin(20).Width(1000).Height(1000).AlignItems("center").FlexDirection("row")

	// Create and run the application
	app := core.NewApplication("Bar Graph Visualization", 800, 600)
	app.SetRoot(root)
	app.Run()
}
