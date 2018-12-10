package tracer

import (
	"fmt"
	"math"
)

type Sphere struct {
	material     Material
	transform    Mat4x4
	invTransform Mat4x4
}

func NewSphere(material *Material) *Sphere {
	id := NewMat4x4ID()

	return &Sphere{
		material:     *material,
		transform:    *id,
		invTransform: *id,
	}
}

func NewTransformedSphere(transform *Mat4x4, material *Material) (*Sphere, error) {
	s := Sphere{
		material: *material,
	}

	if err := s.SetTransform(transform); err == nil {
		return &s, nil
	} else {
		return nil, err
	}
}

func (s *Sphere) Material() *Material {
	return &s.material
}

func (s *Sphere) SetTransform(transform *Mat4x4) error {
	inv, ok := transform.Inverse()

	if !ok {
		return fmt.Errorf("transform matrix is not invertible\n%s", transform)
	}

	s.transform = *transform
	s.invTransform = *inv
	return nil
}

func (s *Sphere) Intersections(ray *Ray) []Intersection {
	ray = ray.Transform(&s.invTransform)
	centerToRayOrig := ray.Origin.SubPoint(Origin()) // Sphere center is at origin
	a := ray.Direction.Dot(&ray.Direction)
	b := 2 * ray.Direction.Dot(centerToRayOrig)
	c := centerToRayOrig.Dot(centerToRayOrig) - 1 // Sphere radius is 1
	disc := b*b - 4*a*c

	if disc < 0 {
		return nil
	}

	sqrtDisc := math.Sqrt(disc)
	res := make([]Intersection, 1, 2)
	res[0] = NewIntersection((-b-sqrtDisc)/(2*a), s)

	if !ApproxEqual(disc, 0) {
		res = append(res, NewIntersection((-b+sqrtDisc)/(2*a), s))
	}

	return res
}

func (s *Sphere) NormalAt(p *Point) *Vector {
	return s.invTransform.T().DotVector(s.invTransform.DotPoint(p).SubPoint(Origin())).MutNormalized()
}
