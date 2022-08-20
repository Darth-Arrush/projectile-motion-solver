/*Solves projectile motion questions given inputs.*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(2))
}

func MaxHeightTime(θ float64, u float64, g float64) float64 {
	return u * math.Sin(θ*math.Pi/180) / g
}

func Range(θ float64, u float64, g float64, heightDiff float64) float64 {
	vY := u * math.Sin(θ*math.Pi/180)
	t := 2 * ((-1 * vY) + math.Sqrt((vY*vY)+(2*g*heightDiff))) / (g)
	vX := u * math.Cos(θ*math.Pi/180)
	return vX * t
}

func MaxHeight(θ float64, u float64, g float64) float64 {
	return u * u * math.Sin(θ*math.Pi/180) * math.Sin(θ*math.Pi/180) / (2 * g)
}

func Position(θ float64, u float64, g float64, t float64) (float64, float64) {
	return u * math.Cos(θ*math.Pi/180) * t, u*math.Sin(θ*math.Pi/180) - 0.5*g*t*t
}

func Velocities(θ float64, u float64, g float64, t float64) (float64, float64) {
	return u * math.Cos(θ*math.Pi/180), u*math.Sin(θ*math.Pi/180) - g*t
}

func Velocity(vX float64, vY float64) float64 {
	return math.Sqrt((vX * vX) + (vY * vY))
}
