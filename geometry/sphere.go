package geometry

import (
	"math"

	"github.com/MarioCdeS/romano/float"
	"github.com/MarioCdeS/romano/linalg"
)

type Sphere struct {
	material     Material
	transform    linalg.Mat4x4
	invTransform linalg.Mat4x4
}

func NewSphere(material *Material) *Sphere {
	id := linalg.NewMat4x4ID()

	return &Sphere{
		material:     *material,
		transform:    *id,
		invTransform: *id,
	}
}

func NewTransformedSphere(transform *linalg.Mat4x4, material *Material) (*Sphere, error) {
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

func (s *Sphere) SetTransform(transform *linalg.Mat4x4) error {
	if inv, err := transform.Inverse(); err == nil {
		s.transform = *transform
		s.invTransform = *inv
		return nil
	} else {
		return err
	}
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

	if !float.ApproxEqual(disc, 0) {
		res = append(res, NewIntersection((-b+sqrtDisc)/(2*a), s))
	}

	return res
}

func (s *Sphere) NormalAt(p *linalg.Point) *linalg.Vector {
	return s.invTransform.T().DotVector(s.invTransform.DotPoint(p).SubPoint(Origin())).MutNormalized()
}
