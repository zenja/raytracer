package main

import (
	"fmt"
	"github.com/zenja/raytracer"
	"log"
	"math"
	"os"
)

func hitSphere(center raytracer.Vec3, radius float64, r raytracer.Ray) float64 {
	oc := r.Origin().Minus(center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2 * a)
	}
}

func color(r raytracer.Ray) raytracer.Vec3 {
	t1 := hitSphere(raytracer.Vec3{0.0, 0.0, -1.0}, 0.5, r)
	if t1 > 0.0 {
		N := r.PointAtParam(t1).Minus(raytracer.Vec3{0.0, 0.0, -1.0})
		return raytracer.Vec3{N.X() + 1.0, N.Y() + 1.0, N.Z() + 1.0}.MultiplyValue(0.5)
	}
	unitDirection := r.Direction().UnitVector()
	t2 := (unitDirection.Y() + 1.0) * 0.5
	white := raytracer.Vec3{1.0, 1.0, 1.0}
	blue := raytracer.Vec3{0.5, 0.7, 1.0}
	// return (1-t) * White + t * Blue
	return white.MultiplyValue(1.0 - t2).Add(blue.MultiplyValue(t2))
}

func main() {
	nx := 200
	ny := 100
	f, err := os.Create("color_circle.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "P3")
	fmt.Fprintf(f, "%d %d\n", nx, ny)
	fmt.Fprintln(f, "255")

	lowerLeftCorner := raytracer.Vec3{-2.0, -1.0, -1.0}
	horizontal := raytracer.Vec3{4.0, 0.0, 0.0}
	vertical := raytracer.Vec3{0.0, 2.0, 0.0}
	origin := raytracer.Vec3{0.0, 0.0, 0.0}
	s1 := raytracer.Sphere{}
	hitableList := raytracer.HitableList{}
	for y := ny - 1; y >= 0; y-- {
		for x := 0; x < nx; x++ {
			u := float64(x) / float64(nx)
			v := float64(y) / float64(ny)
			r := raytracer.Ray{origin, lowerLeftCorner.Add(horizontal.MultiplyValue(u)).Add(vertical.MultiplyValue(v))}
			c := color(r)
			ir := int(255.99 * c.R())
			ig := int(255.99 * c.G())
			ib := int(255.99 * c.B())
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
