package main

import (
	"github.com/noahdw/goui/core"
	. "github.com/noahdw/goui/node"
)

func main() {
	// Create UI components
	root :=
		Layout("column",
			H1(Text("My Application")).
				Color("blue").
				Padding(30).
				FontSize(30),
			H2(Text("My Application 2")).
				Color("black").
				Padding(20).
				FontSize(40).Background("red"),
		).Background("green")

	// Create and run the application
	app := core.NewApplication("My App", 800, 600)
	app.SetRoot(root)
	app.Run()
}
