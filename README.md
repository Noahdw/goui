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
root := Rect(
    Rect().Background("green").Padding(50).Margin(20).Width(100).Height("33%"),
    Rect().Background("blue").Padding(50).Margin(20).Width(100).Height("200"),
    Rect().Background("red").Padding(50).Margin(20).Width("20%").Height("200"),
).Background("black").Padding(50).Margin(20).Width(1000).Height(1000).FlexDirection("row")

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
