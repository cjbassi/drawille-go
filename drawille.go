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
	return []image.Point{}
}
