package primitives

import (
	"math"
	"math/cmplx"
)

// InfinitePoint for representing a non-valid point
var InfinitePoint = Point{1e20, 1e20, 1e20}

// Lambda is to prevent
const lambda = 1e6

func solveQuadratic(floats [3]float64) (x1, x2 float64, ret bool) {
	a, b, c := complex(floats[0], 0.0), complex(floats[1], 0.0), complex(floats[2], 0.0)
	root := cmplx.Sqrt(cmplx.Pow(b, 2) - 4*a*c)
	x1Complex := (-b - root) / (2 * a)
	x2Complex := (-b + root) / (2 * a)
	ret = true
	if imag(x1Complex) != 0.0 || imag(x2Complex) != 0.0 {
		ret = false
	}
	x1 = real(x1Complex)
	x2 = real(x2Complex)
	return
}

// Intersectable are objects that can intercept rays
type Intersectable interface {
	Intersection(r Ray) Point
	Normal(p Point) Point
}

// Point in 3D
type Point struct {
	X, Y, Z float64
}

// Add points
func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Subtract points
func (a Point) Subtract(b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Multiply is the scalar multiplication of a vector
func (a Point) Multiply(b float64) Point {
	return Point{a.X * b, a.Y * b, a.Z * b}
}

// CrossProduct calculates the cross product between two vectors
func (a Point) CrossProduct(b Point) Point {
	return Point{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

// DotProduct calculates the dot product between two vectors
func (a Point) DotProduct(b Point) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Length of the point/vector
func (a Point) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Normalize the reference point/vector
func (a *Point) Normalize() {
	len := a.Length()
	a.X = a.X / len
	a.Y = a.Y / len
	a.Z = a.Z / len
}

// Reflect the ray at the point p with normal
func (r Ray) Reflect(normal, p Point) Ray {
	v := r.V.Subtract((normal.Multiply(2).Multiply((r.V.DotProduct(normal)))))
	return Ray{p, v}
}

// Ray representation from the camera
type Ray struct {
	P, V Point
}

// Triangle one of the objects available
type Triangle struct {
	V1, V2, V3 Point
	N          Point
}

// Sphere one of the objects available
type Sphere struct {
	Center Point
	Radius float64
}

// Intersection between a Ray and a Triangle
func (tri *Triangle) Intersection(r Ray) Point {
	side1 := tri.V2.Subtract(tri.V1)
	side2 := tri.V3.Subtract(tri.V1)
	side3 := r.P.Subtract(tri.V1)

	det := r.V.CrossProduct(side2).DotProduct(side1)
	invDet := 1.0 / det

	u := (r.V.CrossProduct(side2).DotProduct(side1)) * invDet
	if u < 0 || u > 1 {
		return InfinitePoint
	}

	v := (side3.CrossProduct(side1).DotProduct(r.V)) * invDet
	if v < 0 || v > 1 {
		return InfinitePoint
	}

	if u+v >= 1 {
		return InfinitePoint
	}
	t := (side3.CrossProduct(side1).DotProduct(side2)) * invDet
	if t < 0 {
		return InfinitePoint
	}

	t = t - lambda

	return r.P.Add(r.V.Multiply(t))
}

// Intersection calculation between a Ray and a Sphere
func (s *Sphere) Intersection(r Ray) Point {
	var coeff [3]float64
	coeff[0] = r.V.DotProduct(r.V)
	coeff[1] = 2.0 * (r.V.DotProduct(r.P.Subtract(s.Center)))
	coeff[2] = (r.P.Subtract(s.Center)).DotProduct((r.P.Subtract(s.Center))) - s.Radius*s.Radius
	x1, x2, real := solveQuadratic(coeff)
	if !real {
		return InfinitePoint
	}
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if x1 < 0 {
		x1 = x2
		if x1 < 0 {
			return InfinitePoint
		}
	}
	x1 = x1 - lambda
	return r.P.Add(r.V.Multiply(x1))
}

// Normal of the triangle, the point is inconsequential
func (tri *Triangle) Normal(p Point) Point {
	one := tri.V2.Subtract(tri.V1)
	two := tri.V3.Subtract(tri.V1)
	ret := one.CrossProduct(two)
	ret.Normalize()
	return ret
}

// Normal of a sphere at the point
func (s *Sphere) Normal(p Point) Point {
	ret := p.Subtract(s.Center)
	ret.Normalize()
	return ret
}
