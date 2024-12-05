package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
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

    var middleSum int
    for _, update := range updates {
        if !isOrdered(update, rules) {
            ordered := orderUpdate(update, rules)
            midIndex := len(ordered) / 2
            middleSum += ordered[midIndex]
        }
    }
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

    updateSet := make(map[int]bool)
    for _, page := range update {
        updateSet[page] = true
        graph.adjList[page] = make(map[int]bool)
    }

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

    for i := 0; i < len(update); i++ {
        for j := i + 1; j < len(update); j++ {
            if graph.adjList[update[j]][update[i]] {
                return false
            }
        }
    }
    return true
}

func orderUpdate(update []int, rules []string) []int {
    ordered := make([]int, len(update))
    copy(ordered, update)
    
    sort.SliceStable(ordered, func(i, j int) bool {
        graph := buildGraph(rules, update)
        return graph.adjList[ordered[i]][ordered[j]]
    })
    
    // If no direct rules exist, sort in descending order
    if !hasRules(ordered, rules) {
        sort.Sort(sort.Reverse(sort.IntSlice(ordered)))
    }
    
    return ordered
}

func hasRules(update []int, rules []string) bool {
    graph := buildGraph(rules, update)
    for _, from := range update {
        if len(graph.adjList[from]) > 0 {
            return true
        }
    }
    return false
}