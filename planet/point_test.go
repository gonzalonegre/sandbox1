package planet_test;

import (
	"testing";
	"math";
	"github.com/gonzalonegre/sandbox1/planet";
);

func TestDistance(t *testing.T) {
        P1 := planet.Point{1,1}
        P2 := planet.Point{4,5}
	if (P1.Distance(P2) != 5) {
		t.Fatalf("Expected %v, got: %v", math.Sqrt(25), P1.Distance(P2))
	}
	if (P2.Distance(P1) != 5) {
		t.Fatalf("Expected %v, got: %v", math.Sqrt(25), P2.Distance(P2))
	}
}

func TestIsInLine(t *testing.T) {
	P1 := planet.Point{1,1}
        P2 := planet.Point{-1,-1}
        P3 := planet.Point{2.5,2.5}
        P4 := planet.Point{-1,-1.0000001}
	if !P1.IsInLine(P2,P3) {
		t.Fatalf("IsInLine(P1,P2,P3) failed.")
	}
	if  P1.IsInLine(P2,P4) {
		t.Fatalf("IsInLine(P1,P2,P4) failed.")
	}
}


func TestPointInsideTriangle(t *testing.T) {
	P1 := planet.Point{-1,1}
        P2 := planet.Point{-1,-1}
        P3 := planet.Point{2.5,0}
        P4 := planet.Point{0,0}
	P5 := planet.Point{-10,-10}
	if !planet.PointInsideTriangle(P4,P1,P2,P3) {
		t.Fatalf("PointInTriangle(P4,P1,P2,P3) failed.")
	}
	if planet.PointInsideTriangle(P5,P1,P4,P3) {
		t.Fatalf("PointInTriangle(P5,P1,P4,P3) failed.")
	}
}









