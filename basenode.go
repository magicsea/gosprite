package gosprite

import (
	"github.com/hajimehoshi/ebiten"
	"image"
	"image/color"
)

type INode interface {
	GetUUID() int64
	GetPosition() Vector
	GetRotato() float64

	SetParent(parent INode)
	DetachParent()
	GetTreeNode() *TreeNode
	SetTreeNode(*TreeNode)
	//GetParent()INode
	SetDepth(v int)
	GetDepth() int

	SetSize(v Vector)
	GetSize() Vector

	GetScale() Vector
	SetScale(v Vector)

	RecountPos()

	Update(detaTime int64)
	Draw(target *ebiten.Image)
}

type IDrawable interface {
	OnDrawFrame(target *ebiten.Image)
}

type BaseNode struct {
	uuid int64
	position Vector
	localPosition Vector
	rotate float64
	size Vector
	scale Vector
	color color.Color

	treeNode *TreeNode
	anchor Vector
	depth int
	opt    ebiten.DrawImageOptions
}

func NewBaseNode(pos Vector,rotate float64,nodeIns INode) BaseNode{
	node := BaseNode{
		uuid:genID(),
		position:pos,
		rotate:rotate,
		anchor:NewVector(0.5,0.5),
		scale:NewVector(1,1),
		treeNode:NewTreeNode(nodeIns),
	}
	return node
}

func  (node *BaseNode) GetUUID() int64  {
	return node.uuid
}

func (node *BaseNode) SetDepth(v int)  {
	node.depth = v
}
func (node *BaseNode) GetDepth() int  {
	return node.depth
}

func (node *BaseNode) GetAnchor() Vector {
	return node.anchor
}

func (node *BaseNode) SetAnchor(v Vector){
	node.anchor = v
}

func (node *BaseNode) SetColor(c color.Color){
	node.color = c
}
func (node *BaseNode) GetColor() color.Color{
	return node.color
}

func (node *BaseNode) GetPosition() Vector {
	return node.position
}

func (node *BaseNode) SetPosition(v Vector){
	node.position = v
	if node.GetParent()!=nil {
		node.localPosition = node.position.Sub(node.GetParent().renderNode.GetPosition())
	} else {
		node.localPosition = v
	}


	for _, c := range node.treeNode.children {
		c.renderNode.RecountPos()
	}
}

func (node *BaseNode) RecountPos()  {
	if node.GetParent()!=nil {
		node.position = node.localPosition.Add(node.GetParent().renderNode.GetPosition())
	} else {
		node.position = node.localPosition
	}

	for _, c := range node.treeNode.children {
		c.renderNode.RecountPos()
	}
}

func (node *BaseNode) GetLocalPosition() Vector {
	return node.localPosition
}

func (node *BaseNode) SetLocalPosition(v Vector){
	node.localPosition = v
	node.RecountPos()
}


func (node *BaseNode) GetScale() Vector {
	return node.scale
}

func (node *BaseNode) SetScale(v Vector){
	node.scale = v
}


func (node *BaseNode) SetSize(v Vector){
	node.size = v
}


func (node *BaseNode) GetSize() Vector{
	return node.size
}

func (node *BaseNode) GetIRect() image.Rectangle{
	return image.Rect(int(node.position.X),int(node.position.Y),
		int(node.position.X+node.size.X),int(node.position.Y+node.size.Y))
}

func (node *BaseNode) GetRotato() float64  {
	return node.rotate
}


func (node *BaseNode) SetRotato(r float64)  {
	node.rotate = r
}
func  (node *BaseNode) GetTreeNode() *TreeNode {
	return node.treeNode
}
func  (node *BaseNode) SetTreeNode(tn *TreeNode) {
	node.treeNode = tn
}

func  (node *BaseNode) SetParent(parent INode) {
	node.treeNode.SetParent(parent.GetTreeNode())
	pos := node.position
	node.SetPosition(pos)
}

func SetParent(child INode,parent INode)  {
	if child.GetTreeNode()==nil {
		child.SetTreeNode(NewTreeNode(child))
	}
	child.GetTreeNode().SetParent(parent.GetTreeNode())
}


func  (node *BaseNode) GetParent() *TreeNode {
	return node.treeNode.parent
}

func  (node *BaseNode) DetachParent() {
	node.treeNode.DetachParent()
	node.treeNode = nil
}

func (node *BaseNode) Update(detaTime int64) {
	//do nothing
}

func (node *BaseNode) Draw(target *ebiten.Image) {
	//do nothing
}

func  (node *BaseNode) Destory()  {
	node.DetachParent()
}