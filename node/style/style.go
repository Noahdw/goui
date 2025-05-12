// Package style provides a comprehensive styling system for UI nodes.
// It includes support for:
// - Layout properties (width, height, margins, etc.)
// - Visual styling (colors, borders, shadows, etc.)
// - Typography (font family, size, weight, etc.)
// - Flex layout (direction, alignment, etc.)
// - State-based styling
//
// The package is organized into several files:
// - types.go: Core type definitions and constants
// - props.go: Style properties and default values
// - manager.go: Style management and computation
// - utils.go: Debugging and utility functions
package style

// Re-export commonly used types and functions
type (
	ValueType     = valueType
	StyleSource   = styleSource
	StyleProperty = styleProperty
	StyleValue    = styleValue
	Color         = color
	EdgeInsets    = edgeInsets
	ShadowStyle   = shadowStyle
	StyleError    = styleError
	Styles        = styles
	StyleProps    = styleProps
	Size          = size
	Point         = point
	Rect          = rect
)

// Re-export commonly used constants
const (
	// Value types
	PIXEL      = pixel
	PERCENTAGE = percentage
	AUTO       = auto
	EM         = em
	REM        = rem

	// Style sources
	Unset     = unset
	Default   = default_
	Inherited = inherited
	Explicit  = explicit

	// Layout properties
	WidthProp     = widthProp
	HeightProp    = heightProp
	MinWidthProp  = minWidthProp
	MaxWidthProp  = maxWidthProp
	MinHeightProp = minHeightProp
	MaxHeightProp = maxHeightProp

	// Spacing
	MarginProp  = marginProp
	PaddingProp = paddingProp

	// Positioning
	PositionProp = positionProp
	TopProp      = topProp
	RightProp    = rightProp
	BottomProp   = bottomProp
	LeftProp     = leftProp
	ZIndexProp   = zIndexProp

	// Flex Layout
	FlexDirectionProp  = flexDirectionProp
	JustifyContentProp = justifyContentProp
	AlignItemsProp     = alignItemsProp
	FlexWrapProp       = flexWrapProp

	// Typography
	FontFamilyProp = fontFamilyProp
	FontSizeProp   = fontSizeProp
	FontWeightProp = fontWeightProp
	LineHeightProp = lineHeightProp
	TextAlignProp  = textAlignProp
	ColorProp      = colorProp

	// Visual styling
	BackgroundProp   = backgroundProp
	BorderProp       = borderProp
	BorderRadiusProp = borderRadiusProp
	ShadowProp       = shadowProp
	OpacityProp      = opacityProp
	ScaleProp        = scaleProp
)

// Re-export commonly used variables
var (
	// Basic colors
	Black       = black
	White       = white
	Red         = red
	Green       = green
	Blue        = blue
	Yellow      = yellow
	Cyan        = cyan
	Magenta     = magenta
	Gray        = gray
	Transparent = transparent

	// Extended colors
	LightGray   = lightGray
	DarkGray    = darkGray
	Silver      = silver
	DarkRed     = darkRed
	Navy        = navy
	ForestGreen = forestGreen
	Orange      = orange
	Purple      = purple
)

// Re-export commonly used functions
var (
	NewStyles = newStyles
)

// Re-export methods
func (s *Styles) Set(key string, value interface{}) error {
	return s.set(key, value)
}

func (s *Styles) Get(key string) (interface{}, bool) {
	return s.get(key)
}

func (s *Styles) GetFloat(key string) (float64, bool) {
	return s.getFloat(key)
}

func (s *Styles) GetString(key string) (string, bool) {
	return s.getString(key)
}

func (s *Styles) GetColor(key string) (Color, bool) {
	return s.getColor(key)
}

func (s *Styles) GetEdgeInsets(key string) (EdgeInsets, bool) {
	return s.getEdgeInsets(key)
}

func (s *Styles) AddStateStyle(state string, style *Styles) {
	s.addStateStyle(state, style)
}

func (s *Styles) GetStateStyle(state string) *Styles {
	return s.getStateStyle(state)
}

func (s *Styles) RestoreOriginalStyles() {
	s.restoreOriginalStyles()
}

func (s *Styles) MarkPropertyExplicit(prop StyleProperty) {
	s.markPropertyExplicit(prop)
}

func (s *Styles) DumpStyles() {
	s.dumpStyles()
}

// IsExplicit returns true if the property was explicitly set
func (s *Styles) IsExplicit(key string) bool {
	if s == nil {
		return false
	}
	if s.setProperties == nil {
		return false
	}
	source, ok := s.setProperties[key]
	return ok && source == Explicit
}

// GetFinalOpacity returns the final computed opacity
func (s *Styles) GetFinalOpacity() float64 {
	return s.finalOpacity
}

// SetFinalOpacity sets the final computed opacity
func (s *Styles) SetFinalOpacity(opacity float64) {
	s.finalOpacity = opacity
}

// StoreOriginalValue stores the original value for a property (for state style changes)
func (s *Styles) StoreOriginalValue(key string, value interface{}) {
	if s.originalValues == nil {
		s.originalValues = make(map[string]interface{})
	}
	if _, exists := s.originalValues[key]; !exists {
		s.originalValues[key] = value
	}
}
