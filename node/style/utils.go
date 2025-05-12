package style

import "fmt"

// debugLog is a helper function for style debugging
func debugLog(format string, args ...interface{}) {
	fmt.Printf("[STYLE DEBUG] "+format+"\n", args...)
}

// dumpStyles is a debug method to dump style information
func (s *styles) dumpStyles() {
	debugLog("Style Dump:")
	debugLog("  Layout:")
	debugLog("    Width: %v", s.properties[string(widthProp)])
	debugLog("    Height: %v", s.properties[string(heightProp)])
	debugLog("  Visual:")
	debugLog("    Background: %v", s.properties[string(backgroundProp)])
	debugLog("    Opacity: %v", s.finalOpacity)
	debugLog("    Scale: %v", s.properties[string(scaleProp)])
	debugLog("  States:")
	for state, style := range s.stateStyles {
		debugLog("    %s:", state)
		debugLog("      Background: %v", style.properties[string(backgroundProp)])
		debugLog("      Opacity: %v", style.finalOpacity)
		debugLog("      Scale: %v", style.properties[string(scaleProp)])
	}
}
