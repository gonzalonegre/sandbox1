package point_test;

import (
	"testing";
	"math";
	"github.com/gonzalonegre/sandbox1/point";
);

func TestDistance(t *testing.T) {
        P1 := point.Point{1,1}
        P2 := point.Point{4,5}
	if (P1.Distance(P2) != 5) {
		t.Fatalf("Expected %v, got: %v", math.Sqrt(25), P1.Distance(P2))
	}
}

