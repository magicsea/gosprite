package gosprite

import "math"

type Vector struct {
	X float64
	Y float64
}

func NewVector(x,y float64) Vector {
	return Vector{X:x,Y:y}
}

func VectorZero()  Vector{
	return Vector{0,0}
}

func VectorOne()  Vector{
	return Vector{1,1}
}

//加法
func (v Vector) Add(val Vector)  Vector{
	return Vector{v.X+val.X,v.Y+val.Y}
}

//加法自身
func (v Vector) AddBy(val Vector)  {
	v.X = v.X+val.X
	v.Y = v.Y+val.Y
}

//减法
func (v Vector) Sub(val Vector)  Vector{
	return Vector{v.X-val.X,v.Y-val.Y}
}
//减法自身
func (v Vector) SubBy(val Vector)  {
	v.X = v.X-val.X
	v.Y = v.Y-val.Y
}

//乘法
func (v Vector) Mul(val float64)  Vector{
	return Vector{v.X*val,v.Y*val}
}

//乘法自身
func (v Vector) MulBy(val float64)  {
	v.X = v.X*val
	v.Y = v.Y*val
}


func (v Vector) Normal() Vector {
	val := v.X*v.X+v.Y*v.Y
	l := math.Sqrt(val)
	newv := Vector{
		v.X/l,
		v.Y/l,
	}
	return newv
}

