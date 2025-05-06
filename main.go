package main

import (
	"github.com/noahdw/goui/core"
	n "github.com/noahdw/goui/node"
	. "github.com/noahdw/goui/ui"
)

func doSomething(e n.Event) {
	println("Something")
}

func main() {
	// Create UI components
	root :=
		Rect(
			widget1(),
			Rect(
				OnEvent("click", doSomething),
				Image("figure.png"),
			).Background("blue").Padding(5).BorderRadius(.3).Opacity(1).Padding(40),
		).Background("red").Padding(32).Border("solid").BorderWidth(3).Margin(4)

	// Create and run the application
	app := core.NewApplication("My App", 1200, 900)
	app.SetRoot(root)
	app.Run()
}

func widget1() n.Node {
	return Rect(
		H1(
			Text("My Application"),
			OnEvent("click", doSomething),
		).Color("blue").FontSize(30).Background("white").Margin(40),
		H2(Text("My Application 2")).
			Color("black").
			FontSize(20).Background("red").Opacity(.7),
	).Background("gray").Padding(20).BorderRadius(.3).Border("solid").Flex("column")
}
