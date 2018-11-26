package point_test;

import (
	"fmt";
	"testing";
	"math";
	"github.com/gonzalonegre/sandbox1/point";
	"github.com/google/go-cmp/cmp";
);

func TestDistance(t *testing.T) {
        var P1 point.Point
        P1 = point.Point{X:1,1}
        P2 := point.Point{4,5}
	if (P1.Distance(P2) != 5) {
		t.Fatalf("Expected %v, got: %v", math.Sqrt(25), P1.Distance(P2))
	}
}

