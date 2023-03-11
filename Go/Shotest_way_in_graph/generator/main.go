package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const (
	countOfFiles = 10

	nConst = 6

	amplitude = 3
)

func isInSlice(elem [2]int, slice [][2]int) bool {
	for _, n := range slice {
		if n == elem {
			return true
		}
	}
	return false
}

func Generate(name int) {
	//rand.Seed(time.Now().Unix())

	n := nConst
	k := rand.Intn(amplitude) + 1

	s := rand.Intn(n) + 1
	var f int
	for {
		f = rand.Intn(n) + 1
		if s != f {
			break
		}
	}

	pairs := make([][2]int, 0, n)
	pair := [2]int{0, 0}
	for i := 0; i < n; i++ {
		pair = [2]int{rand.Intn(2*amplitude+1) - amplitude, rand.Intn(2*amplitude+1) - amplitude}
		if isInSlice(pair, pairs) {
			i--
			continue
		}
		pairs = append(pairs, pair)
	}

	file, _ := os.Create(strconv.Itoa(name) + ".txt")
	defer file.Close()
	_, _ = file.WriteString(fmt.Sprintf("%d\n", n))
	for i := 0; i < n; i++ {
		_, _ = file.WriteString(fmt.Sprintf("%d %d\n", pairs[i][0], pairs[i][1]))
	}
	_, _ = file.WriteString(fmt.Sprintf("%d\n", k))
	_, _ = file.WriteString(fmt.Sprintf("%d %d\n", s, f))
}

func main() {
	for i := 5; i <= countOfFiles+4; i++ {
		Generate(i)
	}
}
