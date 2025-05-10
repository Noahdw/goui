package node

import (
	"strconv"
	"strings"
)

// ValueType represents the type of a style value
type ValueType int

const (
	PIXEL ValueType = iota
	PERCENTAGE
	AUTO
	EM
	REM
)

// StyleSource tracks where the style value came from
type StyleSource int

const (
	Unset StyleSource = iota
	Default
	Inherited
	Explicit
)

// StyleProperty is an enum representing each style property
type StyleProperty string

const (
	// Layout properties
	WidthProp     StyleProperty = "Width"
	HeightProp    StyleProperty = "Height"
	MinWidthProp  StyleProperty = "MinWidth"
	MaxWidthProp  StyleProperty = "MaxWidth"
	MinHeightProp StyleProperty = "MinHeight"
	MaxHeightProp StyleProperty = "MaxHeight"

	// Spacing
	MarginProp  StyleProperty = "Margin"
	PaddingProp StyleProperty = "Padding"

	// Positioning
	PositionProp StyleProperty = "Position"
	TopProp      StyleProperty = "Top"
	RightProp    StyleProperty = "Right"
	BottomProp   StyleProperty = "Bottom"
	LeftProp     StyleProperty = "Left"
	ZIndexProp   StyleProperty = "ZIndex"

	// Flex Layout
	FlexDirectionProp  StyleProperty = "FlexDirection"
	JustifyContentProp StyleProperty = "JustifyContent"
	AlignItemsProp     StyleProperty = "AlignItems"
	FlexWrapProp       StyleProperty = "FlexWrap"

	// Typography
	FontFamilyProp StyleProperty = "FontFamily"
	FontSizeProp   StyleProperty = "FontSize"
	FontWeightProp StyleProperty = "FontWeight"
	LineHeightProp StyleProperty = "LineHeight"
	TextAlignProp  StyleProperty = "TextAlign"
	ColorProp      StyleProperty = "Color"

	// Visual styling
	BackgroundProp   StyleProperty = "Background"
	BorderProp       StyleProperty = "Border"
	BorderRadiusProp StyleProperty = "BorderRadius"
	ShadowProp       StyleProperty = "Shadow"
	OpacityProp      StyleProperty = "Opacity"
	ScaleProp        StyleProperty = "Scale"
)

// StyleValue represents a value for a style property
type StyleValue struct {
	Type   ValueType
	Value  float64
	Source StyleSource
}

// Color represents an RGBA color
type Color struct {
	R, G, B, A uint8
}

// EdgeInsets represents spacing on all four sides
type EdgeInsets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// BorderStyle represents border properties
type BorderStyle struct {
	Width EdgeInsets
	Style string // "solid", "dashed", "none"
	Color Color
}

// ShadowStyle represents shadow properties
type ShadowStyle struct {
	OffsetX      float64
	OffsetY      float64
	BlurRadius   float64
	SpreadRadius float64
	Color        Color
}

// Styles holds all styling information for a node
type Styles struct {
	// Layout properties
	Width     StyleValue
	Height    StyleValue
	MinWidth  StyleValue
	MinHeight StyleValue
	MaxWidth  StyleValue
	MaxHeight StyleValue

	// Spacing
	Margin  EdgeInsets
	Padding EdgeInsets

	// Positioning
	Position string
	Left     StyleValue
	Top      StyleValue
	Right    StyleValue
	Bottom   StyleValue
	ZIndex   int

	// Flex layout
	FlexDirection  string
	JustifyContent string
	AlignItems     string
	FlexWrap       string

	// Typography
	FontFamily string
	FontSize   StyleValue
	FontWeight StyleValue
	LineHeight StyleValue
	TextAlign  string
	Color      Color

	// Visual styling
	Background   Color
	Border       BorderStyle
	BorderRadius EdgeInsets
	Shadow       ShadowStyle
	Opacity      float64
	Scale        float64

	// State-based style variations
	StateStyles map[string]*Styles

	// Track which properties were explicitly set
	setProperties map[string]StyleSource
	finalOpacity  float64

	// Store original values for state-based style changes
	originalValues map[string]interface{}
}

// StyleProps is used for initializing styles
type StyleProps struct {
	// Layout properties
	Width     *StyleValue
	Height    *StyleValue
	MinWidth  *StyleValue
	MinHeight *StyleValue
	MaxWidth  *StyleValue
	MaxHeight *StyleValue

	// Spacing
	Margin  *EdgeInsets
	Padding *EdgeInsets

	// Positioning
	Position *string
	Left     *StyleValue
	Top      *StyleValue
	Right    *StyleValue
	Bottom   *StyleValue
	ZIndex   *int

	// Flex layout
	FlexDirection  *string
	JustifyContent *string
	AlignItems     *string
	FlexWrap       *string

	// Typography
	FontFamily *string
	FontSize   *StyleValue
	FontWeight *StyleValue
	LineHeight *StyleValue
	TextAlign  *string
	Color      *Color

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
	Black       = Color{0, 0, 0, 255}
	White       = Color{255, 255, 255, 255}
	Red         = Color{255, 0, 0, 255}
	Green       = Color{0, 255, 0, 255}
	Blue        = Color{0, 0, 255, 255}
	Yellow      = Color{255, 255, 0, 255}
	Cyan        = Color{0, 255, 255, 255}
	Magenta     = Color{255, 0, 255, 255}
	Gray        = Color{128, 128, 128, 255}
	Transparent = Color{0, 0, 0, 0}

	// Extended colors
	LightGray   = Color{211, 211, 211, 255}
	DarkGray    = Color{169, 169, 169, 255}
	Silver      = Color{192, 192, 192, 255}
	DarkRed     = Color{139, 0, 0, 255}
	Navy        = Color{0, 0, 128, 255}
	ForestGreen = Color{34, 139, 34, 255}
	Orange      = Color{255, 165, 0, 255}
	Purple      = Color{128, 0, 128, 255}
)

func NewStyles(props StyleProps) Styles {
	debugLog("Creating new styles with props")

	// Create styles with default values
	styles := Styles{
		Width:     StyleValue{Type: AUTO, Value: 0, Source: Default},
		Height:    StyleValue{Type: AUTO, Value: 0, Source: Default},
		MinWidth:  StyleValue{Type: PIXEL, Value: 0, Source: Default},
		MinHeight: StyleValue{Type: PIXEL, Value: 0, Source: Default},
		MaxWidth:  StyleValue{Type: AUTO, Value: 0, Source: Default},
		MaxHeight: StyleValue{Type: AUTO, Value: 0, Source: Default},

		Margin:  EdgeInsets{0, 0, 0, 0},
		Padding: EdgeInsets{0, 0, 0, 0},

		Position: "relative",
		Left:     StyleValue{Type: AUTO, Value: 0, Source: Default},
		Top:      StyleValue{Type: AUTO, Value: 0, Source: Default},
		Right:    StyleValue{Type: AUTO, Value: 0, Source: Default},
		Bottom:   StyleValue{Type: AUTO, Value: 0, Source: Default},
		ZIndex:   0,

		FlexDirection:  "row",
		JustifyContent: "start",
		AlignItems:     "stretch",
		FlexWrap:       "nowrap",

		FontFamily: "sans-serif",
		FontSize:   StyleValue{Type: PIXEL, Value: 16, Source: Default},
		FontWeight: StyleValue{Type: PIXEL, Value: 400, Source: Default},
		LineHeight: StyleValue{Type: EM, Value: 1.2, Source: Default},
		TextAlign:  "left",
		Color:      Black,

		Background:   White,
		Border:       BorderStyle{EdgeInsets{0, 0, 0, 0}, "none", Black},
		BorderRadius: EdgeInsets{0, 0, 0, 0},
		Shadow:       ShadowStyle{0, 0, 0, 0, Transparent},
		Opacity:      1.0,
		Scale:        1.0,

		setProperties: make(map[string]StyleSource),
		finalOpacity:  1.0,

		originalValues: make(map[string]interface{}),
	}

	// Apply any properties passed in
	if props.Width != nil {
		styles.Width = *props.Width
		styles.Width.Source = Default
		styles.setProperties[string(WidthProp)] = Default
		debugLog("Applied Width: %v", styles.Width)
	}
	if props.Height != nil {
		styles.Height = *props.Height
		styles.Height.Source = Default
		styles.setProperties[string(HeightProp)] = Default
		debugLog("Applied Height: %v", styles.Height)
	}
	if props.MinWidth != nil {
		styles.MinWidth = *props.MinWidth
		styles.MinWidth.Source = Default
		styles.setProperties[string(MinWidthProp)] = Default
		debugLog("Applied MinWidth: %v", styles.MinWidth)
	}
	if props.FontFamily != nil {
		styles.FontFamily = *props.FontFamily
		styles.setProperties[string(FontFamilyProp)] = Default
		debugLog("Applied FontFamily: %s", styles.FontFamily)
	}
	if props.FontSize != nil {
		styles.FontSize = *props.FontSize
		styles.FontSize.Source = Default
		styles.setProperties[string(FontSizeProp)] = Default
		debugLog("Applied FontSize: %v", styles.FontSize)
	}
	if props.Color != nil {
		styles.Color = *props.Color
		styles.setProperties[string(ColorProp)] = Default
		debugLog("Applied Color: %v", styles.Color)
	}
	if props.Background != nil {
		styles.Background = *props.Background
		styles.setProperties[string(BackgroundProp)] = Default
		debugLog("Applied Background: %v", styles.Background)
	}
	if props.FlexDirection != nil {
		styles.FlexDirection = *props.FlexDirection
		styles.setProperties[string(FlexDirectionProp)] = Default
		debugLog("Applied FlexDirection: %s", styles.FlexDirection)
	}
	if props.Margin != nil {
		styles.Margin = *props.Margin
		styles.setProperties[string(MarginProp)] = Default
		debugLog("Applied Margin: %v", styles.Margin)
	}
	if props.Padding != nil {
		styles.Padding = *props.Padding
		styles.setProperties[string(PaddingProp)] = Default
		debugLog("Applied Padding: %v", styles.Padding)
	}
	if props.Border != nil {
		styles.Border = *props.Border
		styles.setProperties[string(BorderProp)] = Default
		debugLog("Applied Border: %v", styles.Border)
	}
	if props.BorderRadius != nil {
		styles.BorderRadius = *props.BorderRadius
		styles.setProperties[string(BorderRadiusProp)] = Default
		debugLog("Applied BorderRadius: %v", styles.BorderRadius)
	}
	if props.Scale != nil {
		styles.Scale = *props.Scale
		styles.setProperties["scale"] = Default
		debugLog("Applied Scale: %v", styles.Scale)
	}

	return styles
}

func parseStyleValue(value interface{}) StyleValue {
	switch v := value.(type) {
	case int:
		return StyleValue{Type: PIXEL, Value: float64(v), Source: Explicit}
	case float64:
		return StyleValue{Type: PIXEL, Value: v, Source: Explicit}
	case string:
		if v == "auto" {
			return StyleValue{Type: AUTO, Value: 0, Source: Explicit}
		}
		if strings.HasSuffix(v, "%") {
			percentage, _ := strconv.ParseFloat(v[:len(v)-1], 64)
			return StyleValue{Type: PERCENTAGE, Value: percentage, Source: Explicit}
		}
		if strings.HasSuffix(v, "em") {
			em, _ := strconv.ParseFloat(v[:len(v)-2], 64)
			return StyleValue{Type: EM, Value: em, Source: Explicit}
		}
		if strings.HasSuffix(v, "rem") {
			rem, _ := strconv.ParseFloat(v[:len(v)-3], 64)
			return StyleValue{Type: REM, Value: rem, Source: Explicit}
		}
		// Default to pixels if just a number
		px, _ := strconv.ParseFloat(v, 64)
		return StyleValue{Type: PIXEL, Value: px, Source: Explicit}
	}
	return StyleValue{Type: AUTO, Value: 0, Source: Explicit}
}

func (b BorderStyle) CanDisplay() bool {
	return b.Style != "none" && b.Width.IsNonZero()
}

func (e EdgeInsets) IsNonZero() bool {
	return (e.Top != 0 || e.Bottom != 0 || e.Right != 0 || e.Left != 0)
}

// AddStateStyle adds a style variation for a specific state
func (s *Styles) AddStateStyle(state string, style *Styles) {
	debugLog("Adding state style for state: %s", state)
	if s.StateStyles == nil {
		s.StateStyles = make(map[string]*Styles)
	}
	s.StateStyles[state] = style
	debugLog("State styles count: %d", len(s.StateStyles))
}

// GetStateStyle returns the style variation for a specific state
func (s *Styles) GetStateStyle(state string) *Styles {
	if s.StateStyles == nil {
		debugLog("No state styles found for state: %s", state)
		return nil
	}
	style := s.StateStyles[state]
	if style == nil {
		debugLog("No style found for state: %s", state)
	} else {
		debugLog("Found style for state: %s", state)
	}
	return style
}

// Debug method to dump style information
func (s *Styles) DumpStyles() {
	debugLog("Style Dump:")
	debugLog("  Layout:")
	debugLog("    Width: %v", s.Width)
	debugLog("    Height: %v", s.Height)
	debugLog("  Visual:")
	debugLog("    Background: %v", s.Background)
	debugLog("    Opacity: %v", s.Opacity)
	debugLog("    Scale: %v", s.Scale)
	debugLog("  States:")
	for state, style := range s.StateStyles {
		debugLog("    %s:", state)
		debugLog("      Background: %v", style.Background)
		debugLog("      Opacity: %v", style.Opacity)
		debugLog("      Scale: %v", style.Scale)
	}
}
