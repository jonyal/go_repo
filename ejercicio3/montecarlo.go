package montecarlo

import (
	"math/rand"
)

func MonteCarlo(points int) int {
	insideCircle := 0

	for i := 0; i < points; i++ {
		x := rand.Float64()
		y := rand.Float64()

		genVal := x*x + y*y
		if(genVal<=1)
		{
			insideCircle++;
		}	
	}

	return (4*(insideCircle/points))
}