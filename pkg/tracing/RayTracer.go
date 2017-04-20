package tracing

import (
	"math/rand"

	d "github.com/mcgillowen/gotracer/pkg/design"
	p "github.com/mcgillowen/gotracer/pkg/primitives"
)

// RayTracer is where the magic happens
type RayTracer struct {
	scene    d.Scene
	maxDepth int
	sampling int
}

func (rt RayTracer) intersect(r p.Ray) p.Intersectable {
	var minObj p.Intersectable
	var inter p.Point
	min := p.InfinitePoint
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

func (rt RayTracer) trace(r p.Ray, depth int) p.Color {
	rad := p.Color{R: 0.0, G: 0.0, B: 0.0, A: 0.0}
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

func (rt RayTracer) calculate(x, y int) p.Color {
	c := p.Color{R: 0.0, G: 0.0, B: 0.0, A: 0.0}
	for i := 1; i <= rt.sampling; i++ {
		for j := 1; j <= rt.sampling; j++ {
			r1 := rand.Float64()
			r2 := rand.Float64()
			one := float64(x) - r1 + float64(i)/float64(rt.sampling)
			two := float64(y) - r2 + float64(j)/float64(rt.sampling)
			p1 := p.Point{X: one, Y: two, Z: 0.0}
			v := p.Point{X: 0.0, Y: 0.0, Z: 1.0}
			r := p.Ray{P: p1, V: v}
			c = c.Add(rt.trace(r, 0))
		}
	}
	c = c.Multiply(1.0 / float64(rt.sampling*rt.sampling))
	return c
}
