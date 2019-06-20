package ellipse

import (
	"fmt"
	"math"
)

type Init struct {
	A, B float64
}

func (e* Init) GetEccentricity() float64 {
	fmt.Println("")
	return (math.Sqrt(math.Pow(e.A, 2) - math.Pow(e.B, 2))) / e.A
}
