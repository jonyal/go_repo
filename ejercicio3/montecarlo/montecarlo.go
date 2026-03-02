package montecarlo

import (
	"math/rand"
)

func MonteCarlo(points int) float64 {
	insideCircle := 0

	for i := 0; i < points; i++ {
		x := rand.Float64()
		y := rand.Float64()
		
		genVal := x*x + y*y
		if(genVal<=1) {
			insideCircle++;
		}	
	}

	result := float64(insideCircle)/float64(points)

	return (4.0*result)
}