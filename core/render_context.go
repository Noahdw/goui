package core

import (
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/noahdw/goui/node/style"
)

// RaylibRenderContext implements the RenderContext interface using Raylib
type RaylibRenderContext struct {
	clipRect    style.Rect
	textureMap  map[string]rl.Texture2D
	opacity     float64
	fillColor   style.Color
	strokeColor style.Color
	lineWidth   float64
	fontSize    float64
	transform   struct {
		scaleX float64
		scaleY float64
	}
}

// NewRaylibRenderContext creates a new render context using Raylib
func NewRaylibRenderContext() *RaylibRenderContext {
	return &RaylibRenderContext{
		clipRect: style.Rect{
			Position: style.Point{X: 0, Y: 0},
			Size: style.Size{
				Width:  float64(rl.GetScreenWidth()),
				Height: float64(rl.GetScreenHeight()),
			},
		},
		textureMap:  make(map[string]rl.Texture2D),
		opacity:     1.0,
		fillColor:   style.White,
		strokeColor: style.Black,
		lineWidth:   1.0,
		fontSize:    16.0,
		transform: struct {
			scaleX float64
			scaleY float64
		}{
			scaleX: 1.0,
			scaleY: 1.0,
		},
	}
}

// Clear clears the screen with a background color
func (r *RaylibRenderContext) Clear() {
	rl.ClearBackground(rl.RayWhite)
}

// Save saves the current rendering state
func (r *RaylibRenderContext) Save() {
	// In a real implementation, we would save the current state
	// For now, we just track opacity and transform
	r.opacity = 1.0
	r.transform.scaleX = 1.0
	r.transform.scaleY = 1.0
}

// Restore restores the previously saved rendering state
func (r *RaylibRenderContext) Restore() {
	// In a real implementation, we would restore the saved state
	// For now, we just reset opacity and transform
	r.opacity = 1.0
	r.transform.scaleX = 1.0
	r.transform.scaleY = 1.0
}

// SetOpacity sets the current opacity
func (r *RaylibRenderContext) SetOpacity(opacity float64) {
	r.opacity = opacity
}

// SetFillColor sets the current fill color
func (r *RaylibRenderContext) SetFillColor(color style.Color) {
	r.fillColor = color
}

// SetStrokeColor sets the current stroke color
func (r *RaylibRenderContext) SetStrokeColor(color style.Color) {
	r.strokeColor = color
}

// SetLineWidth sets the current line width
func (r *RaylibRenderContext) SetLineWidth(width float64) {
	r.lineWidth = width
}

// SetFontSize sets the current font size
func (r *RaylibRenderContext) SetFontSize(size float64) {
	r.fontSize = size
}

// Scale sets the current scale transform
func (r *RaylibRenderContext) Scale(x, y float64) {
	r.transform.scaleX *= x
	r.transform.scaleY *= y
}

// StrokeLine draws a line from start to end
func (r *RaylibRenderContext) StrokeLine(start, end style.Point) {
	rl.DrawLineEx(
		rl.Vector2{X: float32(start.X), Y: float32(start.Y)},
		rl.Vector2{X: float32(end.X), Y: float32(end.Y)},
		float32(r.lineWidth),
		rl.Color{
			R: r.strokeColor.R,
			G: r.strokeColor.G,
			B: r.strokeColor.B,
			A: NormalizedFloatToUint8(r.opacity),
		},
	)
}

// FillRect fills a rectangle with the current fill color
func (r *RaylibRenderContext) FillRect(rect style.Rect) {
	rl.DrawRectangle(
		int32(rect.Position.X),
		int32(rect.Position.Y),
		int32(rect.Size.Width),
		int32(rect.Size.Height),
		rl.Color{
			R: r.fillColor.R,
			G: r.fillColor.G,
			B: r.fillColor.B,
			A: NormalizedFloatToUint8(r.opacity),
		},
	)
}

// DrawBackground draws a background with the specified styles
func (r *RaylibRenderContext) DrawBackground(bounds style.Rect, styles style.Styles, opacity float64) {
	// Get background color
	bgColor, ok := styles.GetColor("background")
	if !ok {
		return
	}

	// Get border radius
	borderRadius, ok := styles.GetEdgeInsets("borderRadius")
	if !ok {
		borderRadius = style.EdgeInsets{}
	}

	// Draw background with border radius
	if borderRadius.IsNonZero() {
		// For now, just draw a rectangle since raylib doesn't support border radius
		// TODO: Implement proper border radius drawing
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(bounds.Size.Width),
			int32(bounds.Size.Height),
			rl.Color{
				R: bgColor.R,
				G: bgColor.G,
				B: bgColor.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	} else {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(bounds.Size.Width),
			int32(bounds.Size.Height),
			rl.Color{
				R: bgColor.R,
				G: bgColor.G,
				B: bgColor.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	}
}

// DrawBorders draws borders with the specified style
func (r *RaylibRenderContext) DrawBorders(bounds style.Rect, styles style.Styles, opacity float64) {
	// Get border style
	border, ok := styles.Get("border")
	if !ok {
		return
	}
	borderStyle, ok := border.(style.BorderStyle)
	if !ok {
		return
	}

	// Top border
	if borderStyle.Width.Top > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(bounds.Size.Width),
			int32(borderStyle.Width.Top),
			rl.Color{
				R: borderStyle.Color.R,
				G: borderStyle.Color.G,
				B: borderStyle.Color.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	}

	// Right border
	if borderStyle.Width.Right > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X+bounds.Size.Width-borderStyle.Width.Right),
			int32(bounds.Position.Y),
			int32(borderStyle.Width.Right),
			int32(bounds.Size.Height),
			rl.Color{
				R: borderStyle.Color.R,
				G: borderStyle.Color.G,
				B: borderStyle.Color.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	}

	// Bottom border
	if borderStyle.Width.Bottom > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y+bounds.Size.Height-borderStyle.Width.Bottom),
			int32(bounds.Size.Width),
			int32(borderStyle.Width.Bottom),
			rl.Color{
				R: borderStyle.Color.R,
				G: borderStyle.Color.G,
				B: borderStyle.Color.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	}

	// Left border
	if borderStyle.Width.Left > 0 {
		rl.DrawRectangle(
			int32(bounds.Position.X),
			int32(bounds.Position.Y),
			int32(borderStyle.Width.Left),
			int32(bounds.Size.Height),
			rl.Color{
				R: borderStyle.Color.R,
				G: borderStyle.Color.G,
				B: borderStyle.Color.B,
				A: NormalizedFloatToUint8(opacity * r.opacity),
			},
		)
	}
}

// DrawText draws text with the specified styles
func (r *RaylibRenderContext) DrawText(text string, bounds style.Rect, styles style.Styles, opacity float64) {
	fontSize, _ := styles.GetFloat("fontSize")
	padding, _ := styles.GetEdgeInsets("padding")
	textAlign, _ := styles.GetString("textAlign")
	alignItems, _ := styles.GetString("alignItems")
	textColor, _ := styles.GetColor("color")

	textWidth := rl.MeasureText(text, int32(fontSize))
	textHeight := fontSize * 1.2 // Use line height for better vertical centering

	// Calculate position based on alignment
	var x, y float64

	// Horizontal alignment
	switch textAlign {
	case "center":
		x = bounds.Position.X + (bounds.Size.Width-float64(textWidth))/2
	case "right":
		x = bounds.Position.X + bounds.Size.Width - float64(textWidth) - padding.Right
	default: // "left" or any other value
		x = bounds.Position.X + padding.Left
	}

	// Vertical alignment
	switch alignItems {
	case "center":
		y = bounds.Position.Y + (bounds.Size.Height-textHeight)/2
	case "bottom":
		y = bounds.Position.Y + bounds.Size.Height - textHeight - padding.Bottom
	default: // "top" or any other value
		y = bounds.Position.Y + padding.Top
	}

	// Draw the text
	rl.DrawText(
		text,
		int32(x),
		int32(y),
		int32(fontSize),
		rl.Color{
			R: textColor.R,
			G: textColor.G,
			B: textColor.B,
			A: NormalizedFloatToUint8(opacity * r.opacity),
		},
	)
}

// LoadTexture loads a texture from a URL
func (r *RaylibRenderContext) LoadTexture(sourceURL string) rl.Texture2D {
	texture, has := r.textureMap[sourceURL]
	if !has {
		// Check if the file exists
		if _, err := os.Stat(sourceURL); os.IsNotExist(err) {
			// Return empty texture if file doesn't exist
			return rl.Texture2D{}
		}

		texture = rl.LoadTexture(sourceURL)
		if texture.ID != 0 {
			r.textureMap[sourceURL] = texture
		}
	}
	return texture
}

// UnloadTexture unloads a texture from memory
func (r *RaylibRenderContext) UnloadTexture(sourceURL string) {
	if texture, has := r.textureMap[sourceURL]; has {
		rl.UnloadTexture(texture)
		delete(r.textureMap, sourceURL)
	}
}

// UnloadAllTextures unloads all textures from memory
func (r *RaylibRenderContext) UnloadAllTextures() {
	for _, texture := range r.textureMap {
		rl.UnloadTexture(texture)
	}
	r.textureMap = make(map[string]rl.Texture2D)
}

// ClipRect returns the current clipping rectangle
func (r *RaylibRenderContext) ClipRect() style.Rect {
	return r.clipRect
}

// SetClipRect sets the current clipping rectangle
func (r *RaylibRenderContext) SetClipRect(rect style.Rect) {
	r.clipRect = rect
	// Note: Raylib doesn't directly support clipping rectangles
	// You would need to implement scissoring using OpenGL if needed
}

// Present does nothing in Raylib as it handles frame display automatically
func (r *RaylibRenderContext) Present() {
	// In Raylib, the rendering is presented when EndDrawing() is called
	// This would typically be handled outside this context in the main loop
}

func NormalizedFloatToUint8(value float64) uint8 {
	// Clamp value between 0.0 and 1.0
	if value < 0.0 {
		value = 0.0
	} else if value > 1.0 {
		value = 1.0
	}

	// Scale to 0-255 range and convert to uint8
	return uint8(value * 255.0)
}

// DrawTexture draws a texture with the specified styles
func (r *RaylibRenderContext) DrawTexture(sourceURL string, bounds style.Rect, styles style.Styles, opacity float64) {
	texture := r.LoadTexture(sourceURL)
	if texture.ID == 0 {
		// Skip drawing if texture failed to load
		return
	}

	color := rl.White
	color.A = NormalizedFloatToUint8(opacity * r.opacity)

	// Get object-fit style if specified
	objectFit, _ := styles.GetString("objectFit")
	if objectFit == "" {
		objectFit = "contain" // Default to contain
	}

	texWidth := float32(texture.Width)
	texHeight := float32(texture.Height)
	var destRect rl.Rectangle

	switch objectFit {
	case "cover":
		// Scale to cover the entire bounds while maintaining aspect ratio
		scale := math.Max(float64(bounds.Size.Width)/float64(texWidth), float64(bounds.Size.Height)/float64(texHeight))
		scaledWidth := float32(float64(texWidth) * scale)
		scaledHeight := float32(float64(texHeight) * scale)

		// Center the texture
		x := float32(bounds.Position.X) + (float32(bounds.Size.Width)-scaledWidth)/2
		y := float32(bounds.Position.Y) + (float32(bounds.Size.Height)-scaledHeight)/2

		destRect = rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  scaledWidth,
			Height: scaledHeight,
		}

	case "fill":
		// Stretch to fill the bounds exactly
		destRect = rl.Rectangle{
			X:      float32(bounds.Position.X),
			Y:      float32(bounds.Position.Y),
			Width:  float32(bounds.Size.Width),
			Height: float32(bounds.Size.Height),
		}

	default: // "contain" or any other value
		// Scale to fit within bounds while maintaining aspect ratio
		scale := math.Min(float64(bounds.Size.Width)/float64(texWidth), float64(bounds.Size.Height)/float64(texHeight))
		scaledWidth := float32(float64(texWidth) * scale)
		scaledHeight := float32(float64(texHeight) * scale)

		// Center the texture
		x := float32(bounds.Position.X) + (float32(bounds.Size.Width)-scaledWidth)/2
		y := float32(bounds.Position.Y) + (float32(bounds.Size.Height)-scaledHeight)/2

		destRect = rl.Rectangle{
			X:      x,
			Y:      y,
			Width:  scaledWidth,
			Height: scaledHeight,
		}
	}

	// Get object-position style if specified
	objectPosition, _ := styles.GetString("objectPosition")
	if objectPosition != "" {
		// Parse object-position (e.g., "center", "top left", "50% 50%")
		// For now, we'll just support "center" as the default
		// TODO: Implement full object-position support
	}

	rl.DrawTexturePro(
		texture,
		rl.Rectangle{X: 0, Y: 0, Width: texWidth, Height: texHeight},
		destRect,
		rl.Vector2{X: 0, Y: 0},
		0,
		color,
	)
}
