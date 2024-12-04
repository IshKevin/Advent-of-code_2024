package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))
    mulEnabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	matches := re.FindAllStringSubmatch(string(data), -1)
	//fmt.Println(matches)
    sum := 0
	for _, match := range matches {
		if match[0] == "do()" {
			mulEnabled = true
		} else if match[0] == "don't()" {
			mulEnabled = false
		} else if mulEnabled {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x*y
		}
	}

	fmt.Println("Sum:", sum)
}