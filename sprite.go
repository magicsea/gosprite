package gosprite

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
)

type SpriteType int
const (
	SpriteTypeSingle SpriteType = 0
	SpriteTypeSlice SpriteType = 1
)

type Sprite struct {
	BaseNode
	image *ebiten.Image
	spriteType SpriteType
}

func NewSprite(imageSrc []byte) (*Sprite,error){
	sp := new(Sprite)
	base := NewBaseNode(VectorZero(),0,sp)
	sp.BaseNode = base


	img, _, err := image.Decode(bytes.NewReader(imageSrc))
	if err != nil {
		fmt.Println("decode error:",err)
		return nil,err
	}

	newImage, errNew := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if errNew!=nil {
		fmt.Println("new image error:",errNew)
		return nil,errNew
	}

	sp.image = newImage

	w,h := newImage.Size()
	sp.size = NewVector(float64(w),float64(h))

	return sp,nil
}

func  (sp *Sprite) SetSpriteType(stp SpriteType)  {
	sp.spriteType = stp
}


func (sp *Sprite) Update(detaTime int64) {

}


func (sp *Sprite) Draw(target *ebiten.Image) {

	sp.opt.GeoM.Reset()
	sw,sh := sp.image.Size()
	if sp.spriteType==SpriteTypeSingle {
		sp.opt.GeoM.Scale(sp.scale.X*sp.size.X/float64(sw),sp.scale.Y*sp.size.Y/float64(sh))
		//sp.opt.GeoM.Translate(-float64(sp.size.X)*sp.anchor.X, -float64(sp.size.Y)*sp.anchor.Y)
		sp.opt.GeoM.Translate(-float64(sp.size.X)/2, -float64(sp.size.Y)/2)
		sp.opt.GeoM.Rotate(math.Pi * float64(sp.rotate) /180)
		sp.opt.GeoM.Translate(sp.position.X, sp.position.Y)
		target.DrawImage(sp.image, &sp.opt)
	} else if  sp.spriteType==SpriteTypeSlice {
		xc := int(sp.size.X/float64(sw))
		yc := int(sp.size.Y/float64(sh))
		//fmt.Println(xc,yc)
		for i:=0;i<=xc ;i++  {
			for j:=0;j<=yc ;j++  {
				sp.opt.GeoM.Reset()
				sp.opt.GeoM.Scale(sp.scale.X,sp.scale.Y)
				sp.opt.GeoM.Translate(sp.position.X+float64(sw)*float64(i)*sp.scale.X, sp.position.Y+float64(sh)*float64(j)*sp.scale.Y)
				target.DrawImage(sp.image, &sp.opt)
			}
		}

	}

	//fmt.Println("draw sprite,",sp.position,sp.size)
}

