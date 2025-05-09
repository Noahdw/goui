package main

import (
	"github.com/noahdw/goui/core"
	n "github.com/noahdw/goui/node"
	. "github.com/noahdw/goui/ui"
)

func main() {
	// Define scale values
	scaleUp := 1.05
	scaleDown := 0.98
	opacityFull := 1.0

	// Create UI components
	root := Layout("column",
		// Header section
		Rect(
			Text("Interactive UI Demo").FontSize(40).Color("white").FontWeight(900),
		).Background("#2C3E50").Padding(24).Flex("row").AlignItems("center"),

		// Main content area
		Layout("row",
			// Left sidebar
			Rect(
				Text("Navigation").FontSize(24).Color("white").Margin(8),
				Button(
					Text("Home"),
					OnEvent("click", func(e n.UIEvent) { println("Home clicked!") }),
				).Margin(8),
				Button(
					Text("Profile"),
					OnEvent("click", func(e n.UIEvent) { println("Profile clicked!") }),
				).Margin(8),
				Button(
					Text("Settings"),
					OnEvent("click", func(e n.UIEvent) { println("Settings clicked!") }),
				).Margin(8),
			).Background("#34495E").Padding(16).Flex("column").Width(200),

			// Main content
			Layout("column",
				// Interactive card with state-based styling
				Rect(
					Image("figure.png").Width(300).Height(200),
					Text("Hover and click me!").Color("white").FontSize(18),
					StyleOnEvent("hover", &n.StyleProps{
						Background: &n.Color{52, 152, 219, 255}, // Darker blue
						Opacity:    &opacityFull,
					}),
					StyleOnEvent("active", &n.StyleProps{
						Background: &n.Color{41, 128, 185, 255}, // Even darker blue
						Scale:      &scaleDown,
					}),
				).Background("#3498DB").Padding(24).Margin(16).BorderRadius(8).Opacity(0.9),

				// Widget showcase with state-based styling
				widget1(),

				// Interactive buttons row with state-based styling
				Layout("row",
					Button(
						Text("Primary"),
						OnEvent("click", func(e n.UIEvent) { println("Primary action!") }),
						StyleOnEvent("hover", &n.StyleProps{
							Background: &n.Color{46, 204, 113, 255}, // Darker green
							Scale:      &scaleUp,
						}),
					).Background("#2ECC71").Margin(8),
					Button(
						Text("Secondary"),
						OnEvent("click", func(e n.UIEvent) { println("Secondary action!") }),
						StyleOnEvent("hover", &n.StyleProps{
							Background: &n.Color{231, 76, 60, 255}, // Darker red
							Scale:      &scaleUp,
						}),
					).Background("#E74C3C").Margin(8),
					Button(
						Text("Tertiary"),
						OnEvent("click", func(e n.UIEvent) { println("Tertiary action!") }),
						StyleOnEvent("hover", &n.StyleProps{
							Background: &n.Color{241, 196, 15, 255}, // Darker yellow
							Scale:      &scaleUp,
						}),
					).Background("#F1C40F").Margin(8),
				).Padding(16).AlignItems("center"),
			).Flex("column").Padding(24),
		).Background("#ECF0F1"),
	)

	// Create and run the application
	app := core.NewApplication("Interactive UI Demo", 1200, 900)
	app.SetRoot(root)
	app.Run()
}

func widget1() n.Node {
	opacity := 0.9
	return Rect(
		H1(
			Text("My Application"),
			StyleOnEvent("hover", &n.StyleProps{
				Background: &n.Color{240, 240, 240, 255}, // Light gray
				Color:      &n.Color{0, 0, 255, 255},     // Blue
			}),
		).Color("blue").FontSize(30).Background("white").Margin(40).Padding(20),
		H2(
			Text("My Application 2"),
			StyleOnEvent("hover", &n.StyleProps{
				Background: &n.Color{220, 0, 0, 255}, // Darker red
				Opacity:    &opacity,
			}),
		).Color("black").FontSize(20).Background("red").Opacity(.7).Padding(20),
	).Background("gray").Padding(20).BorderRadius(.3).Border("solid").Flex("column")
}
