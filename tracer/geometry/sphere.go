package geometry

import (
	"math"

	"github.com/MarioCdeS/romano/tracer/float"
	"github.com/MarioCdeS/romano/tracer/linalg"
)

type Sphere struct {
	transform    linalg.Mat4x4
	invTransform linalg.Mat4x4
}

func NewSphere() *Sphere {
	id := linalg.NewMat4x4ID()
	return &Sphere{*id, *id}
}

func NewTransformedSphere(transform *linalg.Mat4x4) (*Sphere, error) {
	var res Sphere

	if err := res.SetTransform(transform); err == nil {
		return &res, nil
	} else {
		return nil, err
	}
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
	centerToRayOrig := ray.Origin.SubPoint(&Origin) // Sphere center is at origin
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

func (s *Sphere) NormalAt(p *linalg.Point) (*linalg.Vector, bool) {
	res := s.invTransform.DotPoint(p).SubPoint(&Origin)
	return res, float.ApproxEqual(res.SquaredMagnitude(), 1)
}
