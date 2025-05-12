package style

import (
	"fmt"
	"strconv"
	"strings"
)

// styles holds all styling information for a node
type styles struct {
	// Core style properties
	properties map[string]styleValue

	// State-based style variations
	stateStyles map[string]*styles

	// Track which properties were explicitly set
	setProperties map[string]styleSource

	// Final computed opacity
	finalOpacity float64

	// Store original values for state-based style changes
	originalValues map[string]interface{}
}

// newStyles creates a new styles instance with the given properties
func newStyles(props map[string]interface{}) styles {
	s := styles{
		properties:     make(map[string]styleValue),
		stateStyles:    make(map[string]*styles),
		setProperties:  make(map[string]styleSource),
		originalValues: make(map[string]interface{}),
		finalOpacity:   1.0,
	}

	// Set default values
	for key, value := range defaultStyleValues {
		if err := s.set(key, value); err != nil {
			fmt.Printf("[STYLE ERROR] Error setting default style %s: %v\n", key, err)
		}
	}

	// Apply any custom properties
	for key, value := range props {
		if err := s.set(key, value); err != nil {
			fmt.Printf("[STYLE ERROR] Error setting style %s: %v\n", key, err)
		}
	}

	return s
}

// set sets a style property value
func (s *styles) set(key string, value interface{}) error {
	// Convert value to styleValue if needed
	var sv styleValue

	switch v := value.(type) {
	case styleValue:
		sv = v
	case int:
		sv = styleValue{Type: pixel, Value: float64(v), Source: explicit}
	case float64:
		sv = styleValue{Type: pixel, Value: v, Source: explicit}
	case string:
		if v == "auto" {
			sv = styleValue{Type: auto, Value: 0, Source: explicit}
		} else if strings.HasSuffix(v, "%") {
			pctVal, err := strconv.ParseFloat(v[:len(v)-1], 64)
			if err != nil {
				return styleError{Property: key, Value: v, Message: "Invalid percentage value"}
			}
			sv = styleValue{Type: percentage, Value: pctVal, Source: explicit}
		} else if strings.HasSuffix(v, "em") {
			emVal, err := strconv.ParseFloat(v[:len(v)-2], 64)
			if err != nil {
				return styleError{Property: key, Value: v, Message: "Invalid em value"}
			}
			sv = styleValue{Type: em, Value: emVal, Source: explicit}
		} else if strings.HasSuffix(v, "rem") {
			remVal, err := strconv.ParseFloat(v[:len(v)-3], 64)
			if err != nil {
				return styleError{Property: key, Value: v, Message: "Invalid rem value"}
			}
			sv = styleValue{Type: rem, Value: remVal, Source: explicit}
		} else {
			// For numeric properties (width, height, etc.), require explicit numeric values
			if isNumericProperty(key) {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					sv = styleValue{Type: pixel, Value: num, Source: explicit}
				} else {
					return styleError{
						Property: key,
						Value:    v,
						Message:  "Numeric properties require explicit numeric values (e.g., 100 instead of \"100\")",
					}
				}
			} else {
				// For non-numeric properties, try to parse as number first
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					sv = styleValue{Type: pixel, Value: num, Source: explicit}
				} else {
					// Treat as a string value
					sv = styleValue{Type: pixel, Value: v, Source: explicit}
				}
			}
		}
	default:
		// For other types (color, edgeInsets, etc.), store as is
		sv = styleValue{Type: pixel, Value: v, Source: explicit}
	}

	s.properties[key] = sv
	s.setProperties[key] = explicit
	return nil
}

// get gets a style property value
func (s *styles) get(key string) (interface{}, bool) {
	if value, ok := s.properties[key]; ok {
		return value.Value, true
	}
	return nil, false
}

// getFloat gets a style property as a float64
func (s *styles) getFloat(key string) (float64, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case float64:
			return v, true
		case int:
			return float64(v), true
		case styleValue:
			if f, ok := v.Value.(float64); ok {
				return f, true
			}
		}
	}
	return 0, false
}

// getString gets a style property as a string
func (s *styles) getString(key string) (string, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case string:
			return v, true
		case styleValue:
			if str, ok := v.Value.(string); ok {
				return str, true
			}
		}
	}
	return "", false
}

// getColor gets a style property as a color
func (s *styles) getColor(key string) (color, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case color:
			return v, true
		case styleValue:
			if c, ok := v.Value.(color); ok {
				return c, true
			}
		}
	}
	return color{}, false
}

// getEdgeInsets gets a style property as edgeInsets
func (s *styles) getEdgeInsets(key string) (edgeInsets, bool) {
	if value, ok := s.properties[key]; ok {
		switch v := value.Value.(type) {
		case edgeInsets:
			return v, true
		case styleValue:
			if e, ok := v.Value.(edgeInsets); ok {
				return e, true
			}
		}
	}
	return edgeInsets{}, false
}

// addStateStyle adds a style variation for a specific state
func (s *styles) addStateStyle(state string, style *styles) {
	s.stateStyles[state] = style
}

// getStateStyle returns the style variation for a specific state
func (s *styles) getStateStyle(state string) *styles {
	return s.stateStyles[state]
}

// restoreOriginalStyles restores the original styles when a state is removed
func (s *styles) restoreOriginalStyles() {
	if s.originalValues == nil {
		return
	}

	for key, value := range s.originalValues {
		s.set(key, value)
	}
	s.originalValues = nil
}

// markPropertyExplicit marks a property as explicitly set
func (s *styles) markPropertyExplicit(prop styleProperty) {
	if s.setProperties == nil {
		s.setProperties = make(map[string]styleSource)
	}
	s.setProperties[string(prop)] = explicit
}
