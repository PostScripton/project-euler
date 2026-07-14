package solution

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const length = 4

var grid [20][20]int

func init() {
	bytes, err := os.ReadFile("../grid.txt")
	if err != nil {
		fmt.Printf("Could not read grid.txt file: %v\n", err)
		os.Exit(1)
	}

	rows := strings.Split(string(bytes), "\n")

	for ri, row := range rows {
		if row == "" {
			continue
		}

		numbers := strings.Split(row, " ")

		for ni, number := range numbers {
			n, _ := strconv.Atoi(number)
			grid[ri][ni] = n
		}
	}
}
