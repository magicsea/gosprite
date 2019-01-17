package gosprite

import (
	"github.com/hajimehoshi/ebiten"
	"sort"
)


//golang的结构和类型应该分开~
type TreeNode struct {
	uuid int64
	parent *TreeNode
	children map[int64]*TreeNode


	renderNode INode
}

func NewTreeNode(rn INode) *TreeNode{
	tn := &TreeNode{}
	tn.children = make(map[int64]*TreeNode)

	tn.renderNode = rn
	tn.uuid = genID()
	return tn
}


func  (node *TreeNode) SetParent(parent *TreeNode) {

	if node.parent!=nil {
		node.DetachParent()
	}
	node.parent = parent
	node.parent.children[node.uuid]=node
}

func  (node *TreeNode) DetachParent() {
	delete(node.parent.children,node.uuid)
	node.parent = nil
}

func (node *TreeNode) Update(detaTime int64)  {
	if node.renderNode!=nil {
		//fmt.Println("draw one")
		node.renderNode.Update(detaTime)
	}

	//fmt.Println("find child ",len(node.children))

	//update
	for _, v := range node.children {
		//fmt.Println("pre draw child:",reflect.TypeOf(v.renderNode),v.renderNode)
		v.Update(detaTime)
	}
}

func (node *TreeNode) Draw(target *ebiten.Image)  {
	if node.renderNode!=nil {
		//fmt.Println("draw one")
		node.renderNode.Draw(target)
	}

	//fmt.Println("find child ",len(node.children))
	//sort
	var list []*TreeNode
	for _, v := range node.children {
		list = append(list,v)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].renderNode.GetDepth()==list[j].renderNode.GetDepth() {
			return list[i].renderNode.GetUUID()<list[j].renderNode.GetUUID()
		}
		return list[i].renderNode.GetDepth()<list[j].renderNode.GetDepth()
	})
	//draw
	for _, v := range list {
		//fmt.Println("pre draw child:",reflect.TypeOf(v.renderNode),v.renderNode)
		v.Draw(target)
	}
}

type IScene interface {
	Init() error
	Update(detaTime float64)
}

type IDrawScene interface {
	initScene()
	draw(target *ebiten.Image)
	update(detatTime int64)
}

type Scene struct {
	renderRoot *TreeNode
	renderRootUI *TreeNode
}

func NewScene() *Scene {
	s :=  &Scene{}
	s.initScene()
	return s
}

func (s *Scene) initScene()  {
	renderNode := NewEmptyNode()
	s.renderRoot = renderNode.treeNode

	renderNodeUI := NewEmptyNode()
	s.renderRootUI = renderNodeUI.treeNode
}

func (s *Scene) update(detaTime int64)  {
	s.renderRoot.Update(detaTime)
	s.renderRootUI.Update(detaTime)
}

func (s *Scene) draw(target *ebiten.Image) {
	s.renderRoot.Draw(target)
	s.renderRootUI.Draw(target)
}

func (s *Scene)  GetRenderRoot() INode{
	return s.renderRoot.renderNode
}


func (s *Scene)  GetRenderUIRoot() INode{
	return s.renderRootUI.renderNode
}

func  (s *Scene) AddNode(node INode)  {
	node.SetParent(s.GetRenderRoot())
}


func  (s *Scene) AddUINode(node INode)  {
	node.SetParent(s.GetRenderUIRoot())
}
