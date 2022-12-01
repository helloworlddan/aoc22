package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const test bool = false

func main() {
	fileName := "./input.txt"
	if test {
		fileName = "./test_input.txt"
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	loads := make(map[int]int)
	var elf int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elf++
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		loads[elf] += calories
	}

	calories := []int{}
	for _, c := range loads {
		calories = append(calories, c)
	}

	sort.Ints(calories)

	l := len(calories)
	topTotal := calories[l-1] + calories[l-2] + calories[l-3]

	fmt.Printf("The top 3 elves carry a total of %d calories.\n", topTotal)
}
