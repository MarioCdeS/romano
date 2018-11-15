package linalg

import (
	"testing"
)

func TestEqualf64(t *testing.T) {
	if !Equalf64(0.1, 0.4-0.3) {
		t.Fail()
	}
}
