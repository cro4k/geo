/**
 * Created by 93201 on 2017/7/12.
 */
package geo

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}
