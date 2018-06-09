package vec_test

import (
	"testing"

	"github.com/pietroglyph/go-vector/2d/64bit/vec"
)

func TestNew(t *testing.T) {
	var x = 1.1
	var y = 2.2

	v := vec.Vector2D{X: x, Y: y}
	if v.X != x || v.Y != y {
		t.Fail()
	}
}

func TestFromScalar(t *testing.T) {
	var s = 1.1

	v := vec.FromScalar(s)
	if v.X != s || v.Y != s {
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	v := vec.Zero()
	if v.X != 0 || v.Y != 0 {
		t.Fail()
	}
}

func TestUnit(t *testing.T) {
	v := vec.Unit()
	if v.X != 1 || v.Y != 1 {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	var x = 1.1
	var y = 2.2

	v1 := vec.Vector2D{X: x, Y: y}
	v2 := v1.Copy()
	if v2.X != x || v2.Y != y {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	var x1 = 1.0
	var y1 = 2.0
	var x2 = 4.0
	var y2 = 8.0

	v1 := vec.Vector2D{X: x1, Y: y1}
	v2 := vec.Vector2D{X: x2, Y: y2}
	result := v1.Add(v2)
	if result.X != 5 || result.Y != 10 {
		t.Fail()
	}
}

func TestSubtract(t *testing.T) {
	var x1 = 1.0
	var y1 = 2.0
	var x2 = 4.0
	var y2 = 8.0

	v1 := vec.Vector2D{X: x1, Y: y1}
	v2 := vec.Vector2D{X: x2, Y: y2}
	result := v1.Subtract(v2)
	if result.X != -3 || result.Y != -6 {
		t.Fail()
	}
}

func TestMultiply(t *testing.T) {
	var x1 = 1.0
	var y1 = 2.0
	var x2 = 4.0
	var y2 = 8.0

	v1 := vec.Vector2D{X: x1, Y: y1}
	v2 := vec.Vector2D{X: x2, Y: y2}
	result := v1.Multiply(v2)
	if result.X != 4 || result.Y != 16 {
		t.Fail()
	}
}

func TestDivide(t *testing.T) {
	var x1 = 1.0
	var y1 = 2.0
	var x2 = 4.0
	var y2 = 10.0

	v1 := vec.Vector2D{X: x1, Y: y1}
	v2 := vec.Vector2D{X: x2, Y: y2}
	result := v1.Divide(v2)
	if result.X != 0.25 || result.Y != 0.2 {
		t.Fail()
	}
}

func TestMultiplyScalar(t *testing.T) {
	var x = 1.0
	var y = 2.0

	v := vec.Vector2D{X: x, Y: y}
	result := v.MultiplyScalar(1.1)
	if result.X != 1.1 || result.Y != 2.2 {
		t.Fail()
	}
}

func TestDivideScalar(t *testing.T) {
	var x = 1.0
	var y = 2.0

	v := vec.Vector2D{X: x, Y: y}
	result := v.DivideScalar(2)
	if result.X != 0.5 || result.Y != 1 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var x = 1.0
	var y = 2.0

	v := vec.Vector2D{X: x, Y: y}
	result := v.String()
	if result != "1:2" {
		t.Fail()
	}
}
