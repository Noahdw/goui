package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/node"
)

// Application represents a UI application window
type Application struct {
	title  string
	width  int
	height int
	root   node.Node
	engine *RenderEngine
}

// NewApplication creates a new application
func NewApplication(title string, width, height int) *Application {
	return &Application{
		title:  title,
		width:  width,
		height: height,
	}
}

// SetRoot sets the root node of the application
func (app *Application) SetRoot(root node.Node) {
	app.root = root
}

// Run starts the application main loop
func (app *Application) Run() {

	initRaylib(app.title, app.width, app.height)
	defer closeRaylib()

	// Create the render context
	context := node.NewRaylibRenderContext()

	// Create the render engine
	app.engine = NewRenderEngine(app.root, context, float64(app.width), float64(app.height))

	// Run the main loop
	for !shouldWindowClose() {
		// Handle input events and dispatch to UI
		processInputEvents(app.root)

		// Begin drawing
		beginDrawing(app.engine.GetCamera())

		// Render the UI
		app.engine.RenderFrame()

		// End drawing
		endDrawing()
	}
}

func initRaylib(title string, width, height int) {
	rl.InitWindow(int32(width), int32(height), title)
	rl.SetTargetFPS(120)
}

func closeRaylib() {
	// Close Raylib
}

func beginDrawing(camera rl.Camera2D) {
	rl.BeginDrawing()
	rl.BeginMode2D(camera)
	rl.ClearBackground(rl.RayWhite)
}

func endDrawing() {
	rl.EndMode2D()
	rl.EndDrawing()
}

func shouldWindowClose() bool {
	return false
}

func processInputEvents(root node.Node) {
	// Process input events and dispatch to UI components
}
