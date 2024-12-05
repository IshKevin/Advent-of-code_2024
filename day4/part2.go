package main

import (
    "bufio"
    "fmt"
    "os"
)

// Direction vectors for all 8 possible directions
var directions = [8][2]int{
    {0, 1},   // right
    {0, -1},  // left
    {1, 0},   // down
    {-1, 0},  // up
    {1, 1},   // down-right
    {1, -1},  // down-left
    {-1, 1},  // up-right
    {-1, -1}, // up-left
}

func isValid(x, y, rows, cols int) bool {
    return x >= 0 && x < rows && y >= 0 && y < cols
}

func hasXMASPattern(grid []string, x, y int) bool {
    rows := len(grid)
    cols := len(grid[0])

    if !(1 <= x && x < rows-1 && 1 <= y && y < cols-1) {
        return false
    }
    if grid[x][y] != 'A' {
        return false
    }

    // Check both diagonals
    diag1 := string([]byte{grid[x-1][y-1], grid[x+1][y+1]})
    diag2 := string([]byte{grid[x-1][y+1], grid[x+1][y-1]})

    return (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM")
}

func countXMASPatterns(grid []string) int {
    count := 0
    rows := len(grid)
    cols := len(grid[0])

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if hasXMASPattern(grid, i, j) {
                count++
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

    result := countXMASPatterns(grid)
    fmt.Println("Total occurrences of X-MAS pattern:", result)
}