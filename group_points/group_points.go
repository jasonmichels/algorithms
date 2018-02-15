package main

import (
	"fmt"
)

func groupPoints(points []float64) [][]float64 {
	currentPoint := 0
	currentGroup := []float64{}
	groups := [][]float64{}

	for index, value := range points {
		if index == 0 {
			currentGroup = append(currentGroup, value)
			continue
		}

		if points[index]-points[currentPoint] <= 1 {
			currentGroup = append(currentGroup, value)
		} else {
			// create a new group
			currentPoint = index
			groups = append(groups, currentGroup)
			currentGroup = []float64{value}
		}
	}

	groups = append(groups, currentGroup)

	return groups
}

func main() {
	points := []float64{1, 1.1, 1.2, 1.3, 1.4, 2.1, 2.2, 2.3, 2.4, 3, 3.5, 3.6, 3.7, 4}
	group := groupPoints(points)
	fmt.Println(group)
}
