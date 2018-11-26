package point_test

import (
	"fmt"
	"testing"
	"math"
	"github.com/gonzalonegre/sandbox1/point"
	"github.com/google/go-cmp/cmp"
)

func TestDistance(t *testing.T) {
        P1 Point(1,1)
        P2 Point(4,5) 
	if !cmp.Equal(P1.Distance(P2), math.Sqrt(25), Float64Comparer) {
		t.Fatalf("Expected %v, got: %v", math.Sqrt(25), P1.Distance(P2))
	}
}
