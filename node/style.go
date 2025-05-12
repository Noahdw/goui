package node

import (
	"fmt"
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
	Value  interface{} // Can hold any type of value
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
	// Core style properties stored in a map
	properties map[string]StyleValue

	// State-based style variations
	stateStyles map[string]*Styles

	// Track which properties were explicitly set
	setProperties map[string]StyleSource

	// Final computed opacity
	finalOpacity float64

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

// StyleError represents an error in style value parsing
type StyleError struct {
	Property string
	Value    interface{}
	Message  string
}

func (e StyleError) Error() string {
	return fmt.Sprintf("Style error for %s: %v - %s", e.Property, e.Value, e.Message)
}

// NewStyles creates a new Styles instance with the given properties
func NewStyles(props map[string]interface{}) Styles {
	styles := Styles{
		properties:     make(map[string]StyleValue),
		stateStyles:    make(map[string]*Styles),
		setProperties:  make(map[string]StyleSource),
		originalValues: make(map[string]interface{}),
		finalOpacity:   1.0,
	}

	// Set default values
	defaultProps := map[string]interface{}{
		"width":          StyleValue{Type: AUTO, Value: 0, Source: Default},
		"height":         StyleValue{Type: AUTO, Value: 0, Source: Default},
		"minWidth":       StyleValue{Type: PIXEL, Value: 0, Source: Default},
		"minHeight":      StyleValue{Type: PIXEL, Value: 0, Source: Default},
		"maxWidth":       StyleValue{Type: AUTO, Value: 0, Source: Default},
		"maxHeight":      StyleValue{Type: AUTO, Value: 0, Source: Default},
		"margin":         EdgeInsets{0, 0, 0, 0},
		"padding":        EdgeInsets{0, 0, 0, 0},
		"position":       "relative",
		"flexDirection":  "row",
		"justifyContent": "start",
		"alignItems":     "stretch",
		"flexWrap":       "nowrap",
		"fontFamily":     "sans-serif",
		"fontSize":       StyleValue{Type: PIXEL, Value: 16, Source: Default},
		"fontWeight":     StyleValue{Type: PIXEL, Value: 400, Source: Default},
		"lineHeight":     StyleValue{Type: EM, Value: 1.2, Source: Default},
		"textAlign":      "left",
		"color":          Black,
		"background":     White,
		"border":         BorderStyle{EdgeInsets{0, 0, 0, 0}, "none", Black},
		"borderRadius":   EdgeInsets{0, 0, 0, 0},
		"shadow":         ShadowStyle{0, 0, 0, 0, Transparent},
		"opacity":        1.0,
		"scale":          1.0,
	}

	for key, value := range defaultProps {
		if err := styles.Set(key, value); err != nil {
			fmt.Printf("[STYLE ERROR] Error setting default style %s: %v\n", key, err)
		}
	}

	// Apply any custom properties
	for key, value := range props {
		if err := styles.Set(key, value); err != nil {
			fmt.Printf("[STYLE ERROR] Error setting style %s: %v\n", key, err)
		}
	}

	return styles
}

// Set sets a style property value
func (s *Styles) Set(key string, value interface{}) error {
	// Convert value to StyleValue if needed
	var styleValue StyleValue

	switch v := value.(type) {
	case StyleValue:
		styleValue = v
	case int:
		styleValue = StyleValue{Type: PIXEL, Value: float64(v), Source: Explicit}
	case float64:
		styleValue = StyleValue{Type: PIXEL, Value: v, Source: Explicit}
	case string:
		if v == "auto" {
			styleValue = StyleValue{Type: AUTO, Value: 0, Source: Explicit}
		} else if strings.HasSuffix(v, "%") {
			percentage, err := strconv.ParseFloat(v[:len(v)-1], 64)
			if err != nil {
				return StyleError{Property: key, Value: v, Message: "Invalid percentage value"}
			}
			styleValue = StyleValue{Type: PERCENTAGE, Value: percentage, Source: Explicit}
		} else if strings.HasSuffix(v, "em") {
			em, err := strconv.ParseFloat(v[:len(v)-2], 64)
			if err != nil {
				return StyleError{Property: key, Value: v, Message: "Invalid em value"}
			}
			styleValue = StyleValue{Type: EM, Value: em, Source: Explicit}
		} else if strings.HasSuffix(v, "rem") {
			rem, err := strconv.ParseFloat(v[:len(v)-3], 64)
			if err != nil {
				return StyleError{Property: key, Value: v, Message: "Invalid rem value"}
			}
			styleValue = StyleValue{Type: REM, Value: rem, Source: Explicit}
		} else {
			// For numeric properties (width, height, etc.), require explicit numeric values
			if isNumericProperty(key) {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					styleValue = StyleValue{Type: PIXEL, Value: num, Source: Explicit}
				} else {
					return StyleError{
						Property: key,
						Value:    v,
						Message:  "Numeric properties require explicit numeric values (e.g., 100 instead of \"100\")",
					}
				}
			} else {
				// For non-numeric properties, try to parse as number first
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					styleValue = StyleValue{Type: PIXEL, Value: num, Source: Explicit}
				} else {
					// Treat as a string value
					styleValue = StyleValue{Type: PIXEL, Value: v, Source: Explicit}
				}
			}
		}
	default:
		// For other types (Color, EdgeInsets, etc.), store as is
		styleValue = StyleValue{Type: PIXEL, Value: v, Source: Explicit}
	}

	s.properties[key] = styleValue
	s.setProperties[key] = Explicit
	return nil
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

// Get gets a style property value
func (s *Styles) Get(key string) (interface{}, bool) {
	if value, ok := s.properties[key]; ok {
		return value.Value, true
	}
	return nil, false
}

// GetFloat gets a style property as a float64
func (s *Styles) GetFloat(key string) (float64, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case float64:
			return v, true
		case int:
			return float64(v), true
		case StyleValue:
			if f, ok := v.Value.(float64); ok {
				return f, true
			}
		}
	}
	return 0, false
}

// GetString gets a style property as a string
func (s *Styles) GetString(key string) (string, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case string:
			return v, true
		case StyleValue:
			if str, ok := v.Value.(string); ok {
				return str, true
			}
		}
	}
	return "", false
}

// GetColor gets a style property as a Color
func (s *Styles) GetColor(key string) (Color, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case Color:
			return v, true
		case StyleValue:
			if c, ok := v.Value.(Color); ok {
				return c, true
			}
		}
	}
	return Color{}, false
}

// GetEdgeInsets gets a style property as EdgeInsets
func (s *Styles) GetEdgeInsets(key string) (EdgeInsets, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case EdgeInsets:
			return v, true
		case StyleValue:
			if e, ok := v.Value.(EdgeInsets); ok {
				return e, true
			}
		}
	}
	return EdgeInsets{}, false
}

// AddStateStyle adds a style variation for a specific state
func (s *Styles) AddStateStyle(state string, style *Styles) {
	s.stateStyles[state] = style
}

// GetStateStyle returns the style variation for a specific state
func (s *Styles) GetStateStyle(state string) *Styles {
	return s.stateStyles[state]
}

// RestoreOriginalStyles restores the original styles when a state is removed
func (s *Styles) RestoreOriginalStyles() {
	if s.originalValues == nil {
		return
	}

	for key, value := range s.originalValues {
		s.Set(key, value)
	}
	s.originalValues = nil
}

func (b BorderStyle) CanDisplay() bool {
	return b.Style != "none" && b.Width.IsNonZero()
}

func (e EdgeInsets) IsNonZero() bool {
	return (e.Top != 0 || e.Bottom != 0 || e.Right != 0 || e.Left != 0)
}

// Debug method to dump style information
func (s *Styles) DumpStyles() {
	debugLog("Style Dump:")
	debugLog("  Layout:")
	debugLog("    Width: %v", s.properties[string(WidthProp)])
	debugLog("    Height: %v", s.properties[string(HeightProp)])
	debugLog("  Visual:")
	debugLog("    Background: %v", s.properties[string(BackgroundProp)])
	debugLog("    Opacity: %v", s.finalOpacity)
	debugLog("    Scale: %v", s.properties[string(ScaleProp)])
	debugLog("  States:")
	for state, style := range s.stateStyles {
		debugLog("    %s:", state)
		debugLog("      Background: %v", style.properties[string(BackgroundProp)])
		debugLog("      Opacity: %v", style.finalOpacity)
		debugLog("      Scale: %v", style.properties[string(ScaleProp)])
	}
}

// MarkPropertyExplicit marks a property as explicitly set
func (s *Styles) MarkPropertyExplicit(prop StyleProperty) {
	if s.setProperties == nil {
		s.setProperties = make(map[string]StyleSource)
	}
	s.setProperties[string(prop)] = Explicit
}
