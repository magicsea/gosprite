package gosprite

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"

)

const CircleTriNum = 32
type Circle struct {
	BaseNode
	radius float64

	vertices []ebiten.Vertex
	indices []uint16
}


func NewCircle(radius float64,color color.Color) *Circle{
	sp := new(Circle)
	base := NewBaseNode(VectorZero(),0,sp)
	sp.BaseNode = base
	sp.radius = radius
	sp.SetColor(color)

	return sp
}


func (sp *Circle) Update(detaTime int64) {

	sp.vertices = genVertices(sp.position,sp.radius,CircleTriNum,sp.color)
	sp.indices = []uint16{}
	for i := 0; i < CircleTriNum; i++ {
		sp.indices = append(sp.indices, uint16(i), uint16(i+1)%uint16(CircleTriNum), uint16(CircleTriNum))
	}
}

func (sp *Circle) Draw(target *ebiten.Image) {
	op := &ebiten.DrawTrianglesOptions{}
	target.DrawTriangles(sp.vertices, sp.indices, emptyImage,op)

}
