package main

import (
	"fmt"
	"math"
	"os"
)

type cityLocation struct {
	x int
	y int
}

type city struct {
	distance int
	parent   int
}

func isInSlice(elem int, slice []int) bool {
	for _, n := range slice {
		if n == elem {
			return true
		}
	}
	return false
}

func cityDistance(city1 cityLocation, city2 cityLocation) int {
	return int(math.Abs(float64(city1.x-city2.x)) + math.Abs(float64(city1.y-city2.y)))
}

func readFile(name string) (int, []cityLocation, int, int, int) {
	fileIn, _ := os.Open(name)
	defer fileIn.Close()
	var n int
	_, _ = fmt.Fscanf(fileIn, "%d\n", &n)
	cities := make([]cityLocation, n, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanf(fileIn, "%d %d\n", &cities[i].x, &cities[i].y)
	}
	var k int
	_, _ = fmt.Fscanf(fileIn, "%d\n", &k)
	var startCity, endCity int
	_, _ = fmt.Fscanf(fileIn, "%d %d\n", &startCity, &endCity)
	return n, cities, k, startCity - 1, endCity - 1
}

func RealMain(name string) int {
	n, cities, k, startCity, endCity := readFile(name)
	visitedCities := make(map[int]int, n)
	citiesTable := make(map[int]city, n)
	queue := make([]int, 0, n)

	if startCity == endCity {
		return 0
	}

	queue = append(queue, startCity)
	citiesTable[startCity] = city{0, startCity}

	var currentCity int
	for {
		if len(queue) == 0 {
			break
		}
		currentCity = queue[0]
		visitedCities[queue[0]] = 1
		queue = queue[1:]
		for j, location := range cities {
			cDistance := cityDistance(location, cities[currentCity])
			if _, visited := visitedCities[j]; !visited && cDistance <= k {
				if inQueue := isInSlice(j, queue); inQueue {
					continue
				}
				queue = append(queue, j)
				citiesTable[j] = city{1 + citiesTable[currentCity].distance, currentCity}
				if j == endCity {
					return citiesTable[j].distance
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(RealMain("input.txt"))
}
