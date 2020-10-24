package shapes

import (
	"fmt"

	"github.com/golang/geo/r2"
)

const (
	rawRectInfoString = "Rectangle at (%+v,%+v), w=%d, h=%d."
)

// RawRect represents raw rectangle input
type RawRect struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	H int     `json:"h"`
	W int     `json:"w"`
}

func (rr RawRect) BuildInfoText() string {
	return fmt.Sprintf(rawRectInfoString, rr.X, rr.Y, rr.W, rr.H)
}

// BuildR2Rect builds an r2 rectangle
func (r RawRect) BuildR2Rect() r2.Rect {
	p1 := r2.Point{
		X: r.X,
		Y: r.Y,
	}

	p2 := r2.Point{
		X: r.X + float64(r.W),
		Y: r.Y,
	}
	p3 := r2.Point{
		X: r.X + float64(r.W),
		Y: r.Y - float64(r.H),
	}
	p4 := r2.Point{
		X: r.X,
		Y: r.Y - float64(r.H),
	}

	return r2.RectFromPoints(p1, p2, p3, p4)
}

func BuildRawRect(rect r2.Rect) RawRect {
	r := RawRect{}
	r.X = rect.Lo().X
	r.Y = rect.Hi().Y
	r.W = int(rect.Hi().X - rect.Lo().X)
	r.H = int(rect.Hi().Y - rect.Lo().Y)
	return r
}
