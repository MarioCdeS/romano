package lighting

import (
	"math"

	"github.com/MarioCdeS/romano/geometry"
	"github.com/MarioCdeS/romano/gfx"
	"github.com/MarioCdeS/romano/linalg"
)

func At(p *linalg.Point, normal, eye *linalg.Vector, light *PointLight, material *geometry.Material) *gfx.Color {
	effectCol := material.Color.Hadamard(&light.Intensity)
	res := effectCol.Scale(material.Ambient())

	lightVec := light.Position.SubPoint(p).MutNormalized()
	lightDotNorm := lightVec.Dot(normal)

	if lightDotNorm >= 0 {
		res.MutAdd(effectCol.Scale(material.Diffuse() * lightDotNorm))

		reflectDotEye := lightVec.MutNeg().Reflect(normal).Dot(eye)

		if reflectDotEye > 0 {
			res.MutAdd(light.Intensity.Scale(material.Specular() * math.Pow(reflectDotEye, material.Shininess())))
		}
	}

	res.A = 1

	return res
}
