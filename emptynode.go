
package gosprite

import (
	"github.com/hajimehoshi/ebiten"
)


type EmptyNode struct {
	BaseNode
}


func NewEmptyNode() *EmptyNode{
	sp := new(EmptyNode)
	base := NewBaseNode(VectorZero(),0,sp)
	sp.BaseNode = base
	return sp
}


func (sp *EmptyNode) Update(detaTime int64) {

}

func (sp *EmptyNode) Draw(target *ebiten.Image) {

}

