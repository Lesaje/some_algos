package main

import (
    "fmt"
    "math/rand"
)

func heapify(arr []int, n, i int) {
    largest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < n && arr[left] > arr[largest] {
        largest = left
    }
    if right < n && arr[right] > arr[largest] {
        largest = right
    }

    if largest != i {
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
    }
}

func HeapSort(arr []int) {
    for i := len(arr)/2 - 1; i >= 0; i-- {
        heapify(arr, len(arr), i)
    }
    for i := len(arr) - 1; i >= 0; i-- {
        arr[0], arr[i] = arr[i], arr[0]
        heapify(arr, i, 0)
    }
}

func main() {
    arr := rand.Perm(10)
    fmt.Println(arr)
    HeapSort(arr)
    fmt.Println(arr)
}
