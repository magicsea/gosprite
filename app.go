package gosprite

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"time"
)


var (
	currScene IScene
	screenWidth,screenHeight int
	appTitle string
	lastTime int64
	detaTime float64
	logicTimeTester *TimeTester
	renderUpdateTimeTester *TimeTester
	drawTimeTester *TimeTester
)

//帧间隔(秒)
func GetDetaTime() float64 {
	return detaTime
}

func Start(scene IScene,width, height int, title string) error {
	if err := Goto(scene);err!=nil {
		fmt.Println(err)
		return err
	}
	logicTimeTester = NewTimeTester(60)
	drawTimeTester = NewTimeTester(60)
	renderUpdateTimeTester = NewTimeTester(60)

	screenWidth = width
	screenHeight = height
	appTitle = title
	ebiten.SetRunnableInBackground(true)

	ret := ebiten.Run(update,screenWidth,screenHeight,1,appTitle)
	return ret
}

func Goto(scene IScene) error {
	scene.(IDrawScene).initScene()
	if err := scene.Init();err!=nil {
		fmt.Println(err)
		return err
	}
	currScene = scene
	return nil
}



func update(screen *ebiten.Image) error {
	now := time.Now().UnixNano()
	ideta :=  now-lastTime
	detaTime = float64(ideta)/float64(time.Second)
	lastTime = now

	if currScene!=nil {
		//fmt.Println("draw frame")

		logicTimeTester.StartTest()
		currScene.Update(detaTime)
		logicTimeTester.EndTest()

		renderUpdateTimeTester.StartTest()
		ds := currScene.(IDrawScene)
		ds.update(ideta)
		renderUpdateTimeTester.EndTest()

	}


	if ebiten.IsDrawingSkipped() {
		return nil
	}

	//render scene
	if currScene!=nil {
		drawTimeTester.StartTest()
		ds := currScene.(IDrawScene)
		ds.draw(screen)
		drawTimeTester.EndTest()
	}
	//debug
	msg := fmt.Sprintf(`TPS: %0.2f FPS:%0.2f`, ebiten.CurrentTPS(),ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

	msgTime := fmt.Sprintf(`Logic:%0.2fms DrawUpdate:%0.2fms Draw:%0.2fms`, logicTimeTester.GetAvgTime(),renderUpdateTimeTester.GetAvgTime(),drawTimeTester.GetAvgTime())
	ebitenutil.DebugPrintAt(screen, msgTime,0,20)
	return nil
}


var idcreator int64
func genID() int64  {
	idcreator++
	return idcreator
}