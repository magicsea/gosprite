package ui

import (
	"errors"
	"github.com/hajimehoshi/ebiten/inpututil"
	. "github.com/magicsea/gosprite"
	"golang.org/x/image/font"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/text"
)

const (
	inputFieldPaddingLeft = 4
)

type InputField struct {
	BaseNode
	//Text string
	ValueText string
	counter int
	contentBuf *ebiten.Image
	vScrollBar *VScrollBar
	offsetX    int
	offsetY    int
	active bool

	fontSize int
	fontFace font.Face
	aliVertType AliVerticalType
	aliHeriType AliHorizontalType
	mouseDown bool
}

func NewInputField(pos Vector,size Vector,text string,aliVertType AliVerticalType,aliHeriType AliHorizontalType, hasScrollBar bool) *InputField {
	tb := &InputField{
		ValueText: text,
		aliHeriType:aliHeriType,
		aliVertType:aliVertType,
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

func (t *InputField) AppendLine(line string) {
	if t.ValueText == "" {
		t.ValueText = line
	} else {
		t.ValueText += "\n" + line
	}
}

// repeatingKeyPressed return true when key is pressed considering the repeat state.
func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func  (t *InputField) updateInput()  {
	// Add a string from InputChars, that returns string input by users.
	// Note that InputChars result changes every frame, so you need to call this
	// every frame.
	r := ebiten.InputChars()

	t.ValueText += string(r)

	// Adjust the string to be at most 10 lines.
	ss := strings.Split(t.ValueText, "\n")
	if len(ss) > 10 {
		t.ValueText = strings.Join(ss[len(ss)-10:], "\n")
	}

	// If the enter key is pressed, add a line break.
	if repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyKPEnter) {
		t.ValueText += "\n"
	}

	// If the backspace key is pressed, remove one character.
	if repeatingKeyPressed(ebiten.KeyBackspace) {
		rs := []rune(t.ValueText)
		if len(rs) >= 1 {
			rs = rs[:len(rs)-1]
		}
		t.ValueText = string(rs)
	}

}

func (t *InputField) updateActive()  {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if t.GetIRect().Min.X <= x && x < t.GetIRect().Max.X && t.GetIRect().Min.Y <= y && y < t.GetIRect().Max.Y {
			t.active = true
		} else {
			t.active = false
		}
	}
}

func (t *InputField) Update(detaTime int64) {

	t.updateActive()

	if t.active {
		t.updateInput()
	}

	if  t.vScrollBar!=nil {
		t.vScrollBar.SetPosition(NewVector(float64(t.GetIRect().Max.X - VScrollBarWidth), float64(t.GetIRect().Min.Y)))
		t.vScrollBar.Height = t.GetIRect().Dy()

		_, h := t.contentSize()
		t.vScrollBar.Update(h)

		t.offsetX = 0
		t.offsetY = t.vScrollBar.ContentOffset()
	}
}

func (t *InputField) contentSize() (int, int) {
	h := len(strings.Split(t.ValueText, "\n")) * lineHeight
	return t.GetIRect().Dx(), h
}

func (t *InputField) viewSize() (int, int) {
	barSize := 0
	if t.vScrollBar!=nil {
		barSize = VScrollBarWidth + inputFieldPaddingLeft
	}
	return t.GetIRect().Dx() - barSize, t.GetIRect().Dy()
}

func (t *InputField) contentOffset() (int, int) {
	return t.offsetX, t.offsetY
}

func (t *InputField) Draw(dst *ebiten.Image) {
	// Blink the cursor.
	var txt = t.ValueText
	if t.active {
		t.counter++
		if t.counter%60 < 30 {
			txt += "|"
		}else {
			txt += " "
		}
	}

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
	cx,ch := t.contentBuf.Size()
	var th int
	var fontH = uiFontMHeight
	var offyTop =  lineHeight - (lineHeight-fontH)/2
	for i, line := range strings.Split(txt, "\n") {
		x := -t.offsetX + inputFieldPaddingLeft
		b,_ := font.BoundString(t.fontFace,line)
		w := (b.Max.X - b.Min.X).Ceil()
		if t.aliHeriType==AliHorizontal_Right {
			x = -t.offsetX + cx - inputFieldPaddingLeft - w
		} else if t.aliHeriType==AliHorizontal_Mid {
			x = -t.offsetX + (cx - w)/2
		}

		th += lineHeight
		y := -t.offsetY + i*lineHeight + offyTop
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

	if t.vScrollBar!=nil {
		t.vScrollBar.Draw(dst,color.White)
	}
}

func (t *InputField) SetText(txt string) {
	t.ValueText = txt
}

func (t *InputField) GetText() string {
	return t.ValueText
}



func (t *InputField) SetFont(fontName string,fontSize int) error {
	t.fontSize = fontSize
	face := GetFontFace(fontName,t.fontSize)
	if face==nil {
		return errors.New("not found")
	}
	t.fontFace = face
	return nil
}