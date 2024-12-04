package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func isSafe(report string) bool {
    levelsStr := strings.Fields(report)
    levels := make([]int, len(levelsStr))
    for i, s := range levelsStr {
        level, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("Error in converting string to int")
            return false
        }
        levels[i] = level
    }

    if len(levels) < 2 {
        return false
    }

   
    for i := 0; i < len(levels)-1; i++ {
        diff := abs(levels[i+1] - levels[i])
        if diff < 1 || diff > 3 {
            return false
        }
    }

    
    isIncreasing := true
    isDecreasing := true

    for i := 0; i < len(levels)-1; i++ {
        if levels[i] >= levels[i+1] {
            isIncreasing = false
        }
        if levels[i] <= levels[i+1] {
            isDecreasing = false
        }
    }

    return isIncreasing || isDecreasing
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func countSafeReports(data string) int {
    reports := strings.Split(strings.TrimSpace(data), "\n")
    safeCount := 0
    for _, report := range reports {
        if isSafe(strings.TrimSpace(report)) {
            safeCount++
        }
    }
    return safeCount
}

func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error:", err)
        return
    }

    if len(data) == 0 {
        fmt.Println("Error: Input file is empty")
        return
    }

    input := string(data)
    safeCount := countSafeReports(input)
    fmt.Printf("Number of safe reports: %d\n", safeCount)
}