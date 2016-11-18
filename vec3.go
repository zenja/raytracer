package raytracer

import "math"

type Vec3 struct {
	E0 float64
	E1 float64
	E2 float64
}

func (v Vec3) X() float64 {
	return v.E0
}

func (v Vec3) Y() float64 {
	return v.E1
}

func (v Vec3) Z() float64 {
	return v.E2
}

func (v Vec3) R() float64 {
	return v.E0
}

func (v Vec3) G() float64 {
	return v.E1
}

func (v Vec3) B() float64 {
	return v.E2
}

func (v *Vec3) AddEqual(other Vec3) {
	v.E0 += other.E0
	v.E1 += other.E1
	v.E2 += other.E2
}

func (v *Vec3) MinusEqual(other Vec3) {
	v.E0 -= other.E0
	v.E1 -= other.E1
	v.E2 -= other.E2
}

func (v *Vec3) MultiplyEqual(other Vec3) {
	v.E0 *= other.E0
	v.E1 *= other.E1
	v.E2 *= other.E2
}

func (v *Vec3) DivideEqual(other Vec3) {
	v.E0 /= other.E0
	v.E1 /= other.E1
	v.E2 /= other.E2
}

func (v *Vec3) MultiplyValueEqual(value float64) {
	v.E0 *= value
	v.E1 *= value
	v.E2 *= value
}

func (v *Vec3) DivideValueEqual(value float64) {
	v.E0 /= value
	v.E1 /= value
	v.E2 /= value
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{v.E0 + other.E0, v.E1 + other.E1, v.E2 + other.E2}
}

func (v Vec3) Minus(other Vec3) Vec3 {
	return Vec3{v.E0 - other.E0, v.E1 - other.E1, v.E2 - other.E2}
}

func (v Vec3) Multiply(other Vec3) Vec3 {
	return Vec3{v.E0 * other.E0, v.E1 * other.E1, v.E2 * other.E2}
}

func (v Vec3) Divide(other Vec3) Vec3 {
	return Vec3{v.E0 / other.E0, v.E1 / other.E1, v.E2 / other.E2}
}

func (v Vec3) AddValue(value float64) Vec3 {
	return Vec3{v.E0 + value, v.E1 + value, v.E2 + value}
}

func (v Vec3) MinusValue(value float64) Vec3 {
	return Vec3{v.E0 - value, v.E1 - value, v.E2 - value}
}

func (v Vec3) MultiplyValue(value float64) Vec3 {
	return Vec3{v.E0 * value, v.E1 * value, v.E2 * value}
}

func (v Vec3) DivideValue(value float64) Vec3 {
	return Vec3{v.E0 / value, v.E1 / value, v.E2 / value}
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.E0*v.E0 + v.E1*v.E1 + v.E2*v.E2)
}

func (v Vec3) SquaredLength() float64 {
	return v.E0*v.E0 + v.E1*v.E1 + v.E2*v.E2
}

func (v Vec3) Dot(other Vec3) float64 {
	return v.E0*other.E0 + v.E1*other.E1 + v.E2*other.E2
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivideValue(v.Length())
}
