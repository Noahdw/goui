package style

// size represents width and height dimensions
type size struct {
	Width, Height float64
}

// point represents x and y coordinates
type point struct {
	X, Y float64
}

// rect represents a rectangle with position and size
type rect struct {
	Position point
	Size     size
}

// Intersects checks if this rectangle intersects with another rectangle
func (r *rect) Intersects(other rect) bool {
	aMaxX := r.Position.X + r.Size.Width
	aMaxY := r.Position.Y + r.Size.Height
	bMaxX := other.Position.X + other.Size.Width
	bMaxY := other.Position.Y + other.Size.Height

	// a is left of b
	if aMaxX < other.Position.X {
		return false
	}

	// a is right of b
	if r.Position.X > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < other.Position.Y {
		return false
	}

	// a is below b
	if r.Position.Y > bMaxY {
		return false
	}

	// The two overlap
	return true
}

// NewSize creates a new size with the given dimensions
func NewSize(width, height float64) size {
	return size{Width: width, Height: height}
}

// NewPoint creates a new point with the given coordinates
func NewPoint(x, y float64) point {
	return point{X: x, Y: y}
}

// NewRect creates a new rect with the given position and size
func NewRect(position point, size size) rect {
	return rect{Position: position, Size: size}
}
