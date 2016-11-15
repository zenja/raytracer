package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	nx := 200
	ny := 100
	f, err := os.Create("image.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "P3")
	fmt.Fprintf(f, "%d %d\n", nx, ny)
	fmt.Fprintln(f, "255")
	for y := ny - 1; y >= 0; y-- {
		for x := 0; x < nx; x++ {
			r := float64(x) / float64(nx)
			g := float64(y) / float64(ny)
			b := 0.2
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
		}
	}
}
