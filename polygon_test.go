/**
 * Created by 93201 on 2017/9/6.
 */
package geo

import "testing"

func TestPointInPolygon(t *testing.T) {
	// convex polygon
	convexPolygon := NewBuilder().
		AddVertex(NewPoint(0, 0)).
		AddVertex(NewPoint(4, 1)).
		AddVertex(NewPoint(3, 5)).
		AddVertex(NewPoint(1, 4)).
		Build()
	if convexPolygon.Contains(NewPoint(6, 6)) { //expect false
		t.Fail()
	}

	if !convexPolygon.Contains(NewPoint(1, 1)) { // expect true
		t.Fail()
	}

	//non-convex polygon
	polygon := NewBuilder().
		AddVertex(NewPoint(0, 0)).
		AddVertex(NewPoint(2, 2)).
		AddVertex(NewPoint(4, 1)).
		AddVertex(NewPoint(5, 5)).
		AddVertex(NewPoint(1, 6)).
		Build()

	if polygon.Contains(NewPoint(2, 1)) { //expect false
		t.Fail()
	}

	if !polygon.Contains(NewPoint(1, 2)) { // expect true
		t.Fail()
	}
}
