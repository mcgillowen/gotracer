package design

import p "github.com/mcgillowen/gotracer/pkg/primitives"

// Scene contains all the intersectable objects
type Scene struct {
	Objects []p.Intersectable
}
