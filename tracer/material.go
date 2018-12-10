package tracer

import "math"

type Material struct {
	Color     Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func DefaultMaterial() *Material {
	return &Material{
		Color:     Color{1, 1, 1, 1},
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200,
	}
}

func NewMaterial(color *Color, ambient, diffuse, specular, shininess float64) *Material {
	m := Material{Color: *color}
	m.SetAmbient(ambient)
	m.SetDiffuse(diffuse)
	m.SetSpecular(specular)
	m.SetShininess(shininess)

	return &m
}

func (m *Material) Ambient() float64 {
	return m.ambient
}

func (m *Material) SetAmbient(ambient float64) {
	m.ambient = Clamp(ambient, 0, 1)
}

func (m *Material) Diffuse() float64 {
	return m.diffuse
}

func (m *Material) SetDiffuse(diffuse float64) {
	m.diffuse = Clamp(diffuse, 0, 1)
}

func (m *Material) Specular() float64 {
	return m.specular
}

func (m *Material) SetSpecular(specular float64) {
	m.specular = Clamp(specular, 0, 1)
}

func (m *Material) Shininess() float64 {
	return m.shininess
}

func (m *Material) SetShininess(shininess float64) {
	m.shininess = math.Max(shininess, 0)
}
