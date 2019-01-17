package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/examples/resources/images/blocks"
	g "github.com/magicsea/gosprite"
	"github.com/magicsea/gosprite/ui"
	"image"
	"image/color"
	"io/ioutil"
)


const (
	screenW = 800
	screenH = 600

)

type testscene1 struct {
	g.Scene
	gsp *g.Sprite
}

func (s *testscene1) Init() error {
	//bgfile,errload := ioutil.ReadFile("xw.jpg")
	//if errload!=nil {
	//	fmt.Println(errload)
	//	return errload
	//}

	//load bg
	//bg,bgErr := g.NewSprite(g.NewVector(screenW/2, screenH/2), 0, bgfile)
	//if bgErr != nil {
	//	fmt.Println("bg load error:", bgErr)
	//	return bgErr
	//}
	//bg.SetSize(g.NewVector(screenW,screenH))
	//s.AddNode(bg)

	bg2,bgErr2 := g.NewSprite(blocks.Background_png)
	if bgErr2 != nil {
		fmt.Println("bg load error:", bgErr2)
		return bgErr2
	}
	bg2.SetSize(g.NewVector(screenW,screenH))
	bg2.SetSpriteType(g.SpriteTypeSlice)
	bg2.SetScale(g.NewVector(50.0/32.0,50.0/32.0))
	s.AddNode(bg2)


	//load box
	box := g.NewBox(45,g.NewVector(100, 50),color.RGBA{0,0,255,128})
	box.SetPosition(g.NewVector(400, 300))
	box.SetDepth(1)
	box.SetAnchor(g.NewVector(0.5,0))
	s.AddNode(box)

	//load line
	cell := 50
	for i:=0;i<screenW/cell;i++  {
		from := g.NewVector(float64(i*cell),0)
		to := g.NewVector(float64(i*cell),screenH)
		line := g.NewLine(from,to,1,color.Black)
		line.SetDepth(2)
		s.AddNode(line)
	}
	for j:=0;j<screenH/cell ;j++  {
		from := g.NewVector(0,float64(j*cell))
		to := g.NewVector(screenW,float64(j*cell))
		line := g.NewLine(from,to,1,color.Black)
		line.SetDepth(2)
		s.AddNode(line)
	}

	//text
	fontbt,errlaodfd := ioutil.ReadFile("C:\\Windows\\Fonts\\msyh.ttc")
	if errlaodfd != nil {
		fmt.Println("loadfont:",errlaodfd)
		return errlaodfd
	}
	g.LoadFont("msyh",fontbt)

	txt := g.NewText("hello world你好",48,color.RGBA{255,0,0,255})
	txt.SetPosition(g.NewVector(200, 200))
	errfont := txt.SetFont("msyh")
	if errfont!=nil {
		fmt.Println("setfont:",errfont)
		return errfont
	}
	txt.SetDepth(3)
	s.AddNode(txt)


	ci := g.NewCircle(40,color.RGBA{0,128,0,128})
	ci.SetPosition(g.NewVector(100, 100))
	ci.SetDepth(4)
	s.AddNode(ci)

	//load role
	sp, spErr := g.NewSprite( images.Ebiten_png)
	sp.SetPosition(g.NewVector(400, 300))
	if spErr != nil {
		fmt.Println("sprite load error:", spErr)
		return spErr
	}
	sp.SetDepth(10)
	s.AddNode(sp)
	s.gsp = sp

	//load ani
	frameAniInfo := map[string]*g.FrameAniData{
		"idle":&g.FrameAniData{
			FrameInterval:10,
			AniName:"idle",
			FrameRects:[]image.Rectangle{
				image.Rect(0,0,32,32),
				image.Rect(1*32,0,2*32,32),
				image.Rect(2*32,0,3*32,32),
				image.Rect(3*32,0,4*32,32),
			},
		},
		"run":&g.FrameAniData{
			FrameInterval:10,
			AniName:"run",
			FrameRects:[]image.Rectangle{
				image.Rect(0,32,32,64),
				image.Rect(1*32,32,2*32,64),
				image.Rect(2*32,32,3*32,64),
				image.Rect(3*32,32,4*32,64),
				image.Rect(4*32,32,5*32,64),
				image.Rect(5*32,32,6*32,64),
				image.Rect(6*32,32,7*32,64),
				image.Rect(7*32,32,8*32,64),
			},
		},
		"jump":&g.FrameAniData{
			FrameInterval:10,
			AniName:"jump",
			FrameRects:[]image.Rectangle{
				image.Rect(0,64,32,96),
				image.Rect(1*32,64,2*32,96),
				image.Rect(2*32,64,3*32,96),
				image.Rect(3*32,64,4*32,96),
			},
		},
	}

	as,_ := g.NewAniSprite(images.Runner_png,frameAniInfo)
	as.SetPosition(g.NewVector(400,300))
	as.SetScale(g.NewVector(3,3))
	as.SetDepth(20)
	as.SetParent(sp)
	//s.AddNode(as)
	as.PlayThen("idle", func(ani string) {
		as.PlayThen("run", func(ani string) {
			as.Play("jump")
		})
	})

	//gui
	inf := ui.NewInputField(g.NewVector(screenW-200,100),g.NewVector(100,40),"",ui.AliVertical_Mid,ui.AliHorizontal_Right,false)
	errinf := inf.SetFont("msyh",16)
	if errinf!=nil {
		fmt.Println(errinf)
		return errinf
	}
	s.AddUINode(inf)
	inf2 := ui.NewInputField(g.NewVector(screenW-200,150),g.NewVector(100,40),"",ui.AliVertical_Mid,ui.AliHorizontal_Right,false)
	s.AddUINode(inf2)

	tb := ui.NewTextBox(g.NewVector(screenW-100,100),g.NewVector(100,100),"zzzz",ui.AliVertical_Mid,ui.AliHorizontal_Right,false)
	s.AddUINode(tb)

	btn := ui.NewButton(g.NewVector(screenW-100,0),g.NewVector(100,40),"changescene")
	btn.SetOnPressed(func(b *ui.Button) {
		fmt.Println("click btn")
		g.Goto(new(testscene2))
	})
	s.AddUINode(btn)

	ck := ui.NewCheckBox(g.NewVector(screenW-100,50),g.NewVector(100,40),"click me")
	ck.SetOnCheckChanged(func(b *ui.CheckBox) {
		fmt.Println("click ck",ck.Checked())
		tb.AppendLine("add:"+inf.ValueText)
	})
	s.AddUINode(ck)

	return nil
}

func (s *testscene1) Update(detaTime float64) {
	//左上是0,0
	var charX float64 = s.gsp.GetPosition().X
	var charY float64 = s.gsp.GetPosition().Y
	//if ebiten.IsKeyPressed(ebiten.KeyEnter) {
	//	fmt.Println("enter")
	//	g.Goto(new(testscene2))
	//}

	//if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	//	x,y := ebiten.CursorPosition()
	//	fmt.Println("mouse press:",x,y)
	//}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		charX -= 3
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		charX += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		charY -= 3
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		charY += 3
	}
	s.gsp.SetPosition(g.NewVector(charX,charY))
	s.gsp.SetRotato(s.gsp.GetRotato() + 1)
}

type testscene2 struct {
	g.Scene
}
func (s *testscene2) Init() error {
	txt := g.NewText("hello world你好",48,color.RGBA{255,255,0,255})
	txt.SetPosition(g.NewVector(200, 200))
	txt.SetDepth(3)
	s.AddNode(txt)
	return nil
}
func (s *testscene2) Update(detaTime float64) {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("enter")
		g.Goto(new(testscene1))
	}
}

func main() {
	fmt.Println("start")
	err := g.Start(new(testscene1),800, 600, "mygame")
	if err != nil {
		fmt.Println("run error:", err)
	}
}
