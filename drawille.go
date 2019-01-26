package drawille

import (
	"image"
)

const BRAILLE_OFFSET = '\u2800'

var BRAILLE = [4][2]rune{
	{'\u0001', '\u0008'},
	{'\u0002', '\u0010'},
	{'\u0004', '\u0020'},
	{'\u0040', '\u0080'},
}

type Color int

type Cell struct {
	Rune  rune
	Color Color
}

type Canvas struct {
	CellMap map[image.Point]Cell
}

func NewCanvas() *Canvas {
	return &Canvas{
		CellMap: make(map[image.Point]Cell),
	}
}

func (self *Canvas) Set(p image.Point, color Color) {
	point := image.Pt(p.X/2, p.Y/4)
	self.CellMap[point] = Cell{
		self.CellMap[point].Rune | BRAILLE[p.X%4][p.Y%2],
		color,
	}
}

func Line(p0, p1 image.Point) []image.Point {
	points := []image.Point{}

	leftPoint, rightPoint := p0, p1
	if leftPoint.X > rightPoint.X {
		leftPoint, rightPoint = rightPoint, leftPoint
	}

	xDistance := absInt(leftPoint.X - rightPoint.X)
	yDistance := absInt(leftPoint.Y - rightPoint.Y)
	slope := float64(yDistance) / float64(xDistance)
	slopeDirection := 1
	if rightPoint.Y < leftPoint.Y {
		slopeDirection = -1
	}

	targetYCoordinate := float64(leftPoint.Y)
	currentYCoordinate := leftPoint.Y
	for i := leftPoint.X; i < rightPoint.X; i++ {
		targetYCoordinate += (slope * float64(slopeDirection))
		if currentYCoordinate == int(targetYCoordinate) {
			points = append(points, image.Pt(i, currentYCoordinate))
		}
		for currentYCoordinate != int(targetYCoordinate) {
			points = append(points, image.Pt(i, currentYCoordinate))
			currentYCoordinate += slopeDirection
		}
	}

	return points
}

func absInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
