package main

// testing graph from here: https://youtu.be/oDqjPvD54Ss
// to launch, take first file from this repo, and place it in list directory, change package main to package list there
// after that import list package.

import (
    "fmt"
    "mainProject/list"
)

func BFS(graph map[uint32][]uint32, start uint32, end uint32) []uint32 {
    parents := findParents(graph, start)
    return reconstructPath(start, end, parents)
}

func findParents(graph map[uint32][]uint32, start uint32) []uint32 {
    var q list.List[uint32]
    q.Insert(start, q.Size())

    n := len(graph)
    isVisited := make([]bool, n)
    isVisited[start] = true
    parent := make([]uint32, n)

    for q.Size() > 0 {
        node, _ := q.Access(0)
        q.Erase(0, 0)
        neighbors := graph[node]

        for _, neighbor := range neighbors {
            if isVisited[neighbor] == false {
                q.Insert(neighbor, q.Size())
                isVisited[neighbor] = true
                parent[neighbor] = node
            }
        }
    }
    return parent
}

func reconstructPath(start, end uint32, prev []uint32) []uint32 {
    path := make([]uint32, 0)
    for e := end; e != start; e = prev[e] {
        path = append(path, e)
    }
    return append(path, start)
}

func main() {
    graph := make(map[uint32][]uint32)
    graph[0] = []uint32{7, 9, 11}
    graph[1] = []uint32{8, 10}
    graph[2] = []uint32{3, 12}
    graph[3] = []uint32{2, 4, 7}
    graph[4] = []uint32{3}
    graph[5] = []uint32{6}
    graph[6] = []uint32{5, 7}
    graph[7] = []uint32{0, 3, 6, 11}
    graph[8] = []uint32{1, 9, 12}
    graph[9] = []uint32{0, 8, 10}
    graph[10] = []uint32{1, 9}
    graph[11] = []uint32{0, 7}
    graph[12] = []uint32{2, 8}

    fmt.Println(BFS(graph, 4, 1))
}
