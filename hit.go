package raytracer

type HitRecord struct {
	T      float64
	P      Vec3
	Normal Vec3
}

type Hitable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord)
}

type HitableList struct {
	Hitables []Hitable
}

func (l *HitableList) hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	var tempRec HitRecord
	var hitAnything bool = false
	var closestSoFar float64 = tMax
	var resultRec HitRecord
	for _, h := range l.Hitables {
		if isHit, rec := h.Hit(r, tMax, closestSoFar); isHit {
			hitAnything = true
			closestSoFar = rec.T
			resultRec = tempRec
		}
	}
	return hitAnything, resultRec
}
