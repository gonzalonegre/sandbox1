package comparer


import (
	"math"
)


const minimalDifference = 0.00000001;

func FloatsEquals(f1,f2 float64) bool {
	if f1==f2 {
	  return true
	}
	if math.Abs(f1-f2)< minimalDifference {
	  return true
	}
	if f1+f2 == 0 {
	   return false
	}
	
	if math.Abs( ((f1-f2)*2) / (f1+f1)) < minimalDifference {
	  return true
	}
	return false
}
