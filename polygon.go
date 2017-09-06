/**
 * Created by 93201 on 2017/9/5.
 */
package geo

type Polygon struct {
	bb    *BoundingBox
	sides []*Line
}

func NewPolygon(sides []*Line, bb *BoundingBox) *Polygon {
	return &Polygon{
		sides: sides,
		bb:    bb,
	}
}

func (p *Polygon) Contains(point *Point) bool {
	if p.inBoundingBox(point) {
		ray := p.createRay(point)
		intersection := 0
		for _, side := range p.sides {
			if p.Intersect(ray, side) {
				intersection++
			}
		}
		if intersection%2 != 0 {
			return true
		}
	}
	return false
}

func (p *Polygon) Intersect(ray, side *Line) bool {
	intersectPoint := &Point{}
	if !ray.vertical && !side.vertical {
		if ray.a-side.a == 0 {
			return false
		}

		x := (side.b - ray.b) / (ray.a - side.a) //x = (b2-b1)/(a1-a2)
		y := side.a*x + side.b                   // y = a2*x+b2
		intersectPoint.X = x
		intersectPoint.Y = y
	} else if ray.vertical && !side.vertical {
		x := ray.start.X
		y := side.a*x + side.b
		intersectPoint.X = x
		intersectPoint.Y = y
	} else if !ray.vertical && side.vertical {
		x := side.start.X
		y := ray.a*x + ray.b
		intersectPoint.X = x
		intersectPoint.Y = y
	} else {
		return false
	}

	return side.IsInside(intersectPoint) && ray.IsInside(intersectPoint)
}

func (p *Polygon) createRay(point *Point) *Line {
	epsilon := (p.bb.xMax - p.bb.xMin) / 10e6
	outsidePoint := NewPoint(p.bb.xMin-epsilon, p.bb.yMin)
	return NewLine(outsidePoint, point)
}

func (p *Polygon) inBoundingBox(point *Point) bool {
	return !(point.X < p.bb.xMin ||
		point.X > p.bb.xMax ||
		point.Y < p.bb.yMin ||
		point.Y > p.bb.yMax)
}
