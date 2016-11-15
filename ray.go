package raytracer

type Ray struct {
	A Vec3
	B Vec3
}

func (r Ray) Origin() Vec3 {
	return r.A
}

func (r Ray) Direction() Vec3 {
	return r.B
}

func (r Ray) PointAtParam(t float64) Vec3 {
	// return A + t*B
	return r.A.Add(r.B.MultiplyValue(t))
}
