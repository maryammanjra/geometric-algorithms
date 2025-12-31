package convexhull

import (
	"geometric-algorithms/geometry"
	"reflect"
	"testing"
)

func pointInSlice(p geometry.Point, pts []geometry.Point) bool {
	for _, q := range pts {
		if reflect.DeepEqual(p, q) {
			return true
		}
	}
	return false
}

func TestGrahamsScanStatic(t *testing.T) {
	tests := []struct {
		name        string
		points      []geometry.Point
		expectedLen int
	}{
		{
			name: "Triangle",
			points: []geometry.Point{
				{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1},
			},
			expectedLen: 3,
		},
		{
			name: "Square",
			points: []geometry.Point{
				{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}, {X: 1, Y: 1},
			},
			expectedLen: 4,
		},
		{
			name: "Pentagon",
			points: []geometry.Point{
				{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 1},
			},
			expectedLen: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hull := grahamsScan(tt.points)
			if len(hull) != tt.expectedLen {
				t.Errorf("expected hull length %d, got %d", tt.expectedLen, len(hull))
			}

			for _, hp := range hull {
				if !pointInSlice(hp, tt.points) {
					t.Errorf("hull point %v not in original points", hp)
				}
			}
		})
	}
}
func TestGrahamsScanConvexity(t *testing.T) {
	points := []geometry.Point{
		{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 0}, {X: 1, Y: -1},
	}

	hull := grahamsScan(points)
	if len(hull) < 3 {
		t.Fatalf("Hull must have at least 3 points, got %d", len(hull))
	}

	n := len(hull)
	sign := 0
	for i := 0; i < n; i++ {
		a := hull[i]
		b := hull[(i+1)%n]
		c := hull[(i+2)%n]

		cross := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
		if cross != 0 {
			s := 1
			if cross < 0 {
				s = -1
			}
			if sign == 0 {
				sign = s
			} else if sign != s {
				t.Errorf("Hull is not convex at points %v, %v, %v", a, b, c)
			}
		}
	}
}
