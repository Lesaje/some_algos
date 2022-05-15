package main

import (
    "fmt"
    "math/rand"
    "time"
)

func mergeSort(input []int) {
    if len(input) > 1 {

        middlePoint := len(input) / 2
        left := make([]int, len(input[:middlePoint]))
        right := make([]int, len(input[middlePoint:]))
        copy(left, input[:middlePoint])
        copy(right, input[middlePoint:])

        mergeSort(left)
        mergeSort(right)
        i, j, k := 0, 0, 0

        for i < len(left) && j < len(right) {
            if left[i] < right[j] {
                input[k] = left[i]
                i++
            } else {
                input[k] = right[j]
                j++
            }
            k++
        }

        for i < len(left) {
            input[k] = left[i]
            i++
            k++
        }
        for j < len(right) {
            input[k] = right[j]
            j++
            k++
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())
    testSlice := rand.Perm(10)
    fmt.Print("Input: ")
    fmt.Println(testSlice)
    mergeSort(testSlice)
    fmt.Print("Output: ")
    fmt.Println(testSlice)
}
