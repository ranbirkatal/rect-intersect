package shapes_test

import (
	"reflect"
	"testing"

	"github.com/golang/geo/r2"
	"github.com/ranbirkatal/rect-intersect/pkg/shapes"
)

var (
	p1 = r2.Point{float64(10), float64(20)}
	p2 = r2.Point{float64(20), float64(10)}

	p3 = r2.Point{float64(15), float64(15)}
	p4 = r2.Point{float64(25), float64(5)}

	p5 = r2.Point{float64(30), float64(25)}
	p6 = r2.Point{float64(40), float64(20)}
)

func TestIntersect(t *testing.T) {
	rect1 := r2.RectFromPoints(p1, p2)
	rect2 := r2.RectFromPoints(p3, p4)
	rect3 := r2.RectFromPoints(p5, p6)
	items := shapes.BuildInitialIntersections([]r2.Rect{rect1, rect2, rect3})
	if len(items) != 3 {
		t.Logf("expected 3 items but got %d", len(items))
		t.Fail()
	}
	result := shapes.Intersect(items)

	if len(result) != 1 {
		t.Logf("expected 1 intersection but got %d", len(result))
		t.Fail()
	}

	isx := result[0]
	if !reflect.DeepEqual([]int{0, 1}, isx.IDXs) {
		t.Logf("Invalid intersection found, expected [0 1] got %+v", isx.IDXs)
		t.Fail()
	}
	rawRect := shapes.BuildRawRect(isx.Rect)

	if rawRect.X != float64(15) {
		t.Errorf("expected first x cordinate to be 15 but got %+v", rawRect.X)
		t.Fail()
	}

}
