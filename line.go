/**
 * Created by 93201 on 2017/9/5.
 */
package geo

type Line struct {
	start    *Point
	end      *Point
	vertical bool
	a, b     float64
}

func NewLine(start, end *Point) *Line {
	l := &Line{}
	l.start = start
	l.end = end
	l.vertical = false
	if l.end.X-l.start.X != 0 {
		l.a = float64((l.end.Y - l.start.Y) / (l.end.X - l.start.X))
		l.b = float64(l.start.Y - l.a*l.start.X)
	} else {
		l.vertical = true
	}
	return l
}

func (l *Line) IsInside(point *Point) bool {
	maxX := maxFloat(l.start.X, l.end.X)
	minX := minFloat(l.start.X, l.end.X)
	maxY := maxFloat(l.start.Y, l.end.Y)
	minY := minFloat(l.start.Y, l.end.Y)

	return (point.X >= minX && point.X <= maxX) &&
		(point.Y >= minY && point.Y <= maxY)
}
