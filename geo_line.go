package gosprite

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)


type Line struct {
	BaseNode
}


func NewLine(from,to Vector,width float64,color color.Color) *Line{
	sp := new(Line)
	var rotate float64 = 180*math.Atan2(to.Y-from.Y, to.X-from.X)/math.Pi
	base := NewBaseNode(from,rotate,sp)
	length := math.Hypot(to.X-from.X, to.Y-from.Y)
	sp.BaseNode = base
	sp.color = color

	sp.SetSize(NewVector(length,width))
	return sp
}
func (sp *Line) Update(detaTime int64) {

}


func (sp *Line) Draw(target *ebiten.Image) {

	ew, eh := emptyImage.Size()

	sp.opt.GeoM.Reset()
	sp.opt.GeoM.Scale(sp.size.X/float64(ew), sp.size.Y/float64(eh))
	sp.opt.GeoM.Rotate(math.Pi * float64(sp.rotate) /180)
	sp.opt.GeoM.Translate(sp.position.X, sp.position.Y)
	sp.opt.ColorM.Scale(colorScale(sp.color))
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	_ = target.DrawImage(emptyImage, &sp.opt)

}
