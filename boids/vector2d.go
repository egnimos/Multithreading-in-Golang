package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

// create a method for this struct

// Add 2 vectors
func (vec Vector2D) Add(v Vector2D) Vector2D {
	return Vector2D{vec.x+v.x, vec.y+v.y}
}

// subtract 2 vectors
func (vec Vector2D) Subtract(v Vector2D) Vector2D {
	return Vector2D{vec.x-v.x, vec.y-v.y}
}

//multiply 2 vectors
func (vec Vector2D) Multiply(v Vector2D) Vector2D {
	return Vector2D{vec.x*v.x, vec.y*v.y}
}

// add vector with single float value
func (vec Vector2D) AddV(d float64) Vector2D {
	return Vector2D{
		x: vec.x+ d,
		y: vec.y+ d,
	}
}

// subtract vector with single float value
func (vec Vector2D) SubtractV(d float64) Vector2D {
	return Vector2D{
		x: vec.x- d,
		y: vec.y- d,
	}
}

// multiply vector with single float value
func (vec Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{
		x: vec.x* d,
		y: vec.y* d,
	}
}

// divide vector with single float value
func (vec Vector2D) DivideV(d float64) Vector2D {
	return Vector2D{
		x: vec.x / d,
		y: vec.y / d,
	}
}

// get the limit value as value cannot cross the upper and lower bound
func (vec *Vector2D) Limit(lower, upper float64) Vector2D {
	// check if the value is more than lower if it is not return lower
	// and then check if the value is less than upper if it is not then return upper value
	// value should be in the range of upper and lower value
	x := math.Min(math.Max(vec.x, lower), upper)
	y := math.Min(math.Max(vec.y, lower), upper)
	
	return Vector2D{x, y}
}

// calculate the distance between 2 vectors in 2d plane
func (vec *Vector2D) Distance(v Vector2D) float64 {
	return math.Sqrt(math.Pow(vec.x - v.x, 2) + math.Pow(vec.y - v.y, 2))
}