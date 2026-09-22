package main

import (
	"fmt"
	"slices"
)

func CompressCoordinates(coords []int, doAlloc bool) (
	distCoords []int, compCoords map[int]int) {

	capacity := 0
	if doAlloc {
		capacity = len(coords)
	}
	added := make(map[int]int, capacity)
	distCoords = make([]int, 0, capacity)
	for _, x := range coords {
		if added[x] == 0 {
			distCoords = append(distCoords, x)
			added[x] = 1
		}
	}
	slices.Sort(distCoords)
	compCoords = added
	for i, x := range distCoords {
		compCoords[x] = i
	}
	return distCoords, compCoords
}

func main() {
	coords := []int{198, -71, 3, 0, 567891}
	distCoords, compCoords := CompressCoordinates(coords, true)
	fmt.Println(distCoords)
	fmt.Println(compCoords)
}
