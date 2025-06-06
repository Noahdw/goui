package node

import (
	"strconv"
	"strings"

	"github.com/noahdw/goui/node/style"
)

// Style builder methods for Node interface
type StyleBuilder interface {
	// Layout properties
	Width(value interface{}) Node  // Can be number, "auto", percentage string, etc.
	Height(value interface{}) Node // Can be number, "auto", percentage string, etc.
	MinWidth(value interface{}) Node
	MaxWidth(value interface{}) Node
	MinHeight(value interface{}) Node
	MaxHeight(value interface{}) Node

	// Spacing
	Margin(value interface{}) Node  // Can be number (all sides), [top, right, bottom, left], or EdgeInsets
	Padding(value interface{}) Node // Can be number (all sides), [top, right, bottom, left], or EdgeInsets

	// Positioning
	Position(value string) Node
	Top(value interface{}) Node    // Can be number, percentage string, etc.
	Right(value interface{}) Node  // Can be number, percentage string, etc.
	Bottom(value interface{}) Node // Can be number, percentage string, etc.
	Left(value interface{}) Node   // Can be number, percentage string, etc.
	ZIndex(value int) Node

	// Flex layout
	FlexDirection(value string) Node
	JustifyContent(value string) Node
	AlignItems(value string) Node
	FlexWrap(value string) Node

	// Typography
	FontFamily(value string) Node
	FontSize(value interface{}) Node   // Can be number, "em" string, "rem" string, etc.
	FontWeight(value interface{}) Node // Can be number, "bold", "normal", etc.
	LineHeight(value interface{}) Node // Can be number, "em" string, etc.
	TextAlign(value string) Node
	Color(value interface{}) Node // Can be Color object, color name string, hex string, etc.

	// Visual styling
	Background(value interface{}) Node   // Can be Color object, color name string, hex string, etc.
	Border(value interface{}) Node       // Can be BorderStyle object, or individual components
	BorderRadius(value interface{}) Node // Can be number (all corners), [top, right, bottom, left], or EdgeInsets
	Shadow(value interface{}) Node       // Can be ShadowStyle object, or individual components
	Opacity(value interface{}) Node      // Can be number, percentage string, etc.
	Scale(value interface{}) Node        // Can be number, percentage string, etc.
}

// Implementation of style builder methods for BaseNode
func (n *BaseNode) Width(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("width", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("width", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if v == "auto" {
			n.styles.Set("width", style.StyleValue{Type: style.AUTO, Value: 0, Source: style.Explicit})
		} else if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("width", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("width", value)
	}
	return n
}

func (n *BaseNode) Height(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("height", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("height", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if v == "auto" {
			n.styles.Set("height", style.StyleValue{Type: style.AUTO, Value: 0, Source: style.Explicit})
		} else if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("height", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("height", value)
	}
	return n
}

func (n *BaseNode) MinWidth(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("minWidth", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("minWidth", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	default:
		n.styles.Set("minWidth", value)
	}
	return n
}

func (n *BaseNode) MaxWidth(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("maxWidth", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("maxWidth", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	default:
		n.styles.Set("maxWidth", value)
	}
	return n
}

func (n *BaseNode) MinHeight(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("minHeight", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("minHeight", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	default:
		n.styles.Set("minHeight", value)
	}
	return n
}

func (n *BaseNode) MaxHeight(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("maxHeight", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("maxHeight", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	default:
		n.styles.Set("maxHeight", value)
	}
	return n
}

func (n *BaseNode) Margin(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("margin", style.EdgeInsets{Top: v, Right: v, Bottom: v, Left: v})
	case int:
		n.styles.Set("margin", style.EdgeInsets{Top: float64(v), Right: float64(v), Bottom: float64(v), Left: float64(v)})
	case []float64:
		if len(v) == 4 {
			n.styles.Set("margin", style.EdgeInsets{Top: v[0], Right: v[1], Bottom: v[2], Left: v[3]})
		}
	case []int:
		if len(v) == 4 {
			n.styles.Set("margin", style.EdgeInsets{
				Top:    float64(v[0]),
				Right:  float64(v[1]),
				Bottom: float64(v[2]),
				Left:   float64(v[3]),
			})
		}
	case style.EdgeInsets:
		n.styles.Set("margin", v)
	default:
		n.styles.Set("margin", value)
	}
	return n
}

func (n *BaseNode) Padding(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("padding", style.EdgeInsets{Top: v, Right: v, Bottom: v, Left: v})
	case int:
		n.styles.Set("padding", style.EdgeInsets{Top: float64(v), Right: float64(v), Bottom: float64(v), Left: float64(v)})
	case []float64:
		if len(v) == 4 {
			n.styles.Set("padding", style.EdgeInsets{Top: v[0], Right: v[1], Bottom: v[2], Left: v[3]})
		}
	case []int:
		if len(v) == 4 {
			n.styles.Set("padding", style.EdgeInsets{
				Top:    float64(v[0]),
				Right:  float64(v[1]),
				Bottom: float64(v[2]),
				Left:   float64(v[3]),
			})
		}
	case style.EdgeInsets:
		n.styles.Set("padding", v)
	default:
		n.styles.Set("padding", value)
	}
	return n
}

func (n *BaseNode) Position(value string) Node {
	n.styles.Set("position", value)
	return n
}

func (n *BaseNode) Top(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("top", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("top", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("top", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("top", value)
	}
	return n
}

func (n *BaseNode) Right(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("right", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("right", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("right", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("right", value)
	}
	return n
}

func (n *BaseNode) Bottom(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("bottom", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("bottom", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("bottom", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("bottom", value)
	}
	return n
}

func (n *BaseNode) Left(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("left", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("left", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("left", style.StyleValue{Type: style.PERCENTAGE, Value: pct, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("left", value)
	}
	return n
}

func (n *BaseNode) ZIndex(value int) Node {
	n.styles.Set("zIndex", value)
	return n
}

func (n *BaseNode) FlexDirection(value string) Node {
	n.styles.Set("flexDirection", value)
	return n
}

func (n *BaseNode) JustifyContent(value string) Node {
	n.styles.Set("justifyContent", value)
	return n
}

func (n *BaseNode) AlignItems(value string) Node {
	n.styles.Set("alignItems", value)
	return n
}

func (n *BaseNode) FlexWrap(value string) Node {
	n.styles.Set("flexWrap", value)
	return n
}

func (n *BaseNode) FontFamily(value string) Node {
	n.styles.Set("fontFamily", value)
	return n
}

func (n *BaseNode) FontSize(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("fontSize", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("fontSize", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "em") {
			if em, err := strconv.ParseFloat(v[:len(v)-2], 64); err == nil {
				n.styles.Set("fontSize", style.StyleValue{Type: style.EM, Value: em, Source: style.Explicit})
			}
		} else if strings.HasSuffix(v, "rem") {
			if rem, err := strconv.ParseFloat(v[:len(v)-3], 64); err == nil {
				n.styles.Set("fontSize", style.StyleValue{Type: style.REM, Value: rem, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("fontSize", value)
	}
	return n
}

func (n *BaseNode) FontWeight(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("fontWeight", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("fontWeight", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if v == "bold" {
			n.styles.Set("fontWeight", style.StyleValue{Type: style.PIXEL, Value: 700.0, Source: style.Explicit})
		} else if v == "normal" {
			n.styles.Set("fontWeight", style.StyleValue{Type: style.PIXEL, Value: 400.0, Source: style.Explicit})
		}
	default:
		n.styles.Set("fontWeight", value)
	}
	return n
}

func (n *BaseNode) LineHeight(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("lineHeight", style.StyleValue{Type: style.EM, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("lineHeight", style.StyleValue{Type: style.EM, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "em") {
			if em, err := strconv.ParseFloat(v[:len(v)-2], 64); err == nil {
				n.styles.Set("lineHeight", style.StyleValue{Type: style.EM, Value: em, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("lineHeight", value)
	}
	return n
}

func (n *BaseNode) TextAlign(value string) Node {
	n.styles.Set("textAlign", value)
	return n
}

func (n *BaseNode) Color(value interface{}) Node {
	switch v := value.(type) {
	case style.Color:
		n.styles.Set("color", v)
	case string:
		if color, ok := parseColorString(v); ok {
			n.styles.Set("color", color)
		} else {
			n.styles.Set("color", v)
		}
	default:
		n.styles.Set("color", value)
	}
	return n
}

func (n *BaseNode) Background(value interface{}) Node {
	switch v := value.(type) {
	case style.Color:
		n.styles.Set("background", v)
	case string:
		if color, ok := parseColorString(v); ok {
			n.styles.Set("background", color)
		} else {
			n.styles.Set("background", v)
		}
	default:
		n.styles.Set("background", value)
	}
	return n
}

func (n *BaseNode) Border(value interface{}) Node {
	switch v := value.(type) {
	case style.BorderStyle:
		n.styles.Set("border", v)
	case []interface{}:
		if len(v) == 3 {
			width, _ := v[0].(style.EdgeInsets)
			styleStr, _ := v[1].(string)
			color, _ := v[2].(style.Color)
			n.styles.Set("border", style.BorderStyle{Width: width, Style: styleStr, Color: color})
		}
	default:
		n.styles.Set("border", value)
	}
	return n
}

func (n *BaseNode) BorderRadius(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("borderRadius", style.EdgeInsets{Top: v, Right: v, Bottom: v, Left: v})
	case int:
		n.styles.Set("borderRadius", style.EdgeInsets{Top: float64(v), Right: float64(v), Bottom: float64(v), Left: float64(v)})
	case []float64:
		if len(v) == 4 {
			n.styles.Set("borderRadius", style.EdgeInsets{Top: v[0], Right: v[1], Bottom: v[2], Left: v[3]})
		}
	case []int:
		if len(v) == 4 {
			n.styles.Set("borderRadius", style.EdgeInsets{
				Top:    float64(v[0]),
				Right:  float64(v[1]),
				Bottom: float64(v[2]),
				Left:   float64(v[3]),
			})
		}
	case style.EdgeInsets:
		n.styles.Set("borderRadius", v)
	default:
		n.styles.Set("borderRadius", value)
	}
	return n
}

func (n *BaseNode) Shadow(value interface{}) Node {
	switch v := value.(type) {
	case style.ShadowStyle:
		n.styles.Set("shadow", v)
	case []interface{}:
		if len(v) == 5 {
			offsetX, _ := v[0].(float64)
			offsetY, _ := v[1].(float64)
			blurRadius, _ := v[2].(float64)
			spreadRadius, _ := v[3].(float64)
			color, _ := v[4].(style.Color)
			n.styles.Set("shadow", style.ShadowStyle{
				OffsetX:      offsetX,
				OffsetY:      offsetY,
				BlurRadius:   blurRadius,
				SpreadRadius: spreadRadius,
				Color:        color,
			})
		}
	default:
		n.styles.Set("shadow", value)
	}
	return n
}

func (n *BaseNode) Opacity(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("opacity", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("opacity", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("opacity", style.StyleValue{Type: style.PERCENTAGE, Value: pct / 100.0, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("opacity", value)
	}
	return n
}

func (n *BaseNode) Scale(value interface{}) Node {
	switch v := value.(type) {
	case float64:
		n.styles.Set("scale", style.StyleValue{Type: style.PIXEL, Value: v, Source: style.Explicit})
	case int:
		n.styles.Set("scale", style.StyleValue{Type: style.PIXEL, Value: float64(v), Source: style.Explicit})
	case string:
		if strings.HasSuffix(v, "%") {
			if pct, err := strconv.ParseFloat(v[:len(v)-1], 64); err == nil {
				n.styles.Set("scale", style.StyleValue{Type: style.PERCENTAGE, Value: pct / 100.0, Source: style.Explicit})
			}
		}
	default:
		n.styles.Set("scale", value)
	}
	return n
}

// Common color names mapped to their hex values
var colorNames = map[string]string{
	"black":   "#000000",
	"white":   "#FFFFFF",
	"red":     "#FF0000",
	"green":   "#00FF00",
	"blue":    "#0000FF",
	"yellow":  "#FFFF00",
	"cyan":    "#00FFFF",
	"magenta": "#FF00FF",
	"gray":    "#808080",
	"grey":    "#808080",
	"purple":  "#800080",
	"orange":  "#FFA500",
	"pink":    "#FFC0CB",
	"brown":   "#A52A2A",
	"silver":  "#C0C0C0",
	"gold":    "#FFD700",
}

// Helper function to parse color strings (hex, color names, etc.)
func parseColorString(s string) (style.Color, bool) {
	// Convert to lowercase for case-insensitive matching
	s = strings.ToLower(strings.TrimSpace(s))

	// Check if it's a named color
	if hex, ok := colorNames[s]; ok {
		return parseHexColor(hex)
	}

	// Check if it's a hex color
	if strings.HasPrefix(s, "#") {
		return parseHexColor(s)
	}

	// Check if it's an rgb/rgba function
	if strings.HasPrefix(s, "rgb(") || strings.HasPrefix(s, "rgba(") {
		return parseRGBFunction(s)
	}

	return style.Color{}, false
}

// parseHexColor parses a hex color string (#RGB, #RGBA, #RRGGBB, #RRGGBBAA)
func parseHexColor(hex string) (style.Color, bool) {
	hex = strings.TrimPrefix(hex, "#")

	var r, g, b, a uint8
	a = 255 // Default alpha to 255 (fully opaque)

	switch len(hex) {
	case 3: // #RGB
		r = parseHexByte(hex[0:1] + hex[0:1])
		g = parseHexByte(hex[1:2] + hex[1:2])
		b = parseHexByte(hex[2:3] + hex[2:3])
	case 4: // #RGBA
		r = parseHexByte(hex[0:1] + hex[0:1])
		g = parseHexByte(hex[1:2] + hex[1:2])
		b = parseHexByte(hex[2:3] + hex[2:3])
		a = parseHexByte(hex[3:4] + hex[3:4])
	case 6: // #RRGGBB
		r = parseHexByte(hex[0:2])
		g = parseHexByte(hex[2:4])
		b = parseHexByte(hex[4:6])
	case 8: // #RRGGBBAA
		r = parseHexByte(hex[0:2])
		g = parseHexByte(hex[2:4])
		b = parseHexByte(hex[4:6])
		a = parseHexByte(hex[6:8])
	default:
		return style.Color{}, false
	}

	return style.Color{R: r, G: g, B: b, A: a}, true
}

// parseHexByte converts a hex byte string to a uint8
func parseHexByte(hex string) uint8 {
	val, _ := strconv.ParseInt(hex, 16, 64)
	return uint8(val)
}

// parseRGBFunction parses rgb() and rgba() function strings
func parseRGBFunction(s string) (style.Color, bool) {
	// Remove rgb( or rgba( and the closing )
	s = strings.TrimPrefix(s, "rgb(")
	s = strings.TrimPrefix(s, "rgba(")
	s = strings.TrimSuffix(s, ")")

	// Split by commas and trim spaces
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	var r, g, b, a uint8
	a = 255 // Default alpha to 255 (fully opaque)

	// Parse RGB values
	if len(parts) < 3 {
		return style.Color{}, false
	}

	// Parse R
	if strings.HasSuffix(parts[0], "%") {
		r = uint8(parsePercentage(parts[0]) * 2.55) // Convert percentage to 0-255
	} else {
		r = uint8(parseNumber(parts[0]))
	}

	// Parse G
	if strings.HasSuffix(parts[1], "%") {
		g = uint8(parsePercentage(parts[1]) * 2.55) // Convert percentage to 0-255
	} else {
		g = uint8(parseNumber(parts[1]))
	}

	// Parse B
	if strings.HasSuffix(parts[2], "%") {
		b = uint8(parsePercentage(parts[2]) * 2.55) // Convert percentage to 0-255
	} else {
		b = uint8(parseNumber(parts[2]))
	}

	// Parse A if present (rgba)
	if len(parts) > 3 {
		// Alpha is typically 0-1 in rgba(), so multiply by 255
		a = uint8(parseNumber(parts[3]) * 255)
	}

	return style.Color{R: r, G: g, B: b, A: a}, true
}

// parsePercentage converts a percentage string to a float64
func parsePercentage(s string) float64 {
	s = strings.TrimSuffix(s, "%")
	val, _ := strconv.ParseFloat(s, 64)
	return val
}

// parseNumber converts a number string to a float64
func parseNumber(s string) float64 {
	val, _ := strconv.ParseFloat(s, 64)
	return val
}
