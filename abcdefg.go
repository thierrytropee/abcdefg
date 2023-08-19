package main

import (
	"fmt"
	"os"
)

type point struct {
	pointName string
	nextPoint string
	distance  int
}

func pointPos(points *[7]point, pointsearch string) int {
	i := -1
	for index, point := range points {
		if point.pointName == pointsearch {
			i = index
			break
		}
	}
	return i

}

func main() {

	// var points [4]point doesn't seem to strongly type points
	points := [7]point{}
	points[0] = point{pointName: "A", nextPoint: "B", distance: 3}
	points[1] = point{pointName: "B", nextPoint: "C", distance: 1}
	points[2] = point{pointName: "C", nextPoint: "D", distance: 4}
	points[3] = point{pointName: "D", nextPoint: "E", distance: 1}
	points[4] = point{pointName: "E", nextPoint: "F", distance: 5}
	points[5] = point{pointName: "F", nextPoint: "G", distance: 9}
	points[6] = point{pointName: "G", nextPoint: "Nil", distance: 0}

	// Input the 2 strings
	var p, q, pq string
	fmt.Scan(&p, &q)

	if p > q {
		pq = p
		p = q
		q = pq
	}
	var p_pos int = pointPos(&points, p)
	var q_pos int = pointPos(&points, q)
	// fmt.Printf("p_pos %d, q_pos %d\n", p_pos, q_pos)
	if (p_pos < 0) || (q_pos < 0) || (p == q) {
		os.Exit(9)
	}

	var totalDistance int = 0
	var currentPoint = p_pos
	for range points {
		totalDistance = totalDistance + points[currentPoint].distance
		if points[currentPoint].nextPoint == " " {
			totalDistance = -1
			break
		}
		if points[currentPoint].nextPoint == q {
			break
		}
		currentPoint = pointPos(&points, points[currentPoint].nextPoint)
	}

	fmt.Println(totalDistance)
}
