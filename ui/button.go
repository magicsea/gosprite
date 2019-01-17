package ui

import (
	. "github.com/magicsea/gosprite"
	"image"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/text"
)

type Button struct {
	BaseNode

	Text string

	mouseDown bool

	onPressed func(b *Button)

}

func NewButton(pos Vector,size Vector,text string) *Button {
	btn := &Button{
		Text: text,
	}
	base := NewBaseNode(pos,0,btn)
	base.SetSize(size)
	btn.BaseNode = base
	btn.SetColor(color.White)
	return btn
}

func (b *Button) Update(detaTime int64) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.GetIRect().Min.X <= x && x < b.GetIRect().Max.X && b.GetIRect().Min.Y <= y && y < b.GetIRect().Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}
func (node *Button) GetIRectScale(scale float64) image.Rectangle{
	rect := node.GetIRect()
	sub := rect.Max.Sub(rect.Min)
	w := int(float64(sub.X/2)*(scale-1))
	h := int(float64(sub.Y/2)*(scale-1))

	rect.Min.X=rect.Min.X-w
	rect.Max.X=rect.Max.X+w
	rect.Min.Y=rect.Min.Y-h
	rect.Max.Y=rect.Max.Y+h


	return rect
}

func (b *Button) Draw(dst *ebiten.Image) {
	if b.mouseDown {
		drawNinePatches(dst, b.GetIRectScale(1.1), imageSrcRects[imageTypeButtonPressed],b.GetColor())
	} else {
		drawNinePatches(dst, b.GetIRect(), imageSrcRects[imageTypeButton],b.GetColor())
	}


	bounds, _ := font.BoundString(uiFont, b.Text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := b.GetIRect().Min.X + (b.GetIRect().Dx()-w)/2
	y := b.GetIRect().Max.Y - (b.GetIRect().Dy()-uiFontMHeight)/2
	text.Draw(dst, b.Text, uiFont, x, y, color.Black)
}

func (b *Button) SetText(txt string) {
	b.Text = txt
}

func (b *Button) SetOnPressed(f func(b *Button)) {
	b.onPressed = f
}