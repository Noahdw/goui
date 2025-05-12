package style

import (
	"fmt"
)

// valueType represents the type of a style value
type valueType int

const (
	pixel valueType = iota
	percentage
	auto
	em
	rem
)

// styleSource tracks where the style value came from
type styleSource int

const (
	unset styleSource = iota
	default_
	inherited
	explicit
)

// styleProperty is an enum representing each style property
type styleProperty string

const (
	// Layout properties
	widthProp     styleProperty = "Width"
	heightProp    styleProperty = "Height"
	minWidthProp  styleProperty = "MinWidth"
	maxWidthProp  styleProperty = "MaxWidth"
	minHeightProp styleProperty = "MinHeight"
	maxHeightProp styleProperty = "MaxHeight"

	// Spacing
	marginProp  styleProperty = "Margin"
	paddingProp styleProperty = "Padding"

	// Positioning
	positionProp styleProperty = "Position"
	topProp      styleProperty = "Top"
	rightProp    styleProperty = "Right"
	bottomProp   styleProperty = "Bottom"
	leftProp     styleProperty = "Left"
	zIndexProp   styleProperty = "ZIndex"

	// Flex Layout
	flexDirectionProp  styleProperty = "FlexDirection"
	justifyContentProp styleProperty = "JustifyContent"
	alignItemsProp     styleProperty = "AlignItems"
	flexWrapProp       styleProperty = "FlexWrap"

	// Typography
	fontFamilyProp styleProperty = "FontFamily"
	fontSizeProp   styleProperty = "FontSize"
	fontWeightProp styleProperty = "FontWeight"
	lineHeightProp styleProperty = "LineHeight"
	textAlignProp  styleProperty = "TextAlign"
	colorProp      styleProperty = "Color"

	// Visual styling
	backgroundProp   styleProperty = "Background"
	borderProp       styleProperty = "Border"
	borderRadiusProp styleProperty = "BorderRadius"
	shadowProp       styleProperty = "Shadow"
	opacityProp      styleProperty = "Opacity"
	scaleProp        styleProperty = "Scale"
)

// styleValue represents a value for a style property
type styleValue struct {
	Type   valueType
	Value  interface{} // Can hold any type of value
	Source styleSource
}

// color represents an RGBA color
type color struct {
	R, G, B, A uint8
}

// edgeInsets represents spacing on all four sides
type edgeInsets struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

// shadowStyle represents shadow properties
type shadowStyle struct {
	OffsetX      float64
	OffsetY      float64
	BlurRadius   float64
	SpreadRadius float64
	Color        color
}

// styleError represents an error in style value parsing
type styleError struct {
	Property string
	Value    interface{}
	Message  string
}

func (e styleError) Error() string {
	return fmt.Sprintf("Style error for %s: %v - %s", e.Property, e.Value, e.Message)
}

func (e edgeInsets) isNonZero() bool {
	return (e.Top != 0 || e.Bottom != 0 || e.Right != 0 || e.Left != 0)
}

// BorderStyle represents border properties
type BorderStyle struct {
	Width EdgeInsets
	Style string // "solid", "dashed", "none"
	Color Color
}

// CanDisplay returns true if the border should be displayed
func (b BorderStyle) CanDisplay() bool {
	return b.Style != "none" && b.Width.IsNonZero()
}

// IsNonZero returns true if any of the edge insets are non-zero
func (e EdgeInsets) IsNonZero() bool {
	return (e.Top != 0 || e.Bottom != 0 || e.Right != 0 || e.Left != 0)
}
