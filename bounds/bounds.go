package bounds

// Bounds - A bounding box with a x,y origin and width and height
type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

//IsPoint - Checks if a bounds object is a point or not (has no width or height)
func (b *Bounds) IsPoint() bool {

	if b.Width == 0 && b.Height == 0 {
		return true
	}

	return false

}

func (b *Bounds) BoundingRect() *Bounds {
	return b
}
