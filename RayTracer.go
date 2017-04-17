package gotracer

import "math/rand"

// RayTracer is where the magic happens
type RayTracer struct {
	scene    Scene
	maxDepth int
	sampling int
}

func (rt RayTracer) intersect(r Ray) Intersectable {
	var minObj Intersectable
	var inter Point
	min := infinitePoint
	for _, o := range rt.scene.Objects {
		inter = o.Intersection(r)
		if inter.Subtract(r.P).Length() > 1e-6 {
			if inter.Subtract(r.P).Length() < min.Subtract(r.P).Length() {
				min = inter
				minObj = o
			}
		}
	}
	return minObj
}

func (rt RayTracer) trace(r Ray, depth int) Color {
	rad := Color{0.0, 0.0, 0.0, 0.0}
	obj := rt.intersect(r)
	if obj == nil {
		return rad
	}
	pt := obj.Intersection(r)
	normal := obj.Normal(pt)
	reflected := r.Reflect(normal, pt)
	if depth > rt.maxDepth {
		rad.Clamp(1.0)
		return rad
	}

	rad = rad.Add(rt.trace(reflected, depth+1))

	rad.Clamp(1.0)
	return rad

}

func (rt RayTracer) calculate(x, y int) Color {
	c := Color{0.0, 0.0, 0.0, 0.0}
	for i := 1; i <= rt.sampling; i++ {
		for j := 1; j <= rt.sampling; j++ {
			r1 := rand.Float64()
			r2 := rand.Float64()
			one := float64(x) - r1 + float64(i)/float64(rt.sampling)
			two := float64(y) - r2 + float64(j)/float64(rt.sampling)
			p := Point{one, two, 0.0}
			v := Point{0.0, 0.0, 1.0}
			r := Ray{p, v}
			c = c.Add(rt.trace(r, 0))
		}
	}
	c = c.Multiply(1.0 / float64(rt.sampling*rt.sampling))
	return c
}
