package node

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// RaylibRenderContext implements the RenderContext interface using Raylib
type RaylibRenderContext struct {
	clipRect   Rect
	textureMap map[string]rl.Texture2D
}

// NewRaylibRenderContext creates a new render context using Raylib
func NewRaylibRenderContext() *RaylibRenderContext {
	return &RaylibRenderContext{
		clipRect: Rect{
			Position: Point{0, 0},
			Size: Size{
				Width:  float64(rl.GetScreenWidth()),
				Height: float64(rl.GetScreenHeight()),
			},
		},
		textureMap: make(map[string]rl.Texture2D),
	}
}

// Clear clears the screen with a background color
func (r *RaylibRenderContext) Clear() {
	//rl.ClearBackground(rl.RayWhite)
}

// DrawBackground draws a rectangle with the specified color
func (r *RaylibRenderContext) DrawBackground(bounds Rect, styles Styles) {
	color := rl.Color{
		R: styles.Background.R,
		G: styles.Background.G,
		B: styles.Background.B,
		A: styles.Background.A,
	}
	if styles.BorderRadius.IsNonZero() {
		rect := rl.Rectangle{
			X:      float32(bounds.Position.X),
			Y:      float32(bounds.Position.Y),
			Width:  float32(bounds.Size.Width),
			Height: float32(bounds.Size.Height),
		}
		rl.DrawRectangleRounded(rect, float32(styles.BorderRadius.Top), 60, color)
	} else {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(bounds.Size.Width),
			int32(bounds.Size.Height),
			color,
		)
	}
}

// DrawBorders draws borders with the specified style
func (r *RaylibRenderContext) DrawBorders(bounds Rect, border BorderStyle) {
	// Top border
	if border.Width.Top > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(bounds.Size.Width),
			int32(border.Width.Top),
			rl.Color{
				R: border.Color.R,
				G: border.Color.G,
				B: border.Color.B,
				A: border.Color.A,
			},
		)
	}

	// Right border
	if border.Width.Right > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X+bounds.Size.Width-border.Width.Right),
			int32(bounds.Position.Y),
			int32(border.Width.Right),
			int32(bounds.Size.Height),
			rl.Color{
				R: border.Color.R,
				G: border.Color.G,
				B: border.Color.B,
				A: border.Color.A,
			},
		)
	}

	// Bottom border
	if border.Width.Bottom > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y+bounds.Size.Height-border.Width.Bottom),
			int32(bounds.Size.Width),
			int32(border.Width.Bottom),
			rl.Color{
				R: border.Color.R,
				G: border.Color.G,
				B: border.Color.B,
				A: border.Color.A,
			},
		)
	}

	// Left border
	if border.Width.Left > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(border.Width.Left),
			int32(bounds.Size.Height),
			rl.Color{
				R: border.Color.R,
				G: border.Color.G,
				B: border.Color.B,
				A: border.Color.A,
			},
		)
	}
}

// DrawText draws text with the specified styles
func (r *RaylibRenderContext) DrawText(text string, bounds Rect, styles Styles) {
	fontSize := int32(styles.FontSize.Value)

	// Calculate text dimensions for alignment
	textWidth := rl.MeasureText(text, fontSize)
	textHeight := fontSize

	// Calculate position based on alignment
	var x, y float64

	// Horizontal alignment
	switch styles.TextAlign {
	case "center":
		x = bounds.Position.X + (bounds.Size.Width-float64(textWidth))/2
	case "right":
		x = bounds.Position.X + bounds.Size.Width - float64(textWidth)
	default: // "left" or any other value
		x = bounds.Position.X
	}

	// Vertical centering (simple approach)
	y = bounds.Position.Y + (bounds.Size.Height-float64(textHeight))/2

	// Draw the text
	rl.DrawText(
		text,
		int32(x),
		int32(y),
		fontSize,
		rl.Color{
			R: styles.Color.R,
			G: styles.Color.G,
			B: styles.Color.B,
			A: styles.Color.A,
		},
	)
}

func (r *RaylibRenderContext) DrawTexture(sourceURL string, bounds Rect) {
	texture, has := r.textureMap[sourceURL]
	if !has {
		texture = rl.LoadTexture(sourceURL)
		r.textureMap[sourceURL] = texture
	}

	rl.DrawTexture(texture, int32(bounds.Position.X), int32(bounds.Position.Y), rl.White)
}

// ClipRect returns the current clipping rectangle
func (r *RaylibRenderContext) ClipRect() Rect {
	return r.clipRect
}

// SetClipRect sets the current clipping rectangle
func (r *RaylibRenderContext) SetClipRect(rect Rect) {
	r.clipRect = rect
	// Note: Raylib doesn't directly support clipping rectangles
	// You would need to implement scissoring using OpenGL if needed
}

// Present does nothing in Raylib as it handles frame display automatically
func (r *RaylibRenderContext) Present() {
	// In Raylib, the rendering is presented when EndDrawing() is called
	// This would typically be handled outside this context in the main loop
}
