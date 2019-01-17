package gosprite

import (

	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"image/color"
	"math"
	"golang.org/x/image/font"
)

var (
	emptyImage *ebiten.Image
)
func init() {
	emptyImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	_ = emptyImage.Fill(color.White)

	LoadFont("default",fonts.ArcadeN_ttf)
}

func colorScale(clr color.Color) (rf, gf, bf, af float64) {
	r, g, b, a := clr.RGBA()
	if a == 0 {
		return 0, 0, 0, 0
	}

	rf = float64(r) / float64(a)
	gf = float64(g) / float64(a)
	bf = float64(b) / float64(a)
	af = float64(a) / 0xffff
	return
}


func genVertices(pos Vector,r float64,num int,color color.Color) []ebiten.Vertex {
	var (
		centerX = float32(pos.X)
		centerY = float32(pos.Y)
		//r      float64 = 160
	)
	cr,cg,cb,ca := color.RGBA()
	_ = ca
	vs := []ebiten.Vertex{}
	for i := 0; i < num; i++ {
		rate := float64(i) / float64(num)

		//fmt.Println(color.RGBA())

		vs = append(vs, ebiten.Vertex{
			DstX:   float32(r*math.Cos(2*math.Pi*rate)) + centerX,
			DstY:   float32(r*math.Sin(2*math.Pi*rate)) + centerY,
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(cr)/65535.0,
			ColorG: float32(cg)/65535.0,
			ColorB: float32(cb)/65535.0,
			ColorA: float32(ca)/65535.0,
		})
	}

	vs = append(vs, ebiten.Vertex{
		DstX:   centerX,
		DstY:   centerY,
		SrcX:   0,
		SrcY:   0,
		ColorR: float32(cr)/65535.0,
		ColorG: float32(cg)/65535.0,
		ColorB: float32(cb)/65535.0,
		ColorA: float32(ca)/65535.0,
	})

	return vs
}



func LoadFont(name string,fontSrc []byte) (*truetype.Font,error) {

	if f,b := fontMap[name];b {
		return f,nil
	}
	tt, err := truetype.Parse(fontSrc)
	if err != nil {
		return nil,err
	}

	fontMap[name]=tt
	return tt,nil
}

func LoadFontFace(name string,ft *truetype.Font,size int) (font.Face,error){
	if size < FontBaseSize {
		size = FontBaseSize
	}
	newName := fmt.Sprintf("%s_%v",name,size)
	if f,b:=fontfaceMap[newName];b {
		return f,nil
	}
	const dpi = 72
	ff := truetype.NewFace(ft, &truetype.Options{
		Size:    float64(size),
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	return ff,nil
}

func GetFontFace(fontname string,size int) font.Face {
	ft,b := fontMap[fontname]
	if !b {
		return nil
	}
	face,errlf := LoadFontFace(fontname,ft,size)
	if errlf!=nil {
		fmt.Println(errlf)
		return nil
	}
	return face
}

var (
	fontMap = map[string]*truetype.Font{}
	fontfaceMap = map[string]font.Face{}

)

const (
	FontBaseSize = 8
)
//
//func getArcadeFonts(scale int) font.Face {
//	if arcadeFonts == nil {
//		tt, err := truetype.Parse(fonts.ArcadeN_ttf)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		arcadeFonts = map[int]font.Face{}
//		for i := 1; i <= 4; i++ {
//			const dpi = 72
//			arcadeFonts[i] = truetype.NewFace(tt, &truetype.Options{
//				Size:    float64(arcadeFontBaseSize * i),
//				DPI:     dpi,
//				Hinting: font.HintingFull,
//			})
//		}
//	}
//	return arcadeFonts[scale]
//}
