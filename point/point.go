package point

import "math"

type Point struct {
	X, Y float64;
}

func (p1 Point) Distance(p2 Point) float64 {
   return math.Sqrt(math.Pow(p1.X-p2.X,2)+math.Pow(p1.Y-p2.Y,2))
 
}