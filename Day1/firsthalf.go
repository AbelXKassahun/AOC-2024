package Day1

import (
	"encoding/csv"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func FirstHalf(filepath string) ([]int, []int, int){
	input := inputParser(filepath)
	left, right := creatLeftNRight(input)

	distance := calcDistance(left, right)
	return left, right, distance
}

// part1
func calcDistance(left, right []int) int{
	var distance float64
	for i, val := range left {
		dif := val - right[i]
		distance += math.Abs(float64(dif))
	}

	return int(distance)
}

func creatLeftNRight(input [][]string) ([]int, []int){
	var left = make([]int, 1000)
	var right = make([]int, 1000)
	for i, val := range input {
		row := strings.Split(val[0], "   ")
		left[i], _ = strconv.Atoi(row[0])
		right[i], _ = strconv.Atoi(row[1])
	}
	sort.Ints(left) 
	sort.Ints(right)
	return left, right
}

func inputParser(filepath string) [][]string{
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	
	r := csv.NewReader(file)
	input, err := r.ReadAll()

	if err != nil {
		panic(err)
	}

	defer file.Close()

	return input
}