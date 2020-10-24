package shapes

import (
	"fmt"
	"sort"
	"strings"

	"github.com/golang/geo/r2"
)

const (
	rectInterInfo = "Between rectangle %s at (%+v,%+v), w=%d, h=%d."
)

// RectangleIntersction represents the intersectin of 2 or more rectangles
type RectIntersection struct {
	// IDXs specify indexes for all the rectangles intersecting
	IDXs []int
	// Rect is the resulting rectangle being created after interection of rectanles specified int the index
	Rect r2.Rect
}

func (ri RectIntersection) BuildInfoText() string {
	var indexes []string
	for _, v := range ri.IDXs {
		indexes = append(indexes, fmt.Sprint(v))
	}
	rects := strings.Join(indexes, ",")
	rawRect := BuildRawRect(ri.Rect)
	return fmt.Sprintf(rectInterInfo, rects, rawRect.X, rawRect.Y, rawRect.W, rawRect.H)
}

// GenerateIntersections generates all the possbile intersections of a rectangle
func GenerateIntersections(rects []r2.Rect) {
	var i = 1
	// first Level rectangles without intersection
	items := BuildInitialIntersections(rects)
	var result []RectIntersection
	dupMap := make(map[string]bool)

	for {
		intersections := Intersect(items)
		if intersections == nil {
			break
		}
		var filter []RectIntersection
		for _, isx := range intersections {
			uniquID := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(isx.IDXs)), ","), "[]")
			if !dupMap[uniquID] {
				dupMap[uniquID] = true
				result = append(result, isx)
				filter = append(filter, isx)
				fmt.Printf("%d: %s\n", i, isx.BuildInfoText())
				i++
			}
		}
		if len(intersections) == 1 {
			break
		}
		items = filter
	}

}

// BuildInitialIntersections builds initial intersections for set of rectangles
func BuildInitialIntersections(items []r2.Rect) (result []RectIntersection) {
	for k, v := range items {
		result = append(result, RectIntersection{IDXs: []int{k}, Rect: v})
	}
	return
}

// Intersect checks for intersections in between given array of rectangles
func Intersect(items []RectIntersection) []RectIntersection {
	var isx []RectIntersection
	idx := 0
	dupMap := make(map[string]bool)
	for idx < (len(items) - 1) {
		first := items[idx]
		idSec := RectIntersection{}
		for i := idx + 1; i < len(items); i++ {
			second := items[i]
			result := first.Rect.Intersection(second.Rect)
			if !result.IsEmpty() {
				ids := normaliseSlices(first.IDXs, second.IDXs)
				uniquID := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]")
				if dupMap[uniquID] {
					continue
				} else {
					dupMap[uniquID] = true
				}
				idSec.Rect = result
				idSec.IDXs = ids
				isx = append(isx, idSec)
			}
		}
		idx++
	}
	return isx
}

// normaliseSlices reuturns a uniqe and sorted slice after joining 2 different slices
func normaliseSlices(a, b []int) []int {
	exists := make(map[int]bool)
	for _, v := range a {
		exists[v] = true
	}
	for _, v := range b {
		if !exists[v] {
			a = append(a, v)
		}
	}
	sort.Ints(a)
	return a
}
