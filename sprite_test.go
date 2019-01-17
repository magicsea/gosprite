package gosprite

import (
	"fmt"
	"testing"
)

func TestMath(t *testing.T) {
	v := Vector{-3,4}
	n := v.Normal()
	fmt.Println(n)
}