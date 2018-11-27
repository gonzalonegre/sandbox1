package utils_test;

import (
	"testing"
	"github.com/gonzalonegre/sandbox1/utils"
)

func TestFloatsEquals(t *testing.T) {
	if (utils.FloatsEquals(.1,.2) ) {
		t.Fatalf("Expected FloatsEquals(.1,.2) ")
	}
	if !(utils.FloatsEquals(.1,.100000000000000001) ) {
		t.Fatalf("Expected FloatsEquals(.1,.100000000000000001) ")
	}
	if !(utils.FloatsEquals(.1,.1) ) {
		t.Fatalf("Expected FloatsEquals(.1,.1) ")
	}
}

