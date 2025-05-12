# Goui

A UI framework for Go that lets you build graphical interfaces using a declarative approach. Built on raylib-go.

## What it does

Goui is a declarative UI framework for Go that provides:

- **Component System**

  - Built-in components (Text, Button, Rect, Image, Layout)
  - Component composition and nesting
  - Declarative component creation

- **Layout & Styling**

  - Flexbox-based layout system
  - Responsive layouts with percentage-based sizing
  - Comprehensive styling (colors, padding, margins, borders, shadows)
  - Text rendering with font styling
  - Image support

- **Interactive Features**
  - Event handling (mouse, keyboard, focus)
  - State management (hover, active, focus, disabled)
  - Component state transitions

## Status

The framework is currently in active development. The core features are implemented and working:

- Core rendering engine and application lifecycle
- All basic components and layout system
- Complete styling system
- Event system and state management
- Component composition and responsive layouts

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

- More advanced components (Input, Select, Checkbox, etc.)
- More layout options (Grid, Table)
- Animations
- Documentation
- WYSIWYG editor (end game)

## License

[License information to be added]
