package point

import "math"

type Point struct 
{
    x float64
    y float64
}

func (p1 Point) Distance(p2 Point) float64 {
   return math.Sqrt(math.Pow(p1.x-p2.x,2)+math.Pow(p1.y-p2.y,2))
 
}