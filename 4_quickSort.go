package main

func quicksort(input []int) {
    if len(input) > 1 {
        pivot := input[0]
        lesser := make([]int, 0, len(input)/2)
        greater := make([]int, 0, len(input)/2)
        for _, v := range input[1:] {
            if v <= pivot {
                lesser = append(lesser, v)
            } else {
                greater = append(greater, v)
            }
        }
        quicksort(lesser)
        quicksort(greater)
        copy(input, append(append(lesser, pivot), greater...))
    }
}

func main() {
    
}
