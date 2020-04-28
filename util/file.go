package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func createMatrix(file string) [][]int {
	lines := strings.Split(file, "\n")

	var matrix [][]int

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}

		parsedLine := strings.Split(lines[i], " ")

		var numLine []int
		for _, v := range parsedLine {
			num, _ := strconv.Atoi(v)
			numLine = append(numLine, num)
		}

		matrix = append(matrix, numLine)
	}

	return matrix
}

// ParseFile reads a file from the path and returns the matrix of the data
func ParseFile(path string) [][]int {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("failed to read file (%v): %v", path, err))
	}

	return createMatrix(string(data))
}
