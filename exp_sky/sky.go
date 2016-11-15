package main

import (
	"fmt"
	"github.com/zenja/raytracer"
	"log"
	"os"
)

func color(r *raytracer.Ray) raytracer.Vec3 {
	unitDirection := r.Direction().UnitVector()
	t := (unitDirection.Y() + 1.0) * 0.5
	white := raytracer.Vec3{1.0, 1.0, 1.0}
	blue := raytracer.Vec3{0.5, 0.7, 1.0}
	// return (1-t) * White + t * Blue
	return white.MultiplyValue(1.0 - t).Add(blue.MultiplyValue(t))
}

func main() {
	nx := 200
	ny := 100
	f, err := os.Create("sky.ppm")
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
	for y := ny - 1; y >= 0; y-- {
		for x := 0; x < nx; x++ {
			u := float64(x) / float64(nx)
			v := float64(y) / float64(ny)
			r := raytracer.Ray{origin, lowerLeftCorner.Add(horizontal.MultiplyValue(u)).Add(vertical.MultiplyValue(v))}
			c := color(&r)
			ir := int(255.99 * c.R())
			ig := int(255.99 * c.G())
			ib := int(255.99 * c.B())
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
