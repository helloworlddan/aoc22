package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var max int
	var heaviest int
	for e, c := range loads {
		if c > max {
			max, heaviest = c, e
		}
	}

	fmt.Printf("Elf %d carries %d calories.\n", heaviest, loads[heaviest])
}
