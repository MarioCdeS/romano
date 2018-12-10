package tracer

import "math"

type PointLight struct {
	Position  Point
	Intensity Color
}

func At(p *Point, normal, eye *Vector, light *PointLight, material *Material) *Color {
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
