package shapes

// RawRect represents raw rectangle input
type RawRect struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

// Build Rect builds a rectangle
func (r RawRect) BuildRect() Rectangle {
	p1 := Point{X: r.X, Y: r.Y}
	p2 := Point{X: r.X + r.W, Y: r.Y}
	p3 := Point{X: r.X + r.W, Y: r.Y - r.H}
	p4 := Point{X: r.X, Y: r.Y - r.H}
	return Rectangle{p1, p2, p3, p4}
}
