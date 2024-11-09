package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu_buffer.pprof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	var size int
	fmt.Scanf("%d", &size)

	points := make([][]float64, size+2, size+2)
	points[0], points[size+1] = []float64{0, 0}, []float64{0, 0}

	input := bufio.NewReader(os.Stdin)
	for i := 1; i < size+1; i++ {
		points[i] = []float64{0, 0}
		fmt.Fscanf(input, "%f %f\n", &points[i][0], &points[i][1])
	}

	fmt.Printf("%.20f\n", traveling(points))
}

func traveling(p [][]float64) float64 {
	var cost float64

	for i := 0; i < len(p)-1; i++ {
		x, y := p[i][0]-p[i+1][0], p[i][1]-p[i+1][1]
		cost += math.Sqrt(x*x + y*y)
	}

	return cost
}
