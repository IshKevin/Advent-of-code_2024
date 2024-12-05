package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    data, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error reading file")
        return
    }
    defer data.Close()

    scanner := bufio.NewScanner(data)
    var rules []string
    var updates [][]int
    isRules := true

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            isRules = false
            continue
        }
        if isRules {
            rules = append(rules, line)
        } else {
            update := parseUpdate(line)
            updates = append(updates, update)
        }
    }

    orderedUpdates := findOrderedUpdates(rules, updates)
    middleSum := sumMiddlePages(orderedUpdates)
    fmt.Println(middleSum)
}

func parseUpdate(line string) []int {
    parts := strings.Split(line, ",")
    var update []int
    for _, part := range parts {
        num, _ := strconv.Atoi(strings.TrimSpace(part))
        update = append(update, num)
    }
    return update
}

type Graph struct {
    adjList map[int]map[int]bool
}

func buildGraph(rules []string, update []int) *Graph {
    graph := &Graph{
        adjList: make(map[int]map[int]bool),
    }
    
    // Initialize adjacency lists for pages in the update
    updateSet := make(map[int]bool)
    for _, page := range update {
        updateSet[page] = true
        graph.adjList[page] = make(map[int]bool)
    }

    // Process rules that apply to pages in the update
    for _, rule := range rules {
        parts := strings.Split(rule, "|")
        from, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
        to, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
        
        if updateSet[from] && updateSet[to] {
            graph.adjList[from][to] = true
        }
    }

    return graph
}

func isOrdered(update []int, rules []string) bool {
    graph := buildGraph(rules, update)
    
    // Check each pair of pages
    for i := 0; i < len(update); i++ {
        for j := i + 1; j < len(update); j++ {
            // If there's a rule saying j should come before i, the order is wrong
            if graph.adjList[update[j]][update[i]] {
                return false
            }
        }
    }
    return true
}

func findOrderedUpdates(rules []string, updates [][]int) [][]int {
    var orderedUpdates [][]int
    for _, update := range updates {
        if isOrdered(update, rules) {
            orderedUpdates = append(orderedUpdates, update)
        }
    }
    return orderedUpdates
}

func sumMiddlePages(updates [][]int) int {
    sum := 0
    for _, update := range updates {
        midIndex := len(update) / 2
        sum += update[midIndex]
    }
    return sum
}