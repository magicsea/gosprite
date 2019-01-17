package ui

import (
	"errors"
	. "github.com/magicsea/gosprite"
	"image/color"
	"strings"
	"golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/text"
)

const (
	textBoxPaddingLeft = 8
)



type TextBox struct {
	BaseNode
	Text string

	contentBuf *ebiten.Image
	vScrollBar *VScrollBar
	offsetX    int
	offsetY    int


	fontSize int
	fontFace font.Face

	aliVertType AliVerticalType
	aliHeriType AliHorizontalType
}

func NewTextBox(pos Vector,size Vector,text string,aliVertType AliVerticalType,aliHeriType AliHorizontalType, hasScrollBar bool) *TextBox {
	tb := &TextBox{
		aliHeriType:aliHeriType,
		aliVertType:aliVertType,
		Text: text,
	}
	base := NewBaseNode(pos,0,tb)
	base.SetSize(size)
	tb.BaseNode = base
	if hasScrollBar {
		tb.vScrollBar = &VScrollBar{}
	}
	tb.fontFace = uiFont
	tb.SetColor(color.White)
	return tb
}

func (t *TextBox) AppendLine(line string) {
	if t.Text == "" {
		t.Text = line
	} else {
		t.Text += "\n" + line
	}
}

func (t *TextBox) Update(detaTime int64) {
	if  t.vScrollBar!=nil {
		t.vScrollBar.SetPosition(NewVector(float64(t.GetIRect().Max.X - VScrollBarWidth), float64(t.GetIRect().Min.Y)))
		t.vScrollBar.Height = t.GetIRect().Dy()

		_, h := t.contentSize()
		t.vScrollBar.Update(h)

		t.offsetX = 0
		t.offsetY = t.vScrollBar.ContentOffset()
	}
}

func (t *TextBox) contentSize() (int, int) {
	h := len(strings.Split(t.Text, "\n")) * lineHeight
	return t.GetIRect().Dx(), h
}

func (t *TextBox) viewSize() (int, int) {
	barSize := 0
	if t.vScrollBar!=nil {
		barSize = VScrollBarWidth + textBoxPaddingLeft
	}
	return t.GetIRect().Dx() - barSize, t.GetIRect().Dy()
}

func (t *TextBox) contentOffset() (int, int) {
	return t.offsetX, t.offsetY
}

func (t *TextBox) Draw(dst *ebiten.Image) {
	drawNinePatches(dst, t.GetIRect(), imageSrcRects[imageTypeTextBox],t.GetColor())

	if t.contentBuf != nil {
		vw, vh := t.viewSize()
		w, h := t.contentBuf.Size()
		if vw > w || vh > h {
			t.contentBuf.Dispose()
			t.contentBuf = nil
		}
	}
	if t.contentBuf == nil {
		w, h := t.viewSize()
		t.contentBuf, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}

	t.contentBuf.Clear()
	var th int
	var fontH = uiFontMHeight
	cx,ch := t.contentBuf.Size()
	var offyTop =  lineHeight - (lineHeight-fontH)/2
	for i, line := range strings.Split(t.Text, "\n") {
		th += lineHeight
		x := -t.offsetX + textBoxPaddingLeft
		b,_ := font.BoundString(t.fontFace,line)
		w := (b.Max.X - b.Min.X).Ceil()
		if t.aliHeriType==AliHorizontal_Right {
			x = -t.offsetX + cx - textBoxPaddingLeft - w
		} else if t.aliHeriType==AliHorizontal_Mid {
			x = -t.offsetX + (cx - w)/2
		}
		y :=  -t.offsetY + i*lineHeight + offyTop
		if y < -lineHeight {
			continue
		}
		if _, h := t.viewSize(); y >= h+lineHeight {
			continue
		}
		text.Draw(t.contentBuf, line, t.fontFace, x, y, color.Black)
	}
	op := &ebiten.DrawImageOptions{}


	if  t.vScrollBar==nil {
		//有scroll不支持竖排列
		var _,offy float64
		if t.aliVertType==AliVertical_Mid {
			offy = float64(ch-th)/2-float64(lineHeight-fontH)/2
		} else if t.aliVertType==AliVertical_Bottom {
			offy = float64(ch-(lineHeight-fontH)/2-th)
		}
		op.GeoM.Translate(0, offy)
	}


	op.GeoM.Translate(float64(t.GetIRect().Min.X), float64(t.GetIRect().Min.Y))

	dst.DrawImage(t.contentBuf, op)

	if  t.vScrollBar!=nil {
		t.vScrollBar.Draw(dst,color.White)
	}
}

func (t *TextBox) SetText(txt string) {
	t.Text = txt
}


func (t *TextBox) SetFont(fontName string,fontSize int) error {
	t.fontSize = fontSize
	face := GetFontFace(fontName,t.fontSize)
	if face==nil {
		return errors.New("not found")
	}
	t.fontFace = face
	return nil
}