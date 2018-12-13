package tracer

import "math"

type PointLight struct {
	Position  Point
	Intensity Color
}

func ColorAt(world *World, ray *Ray) *Color {
	hit, ok := FindHit(ray, world.Intersections(ray))

	if !ok {
		return &Color{0, 0, 0, 1}
	}

	var res Color

	for _, light := range world.Lights {
		material := hit.obj.Material()
		effectCol := material.Color.Hadamard(&light.Intensity)

		res.MutAdd(effectCol.Scale(material.Ambient()))

		lightVec := light.Position.SubPoint(hit.Point).MutNormalized()

		if lightDotNorm := lightVec.Dot(hit.Normal); lightDotNorm >= 0 {
			res.MutAdd(effectCol.Scale(material.Diffuse() * lightDotNorm))

			if reflectDotCam := lightVec.MutNeg().Reflect(hit.Normal).Dot(hit.Camera); reflectDotCam > 0 {
				res.MutAdd(light.Intensity.Scale(material.Specular() * math.Pow(reflectDotCam, material.Shininess())))
			}
		}
	}

	res.A = 1
	return &res
}
