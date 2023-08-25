package ui

import (
	. "github.com/magicsea/gosprite"
	"image"
	"image/color"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/text"
)

const (
	checkboxWidth       = 16
	checkboxHeight      = 16
	checkboxPaddingLeft = 8
)

type CheckBox struct {
	BaseNode

	Text string

	checked   bool
	mouseDown bool

	onCheckChanged func(c *CheckBox)
}

func NewCheckBox(pos Vector, size Vector, text string) *CheckBox {
	ck := &CheckBox{
		Text: text,
	}
	base := NewBaseNode(pos, 0, ck)
	base.SetSize(size)
	ck.BaseNode = base
	ck.SetColor(color.White)
	return ck
}

func (c *CheckBox) width() int {
	b, _ := font.BoundString(uiFont, c.Text)
	w := (b.Max.X - b.Min.X).Ceil()
	return checkboxWidth + checkboxPaddingLeft + w
}

func (c *CheckBox) Update(detaTime int64) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if int(c.GetPosition().X) <= x && x < int(c.GetPosition().X)+c.width() && int(c.GetPosition().Y) <= y && y < int(c.GetPosition().Y)+checkboxHeight {
			c.mouseDown = true
		} else {
			c.mouseDown = false
		}
	} else {
		if c.mouseDown {
			c.checked = !c.checked
			if c.onCheckChanged != nil {
				c.onCheckChanged(c)
			}
		}
		c.mouseDown = false
	}
}

func (c *CheckBox) Draw(dst *ebiten.Image) {
	t := imageTypeCheckBox
	if c.mouseDown {
		t = imageTypeCheckBoxPressed
	}
	r := image.Rect(int(c.GetPosition().X), int(c.GetPosition().Y), int(c.GetPosition().X)+checkboxWidth, int(c.GetPosition().Y)+checkboxHeight)
	drawNinePatches(dst, r, imageSrcRects[t], c.GetColor())
	if c.checked {
		drawNinePatches(dst, r, imageSrcRects[imageTypeCheckBoxMark], c.GetColor())
	}

	x := int(c.GetPosition().X) + checkboxWidth + checkboxPaddingLeft
	y := int(c.GetPosition().Y+16) - (16-uiFontMHeight)/2
	text.Draw(dst, c.Text, uiFont, x, y, color.Black)
}

func (c *CheckBox) Checked() bool {
	return c.checked
}

func (c *CheckBox) SetChecked(b bool) {
	c.checked = b
}

func (c *CheckBox) SetOnCheckChanged(f func(c *CheckBox)) {
	c.onCheckChanged = f
}

func (b *CheckBox) SetText(txt string) {
	b.Text = txt
}
