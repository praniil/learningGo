package arrays

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perimeter() float64
}

type Rectangle struct{
	length, breadth float64
}
type Triangle struct{
	base, perpendicular float64
}

func (reactangle1 *Rectangle) area() float64{
	return reactangle1.length * reactangle1.breadth
}
func (reactangle1 *Rectangle) perimeter() float64{
	return 2*reactangle1.length + 2*reactangle1.breadth
}

func (rightTriangle *Triangle) area() float64{
	return 0.5 * rightTriangle.base * rightTriangle.perpendicular
}

func (rightTriangle *Triangle) perimeter() float64{
	return rightTriangle.base + rightTriangle.perpendicular + math.Sqrt(math.Pow(rightTriangle.base,2)+ math.Pow(rightTriangle.perpendicular,2))
}

func measure (g Geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func Interface () {
	rectangle1 := Rectangle{length: 10, breadth: 10}
	measure(&rectangle1)
	rightTriangle := Triangle{base: 10, perpendicular: 10}
	measure((&rightTriangle))
}


