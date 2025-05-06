package main

import (
	"github.com/noahdw/goui/core"
	n "github.com/noahdw/goui/node"
	. "github.com/noahdw/goui/ui"
)

func main() {
	// Create UI components
	root :=
		Rect(
			widget1(),
			Rect(
				Image("figure.png"),
			).Background("blue").Padding(5).BorderRadius(.3).Opacity(.5),
		).Border("solid").BorderWidth(3).Background("red").Padding(32).Margin(4).Opacity(.5)

	// Create and run the application
	app := core.NewApplication("My App", 1200, 900)
	app.SetRoot(root)
	app.Run()
}

func widget1() n.Node {
	return Rect(
		H1(Text("My Application")).
			Color("blue").
			FontSize(30).Background("white").Margin(40),
		H2(Text("My Application 2")).
			Color("black").
			FontSize(20).Background("red"),
	).Background("gray").Padding(20).BorderRadius(.3).Border("solid").Flex("column")
}
