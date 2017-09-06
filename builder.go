/**
 * Created by 93201 on 2017/9/5.
 */
package geo

import "errors"

type Builder struct {
	vertexes   []*Point
	sides      []*Line
	bb         *BoundingBox
	firstPoint bool
	isClosed   bool
}

func NewBuilder() *Builder {
	return &Builder{
		firstPoint: true,
	}
}

func (b *Builder) AddVertex(p *Point) *Builder {
	if b.isClosed {
		b.vertexes = []*Point{}
		b.isClosed = false
	}
	b.updateBoundingBox(p)
	b.vertexes = append(b.vertexes, p)
	if len(b.vertexes) > 1 {
		b.sides = append(b.sides, NewLine(b.vertexes[len(b.vertexes)-2], p))
	}
	return b
}

func (b *Builder) Close() *Builder {
	b.validate()
	b.sides = append(b.sides, NewLine(b.vertexes[len(b.vertexes)-1], b.vertexes[0]))
	b.isClosed = true
	return b
}

func (b *Builder) updateBoundingBox(p *Point) {
	if b.firstPoint {
		b.bb = &BoundingBox{
			xMax: p.X,
			xMin: p.X,
			yMax: p.Y,
			yMin: p.Y,
		}
		b.firstPoint = false
	} else {
		if p.X > b.bb.xMax {
			b.bb.xMax = p.X
		} else if p.X < b.bb.xMin {
			b.bb.xMax = p.X
		}

		if p.Y > b.bb.yMax {
			b.bb.yMax = p.Y
		} else if p.Y < b.bb.yMin {
			b.bb.yMin = p.Y
		}
	}
}

func (b *Builder) validate() {
	if len(b.vertexes) < 3 {
		panic(errors.New("polygon must have at least 3 points"))
	}
}

func (b *Builder) Build() *Polygon {
	b.validate()

	if !b.isClosed {
		b.sides = append(b.sides, NewLine(b.vertexes[len(b.vertexes)-1], b.vertexes[0]))
	}
	return NewPolygon(b.sides, b.bb)
}
