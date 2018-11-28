package planet

import "math"


const (
	TOTAL_ANGLE = 360
	Pi float64 = 3.14159265358979323846264338327950288419716939937510582097494459 // https://oeis.org/A000796
)

type Planet struct {
	RadiusKm float64;
	InitialAngle int64;
	RotationSpeed int64; // Clockwise with negative values.
}

func (p1 Planet) AngularPos(day int) uint {
   var currentAngle = p1.InitialAngle + p1.RotationSpeed * int64(day)	   
   currentAngle %=TOTAL_ANGLE
   if(currentAngle<0){
   	   currentAngle += TOTAL_ANGLE; 
   }
   return uint(currentAngle)
}

func (p1 Planet) CurrentPos(day int) Point {
   var currentAngle = p1.AngularPos(day)
   var radianAngle = float64(currentAngle) * Pi *2 / TOTAL_ANGLE
   var x = p1.RadiusKm * math.Cos (radianAngle)
   var y = p1.RadiusKm * math.Sin (radianAngle)
   return Point{x , y }
 
}

