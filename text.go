package gosprite

import (
	"errors"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"

	"image/color"
)


type Text struct {
	BaseNode
	fontSize int
	text string
	fontFace font.Face
}


func NewText(text string,fontsize int,color color.Color) *Text{
	sp := new(Text)
	base := NewBaseNode(VectorZero(),0,sp)
	sp.BaseNode = base
	sp.color = color
	sp.fontSize = fontsize
	sp.text = text
	sp.SetFont("default")
	return sp
}



func (sp *Text) Draw(target *ebiten.Image) {
	text.Draw(target, sp.text, sp.fontFace, int(sp.position.X), int(sp.position.Y), sp.color)
}

func (sp *Text) SetText(txt  string) {
	sp.text = txt
}

func (sp *Text) SetFont(fontName string) error {
	face := GetFontFace(fontName,sp.fontSize)
	if face==nil {
		return errors.New("not found")
	}
	sp.fontFace = face
	return nil
}
