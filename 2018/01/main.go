package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/input.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	lines := make([]string, 0)

	for {
		command, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		lines = append(lines, strings.TrimSpace(command))
	}

	p1, err := part1(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 1: %d\n", p1)

	p2, err := part2(lines)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part 2: %d", p2)
}

func part1(lines []string) (int, error) {
	freq := 0
	for _, command := range lines {
		change, err := strconv.Atoi(command)
		if err != nil {
			return 0, err
		}
		freq += change
	}

	return freq, nil
}

func part2(lines []string) (int, error) {
	freq := 0
	memory := map[int]int{
		0: 1,
	}

	for {
		for _, command := range lines {
			change, err := strconv.Atoi(command)
			if err != nil {
				return 0, err
			}
			freq += change
			memory[freq] += 1

			if memory[freq] > 1 {
				return freq, nil
			}
		}
	}
}
