package compare;

import (
	"testing"
)

func TestFloatsEquals(t *testing.T) {
	if (compare.FloatsEquals(.1,.2) ) {
		t.Fatalf("Expected FloatsEquals(.1,.2) ")
	}
	if !(compare.FloatsEquals(.1,.100000000000000001) ) {
		t.Fatalf("Expected FloatsEquals(.1,.100000000000000001) ")
	}
	if !(compare.FloatsEquals(.1,.1) ) {
		t.Fatalf("Expected FloatsEquals(.1,.1) ")
	}
}

