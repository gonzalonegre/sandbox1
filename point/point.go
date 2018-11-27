package point

import "math"
import "github.com/gonzalonegre/sandbox1/comparer"

type Point struct {
	X, Y float64;
}

func (p1 Point) Distance(p2 Point) float64 {
   return math.Sqrt(math.Pow(p1.X-p2.X,2)+math.Pow(p1.Y-p2.Y,2))
 
}

func (p1 Point) IsInLine(p2,p3 Point) bool {

  // The easy way to check is that they have the same tang.
  // y2-y1/x2-x1 = y3-y2/x3-x2
  //First, a boundary condition
  if utils.FloatsEquals(p1.X,p2.X) && utils.FloatsEquals(p2.X,p3.X) {
  	return true
  }
  return utils.FloatsEquals( (p2.Y-p1.Y)*(p3.X-p2.X) , (p3.Y-p2.Y)*(p2.X-p1.X)  ) 

}


func sign ( p1, p2, p3 Point) float64 {
    return (p1.X - p3.X) * (p2.Y - p3.Y) - (p2.X - p3.X) * (p1.Y - p3.Y)
}

func PointInTriangle ( pt, v1, v2, v3 Point) bool {
    var d1, d2, d3 float64; 
    var has_neg, has_pos bool 

    d1 = sign(pt, v1, v2);
    d2 = sign(pt, v2, v3);
    d3 = sign(pt, v3, v1);

    has_neg = (d1 < 0) || (d2 < 0) || (d3 < 0);
    has_pos = (d1 > 0) || (d2 > 0) || (d3 > 0);

    return !(has_neg && has_pos);
}


