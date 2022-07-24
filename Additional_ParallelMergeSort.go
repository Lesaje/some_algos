package main

//this wouldn't be fast, because at low level we will have to many calls
//we need composite sorts, for example, parallel merge sort on top, and american sort at the bottom of recursive calls
//but this is just a demonstration for concurrent programming

import (
    "fmt"
    "math/rand"
    "sync"
)

func ParallelMergeSort(input []int64) {
    var wait sync.WaitGroup
    wait.Add(1)
    mergeSort(input, &wait)
    wait.Wait()
}

func mergeSort(input []int64, wg *sync.WaitGroup) {
    var start sync.WaitGroup
    defer wg.Done()
    if len(input) > 1 {

        middlePoint := len(input) / 2
        left := make([]int64, len(input[:middlePoint]))
        right := make([]int64, len(input[middlePoint:]))
        copy(left, input[:middlePoint])
        copy(right, input[middlePoint:])

        start.Add(2)
        go mergeSort(left, &start)
        go mergeSort(right, &start)
        start.Wait()
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
    input := make([]int64, 16)
    testSlice := rand.Perm(16)
    for i := 16 - 1; i >= 0; i-- {
        input[i] = int64(testSlice[i])
    }
    ParallelMergeSort(input)
    fmt.Println(input)
}
