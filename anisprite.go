package gosprite

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"image"
	"math"
)

//帧动画对象
type AniSprite struct {
	BaseNode
	image *ebiten.Image
	frameAniInfo map[string]*FrameAniData

	currFrameTimer int
	currFrame int
	currAni *FrameAniData
	endFun func(string)
}
//一个动作数据
type FrameAniData struct {
	AniName string
	FrameInterval int
	FrameRects []image.Rectangle
}
func NewAniSprite(imageSrc []byte,frameAniInfo map[string]*FrameAniData) (*AniSprite,error){
	sp := new(AniSprite)
	base := NewBaseNode(VectorZero(),0,sp)
	sp.BaseNode = base
	sp.frameAniInfo = frameAniInfo

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

func  (sp *AniSprite) Play(aniName string) error {
	ani,b:= sp.frameAniInfo[aniName]
	if !b {
		return errors.New("not found")
	}
	sp.endFun = nil
	sp.currAni = ani
	sp.currFrameTimer = 0
	sp.currFrame = 0
	return nil
}

func  (sp *AniSprite) PlayThen(aniName string,endFun func(ani string)) error {
	if err := sp.Play(aniName);err!=nil {
		return err
	}
	sp.endFun = endFun
	return nil
}

func (sp *AniSprite) Update(detaTime int64) {
	if sp.currAni==nil {
		return
	}
	sp.currFrameTimer++
	if sp.currFrameTimer>sp.currAni.FrameInterval {
		sp.currFrame++
		sp.currFrameTimer = 0
	}
	//fmt.Println("play:",sp.currFrame,sp.currAni.AniName)
	if sp.currFrame>=len(sp.currAni.FrameRects) {
		sp.currFrame = 0
		if sp.endFun!=nil {
			fun := sp.endFun
			sp.endFun = nil

			fun(sp.currAni.AniName)

		}
	}


}


func (sp *AniSprite) Draw(target *ebiten.Image) {
	if sp.currAni==nil {
		return
	}

	rect := sp.currAni.FrameRects[sp.currFrame]
	w,h := float64(rect.Size().X),float64(rect.Size().Y)
	sp.opt.GeoM.Reset()

	sp.opt.GeoM.Scale(sp.scale.X,sp.scale.Y)
	sp.opt.GeoM.Translate(-w/2, -h/2)
	sp.opt.GeoM.Rotate(math.Pi * float64(sp.rotate) /180)

	sp.opt.GeoM.Translate(sp.position.X, sp.position.Y)

	target.DrawImage(sp.image.SubImage(rect).(*ebiten.Image), &sp.opt)
	//fmt.Println("draw sprite,",sp.position,sp.size)
}

