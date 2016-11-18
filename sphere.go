package raytracer

import "github.com/engoengine/math"

type Sphere struct {
	Center Vec3
	Radius float64
}

func (s *Sphere) hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := r.Origin().Minus(s.Center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c
	var rec HitRecord
	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / (2 * a)
		if tMin < temp && temp < tMax {
			rec.T = temp
			rec.P = r.PointAtParam(rec.T)
			rec.Normal = rec.P.Minus(s.Center).DivideValue(s.Radius)
			return true, rec
		}
		temp = (-b + math.Sqrt(discriminant)) / (2 * a)
		if tMin < temp && temp < tMax {
			rec.T = temp
			rec.P = r.PointAtParam(rec.T)
			rec.Normal = rec.P.Minus(s.Center).DivideValue(s.Radius)
			return true, rec
		}
	}
	return false, rec
}
