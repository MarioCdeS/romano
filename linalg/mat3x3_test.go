package linalg

import (
	"testing"

	"github.com/MarioCdeS/romano/float"
)

func TestMat3x3_SubMat(t *testing.T) {
	m := NewMat3x3(1, 2, 3, 4, 5, 6, 7, 8, 9)
	exp := NewMat2x2(1, 3, 7, 9)
	got := m.SubMat(1, 1)

	if !exp.Equal(got) {
		t.Error(expectedGotErrorString(exp, got))
	}
}

func TestMat3x3_Det(t *testing.T) {
	m := NewMat3x3(1, 2, 6, -5, 8, -4, 2, 6, 4)
	got := m.Det()

	if !float.ApproxEqual(-196, got) {
		t.Errorf("expected -196, got %g", got)
	}
}
