package main

import (
	"github.com/noahdw/goui/core"
	. "github.com/noahdw/goui/node"
)

func main() {
	// Create UI components
	root :=
		Layout("column",
			Layout("row",
				widget1(),
				widget1(),
			).Background("green").Padding(40).BorderRadius(1).Border("solid").Margin(40),
			Layout("row",
				widget1(),
				widget1(),
			).Background("green").Padding(40).BorderRadius(1).Border("solid").Margin(40),
			Layout("column",
				Image("figure.png"),
			).Border("solid").BorderWidth(4).Margin(20).Background("red").Padding(20),
		).Margin(20).Background("blue").Padding(20)

	// Create and run the application
	app := core.NewApplication("My App", 1200, 900)
	app.SetRoot(root)
	app.Run()
}

func widget1() Node {
	return Layout("column",
		H1(Text("My Application")).
			Color("blue").
			FontSize(30).Background("white").Margin(40),
		H2(Text("My Application 2")).
			Color("black").
			FontSize(20).Background("red"),
	).Background("gray").Padding(20).BorderRadius(.3).Border("solid")
}
