# Goui

A UI framework for Go that lets you build graphical interfaces using a declarative approach. Built on raylib-go.

## What it does

- Declarative UI components with a simple API
- Layout system with flexbox-like properties
- Basic styling (colors, padding, margins)
- Component-based architecture

## Status

Working features:

- Core rendering engine
- Basic components (images, text, rects)
- Layout system
- Basic styling

## Example

```go
// Create a header component with text and image
header := Rect(
    Text("Welcome to Goui").FontSize(24).Color("white"),
    Image("logo.png").Width(50).Height(50),
).Background("navy").Padding(20).FlexDirection("row").JustifyContent("space-between")

// Create a content section with nested components
content := Rect(
    Rect(
        Text("Left Panel").Color("white"),
        Image("icon1.png").Width(32).Height(32),
    ).Background("green").Padding(20).Width("30%"),
    Rect(
        Text("Main Content").Color("black"),
        Text("A simple example showing text, images, and nested components").FontSize(14),
    ).Background("white").Padding(20).Width("70%"),
).Background("lightgray").Padding(10).FlexDirection("row").Height(300)

// Combine components into the root layout
root := Rect(
    header,
    content,
).Background("black").Padding(20).Width(800).Height(600).FlexDirection("column")

app := core.NewApplication("My App", 800, 600)
app.SetRoot(root)
app.Run()
```

## Structure

- `core/` - Framework core
  - `application.go` - App lifecycle
  - `render_engine.go` - Rendering
  - `render_context.go` - Graphics context
- `ui/` - Components
  - `basic_components.go` - Basic UI elements

## Dependencies

- raylib-go - Graphics and window management

## Planned

- More components (Text, Button, Input)
- Better styling system
- Event handling
- Animations
- More layout options
- Documentation
- WYSIWYG editor (end game)

## License

[License information to be added]
