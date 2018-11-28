package planet_test
import "github.com/gonzalonegre/sandbox1/planet"

import "testing"


func TestAngularPos(t *testing.T) {
	P1 := planet.Planet{10, planet.TOTAL_ANGLE/4, -1}
	if P1.AngularPos(planet.TOTAL_ANGLE/4)!=0 {
		t.Fatalf("P1.AngularPos(planet.TOTAL_ANGLE/4)!=0 failed.")
	}
	if  P1.AngularPos(planet.TOTAL_ANGLE/2)!=planet.TOTAL_ANGLE/4*3 {
		t.Fatalf("P1.AngularPos(TOTAL_ANGLE/2)!=TOTAL_ANGLE/4*3 failed.")
	}
}

func CurrentPos(t *testing.T) {
	P1 := planet.Planet{10, planet.TOTAL_ANGLE/4, planet.TOTAL_ANGLE/4}
	point := P1.CurrentPos(3) 
	if  point.X !=21 {
	t.Fatalf("P1.AngularPos(planet.TOTAL_ANGLE/4)!=0 failed.")
	}
}
