package main

import (
    "fmt"
    "os"
    "sort"
    "strings"
    "strconv"
)

func main() {
    //fmt.Println("reading content of file ")
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    //fmt.Println("content of file is : ", string(data))

    var leftList, rightList []int
    for _, line := range strings.Split(string(data), "\n") {
    if len(line) == 0 {
        continue
    }
    parts := strings.Fields(line)
     if len(parts) != 2 {
            fmt.Println("Invalid line format:", line)
            continue
    }
    left, err1 := strconv.Atoi(parts[0])
    right, err2 := strconv.Atoi(parts[1])
    if err1 != nil || err2 != nil {
        fmt.Println("Error converting string to int:", err1, err2)
        continue
    }
    leftList = append(leftList, left)
    rightList = append(rightList, right)
}

  sort.Ints(leftList)
  sort.Ints(rightList)
  //fmt.Println("leftList: ", leftList, "\n" )
  //fmt.Println("rightList: ", rightList)

  totalDistance := 0
    for i := 0; i < len(leftList); i++ {
        distance := abs(leftList[i] - rightList[i])
        totalDistance += distance
    }

     fmt.Println("Total Distance:", totalDistance)

}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}