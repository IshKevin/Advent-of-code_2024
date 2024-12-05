package main

import (
    "bufio"
    "fmt"
    "os"
)


var directions = [8][2]int{
    {0, 1},   
    {0, -1},  
    {1, 0},   
    {-1, 0}, 
    {1, 1},   
    {1, -1},
    {-1, 1},  
    {-1, -1}, 
}

func isValid(x, y, rows, cols int) bool {
    return x >= 0 && x < rows && y >= 0 && y < cols
}

func hasXMAS(grid []string, x, y int, direction [2]int) bool {
    word := "XMAS"
    wordLen := len(word)
    rows := len(grid)
    cols := len(grid[0])

    for k := 0; k < wordLen; k++ {
        newX := x + k*direction[0]
        newY := y + k*direction[1]
        if !isValid(newX, newY, rows, cols) || grid[newX][newY] != word[k] {
            return false
        }
    }
    return true
}

func countOccurrences(grid []string, word string) int {
    count := 0
    rows := len(grid)
    cols := len(grid[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            for _, direction := range directions {
                if hasXMAS(grid, i, j, direction) {
                    count++
                }
            }
        }
    }
    return count
}

func readGridFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var grid []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        grid = append(grid, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return grid, nil
}

func main() {
    grid, err := readGridFromFile("input.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    word := "XMAS"
    result := countOccurrences(grid, word)
    fmt.Println("Total occurrences of", word, ":", result)
}