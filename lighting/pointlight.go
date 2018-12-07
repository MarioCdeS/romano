package lighting

import (
	"github.com/MarioCdeS/romano/gfx"
	"github.com/MarioCdeS/romano/linalg"
)

type PointLight struct {
	Position  linalg.Point
	Intensity gfx.Color
}
