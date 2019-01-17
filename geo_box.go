package gosprite

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"math"
)


type Box struct {
	BaseNode
}


func NewBox(rotate float64,size Vector,color color.Color) *Box{
	sp := new(Box)
	base := NewBaseNode(VectorZero(),rotate,sp)
	base.SetSize(size)
	sp.BaseNode = base
	sp.SetColor(color)
	return sp
}


func (sp *Box) Update(detaTime int64) {

}

func (sp *Box) Draw(target *ebiten.Image) {

	//sp.opt.GeoM.Reset()
	//sp.opt.GeoM.Translate(-float64(sp.size.X)/2, -float64(sp.size.Y)/2)
	//sp.opt.GeoM.Rotate(math.Pi * float64(sp.rotate) /180)
	//sp.opt.GeoM.Translate(sp.size.X/2+sp.position.X, sp.size.Y/2+sp.position.Y)
	//target.DrawImage(sp.image, &sp.opt)
	////fmt.Println("draw sprite,",sp.position,sp.size)
	//fmt.Println("draw sprite,",sp.rotate)
	ew, eh := emptyImage.Size()

	sp.opt.GeoM.Reset()
	sp.opt.GeoM.Scale(sp.size.X/float64(ew), sp.size.Y/float64(eh))
	sp.opt.GeoM.Translate(-sp.size.X*sp.anchor.X, -sp.size.Y*sp.anchor.Y)
	sp.opt.GeoM.Rotate(math.Pi * float64(sp.rotate) /180)
	sp.opt.GeoM.Translate(sp.position.X, sp.position.Y)
	sp.opt.ColorM.Reset()
	sp.opt.ColorM.Scale(colorScale(sp.color))
	// Filter must be 'nearest' filter (default).
	// Linear filtering would make edges blurred.
	_ = target.DrawImage(emptyImage, &sp.opt)
}
