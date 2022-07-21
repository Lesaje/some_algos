package main

import (
    "fmt"
    "mainProject/dfs"
)

func TopSort(graph map[uint32][]uint32) []uint32 {
    _, _, output := dfs.DFS(graph)
    return output
}

func main() {
    graph := make(map[uint32][]uint32)
    graph[0] = []uint32{1}
    graph[1] = []uint32{2}
    graph[2] = []uint32{3, 4, 5}
    graph[3] = []uint32{}
    graph[4] = []uint32{}
    graph[5] = []uint32{}
    fmt.Println(TopSort(graph))
}
