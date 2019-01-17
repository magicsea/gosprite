package gosprite

import "time"

type TimeTester struct {
	testCount int
	lastTimeList []float64
	lastTime int64

	avgTimer float64
}

func NewTimeTester(testCount int)  *TimeTester{
	return &TimeTester{testCount:testCount}
}

func (t *TimeTester) StartTest()  {
	t.lastTime = time.Now().UnixNano()
}


func (t *TimeTester) EndTest() float64 {
	avlOffSet := float64(time.Now().UnixNano()-t.lastTime)/float64(time.Millisecond)

	t.lastTimeList = append(t.lastTimeList,avlOffSet)
	if len(t.lastTimeList)>int(t.testCount) {
		t.lastTimeList = append([]float64{},t.lastTimeList[1:]...)
	}
	var sum float64
	for _, v := range t.lastTimeList {
		sum += v
	}
	t.avgTimer = sum/float64(len(t.lastTimeList))
	return t.avgTimer
}

func  (t *TimeTester) GetAvgTime() float64  {
	return t.avgTimer
}