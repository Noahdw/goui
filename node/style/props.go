package style

// styleProps is used for initializing styles
type styleProps struct {
	// Layout properties
	Width     *styleValue
	Height    *styleValue
	MinWidth  *styleValue
	MinHeight *styleValue
	MaxWidth  *styleValue
	MaxHeight *styleValue

	// Spacing
	Margin  *edgeInsets
	Padding *edgeInsets

	// Positioning
	Position *string
	Left     *styleValue
	Top      *styleValue
	Right    *styleValue
	Bottom   *styleValue
	ZIndex   *int

	// Flex layout
	FlexDirection  *string
	JustifyContent *string
	AlignItems     *string
	FlexWrap       *string

	// Typography
	FontFamily *string
	FontSize   *styleValue
	FontWeight *styleValue
	LineHeight *styleValue
	TextAlign  *string
	Color      *color

	// Visual styling
	Background   *Color
	Border       *BorderStyle
	BorderRadius *EdgeInsets
	Shadow       *ShadowStyle
	Opacity      *float64
	Scale        *float64
}

// Standard color definitions
var (
	// Basic colors
	black       = color{0, 0, 0, 255}
	white       = color{255, 255, 255, 255}
	red         = color{255, 0, 0, 255}
	green       = color{0, 255, 0, 255}
	blue        = color{0, 0, 255, 255}
	yellow      = color{255, 255, 0, 255}
	cyan        = color{0, 255, 255, 255}
	magenta     = color{255, 0, 255, 255}
	gray        = color{128, 128, 128, 255}
	transparent = color{0, 0, 0, 0}

	// Extended colors
	lightGray   = color{211, 211, 211, 255}
	darkGray    = color{169, 169, 169, 255}
	silver      = color{192, 192, 192, 255}
	darkRed     = color{139, 0, 0, 255}
	navy        = color{0, 0, 128, 255}
	forestGreen = color{34, 139, 34, 255}
	orange      = color{255, 165, 0, 255}
	purple      = color{128, 0, 128, 255}
)

// Default style values
var defaultStyleValues = map[string]interface{}{
	"width":          styleValue{Type: auto, Value: 0, Source: default_},
	"height":         styleValue{Type: auto, Value: 0, Source: default_},
	"minWidth":       styleValue{Type: pixel, Value: 0, Source: default_},
	"minHeight":      styleValue{Type: pixel, Value: 0, Source: default_},
	"maxWidth":       styleValue{Type: auto, Value: 0, Source: default_},
	"maxHeight":      styleValue{Type: auto, Value: 0, Source: default_},
	"margin":         edgeInsets{0, 0, 0, 0},
	"padding":        edgeInsets{0, 0, 0, 0},
	"position":       "relative",
	"flexDirection":  "row",
	"justifyContent": "start",
	"alignItems":     "stretch",
	"flexWrap":       "nowrap",
	"fontFamily":     "sans-serif",
	"fontSize":       styleValue{Type: pixel, Value: 16, Source: default_},
	"fontWeight":     styleValue{Type: pixel, Value: 400, Source: default_},
	"lineHeight":     styleValue{Type: em, Value: 1.2, Source: default_},
	"textAlign":      "left",
	"color":          black,
	"background":     white,
	"border":         BorderStyle{Width: EdgeInsets{0, 0, 0, 0}, Style: "none", Color: Black},
	"borderRadius":   edgeInsets{0, 0, 0, 0},
	"shadow":         shadowStyle{0, 0, 0, 0, transparent},
	"opacity":        1.0,
	"scale":          1.0,
}

// isNumericProperty returns true if the property typically expects a numeric value
func isNumericProperty(key string) bool {
	numericProps := map[string]bool{
		"width":      true,
		"height":     true,
		"minWidth":   true,
		"maxWidth":   true,
		"minHeight":  true,
		"maxHeight":  true,
		"fontSize":   true,
		"lineHeight": true,
		"opacity":    true,
		"scale":      true,
		"zIndex":     true,
	}
	return numericProps[key]
}
