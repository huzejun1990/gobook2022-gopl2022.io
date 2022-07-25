// Embed demonstrates basic struct embedding.
// Embed 演示了基本的结构嵌入。
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // 注意：此处（和半径处）需要尾随逗号
	}

	fmt.Printf("%#v\n", w)
	//输出
	//main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	w.X = 42

	fmt.Printf("%#v\n", w)
	//输出
	// main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
}
