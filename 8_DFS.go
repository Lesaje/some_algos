package main

import (
    "fmt"
)

// DFS returns enter, exit "timestep" for every vertex in graph
func DFS(graph map[uint32][]uint32) ([]uint32, []uint32) {
    startTime := make([]uint32, len(graph))
    finishTime := make([]uint32, len(graph))
    ifVertexVisited := make([]bool, len(graph))
    predecessor := make([]uint32, len(graph))
    time := uint32(0)

    //declare "closure" to avoid passing too many variables to dfsVisit function
    //not a real closure, because I need recursion
    var dfsVisit func(vertex uint32)
    dfsVisit = func(vertex uint32) {
        time += 1
        startTime[vertex] = time
        ifVertexVisited[vertex] = true
        for _, v := range graph[vertex] {
            if ifVertexVisited[v] == false {
                predecessor[v] = vertex
                dfsVisit(v)
            }
        }
        time += 1
        finishTime[vertex] = time
    }

    for i := 0; i < len(graph); i++ {
        if ifVertexVisited[i] == false {
            dfsVisit(uint32(i))
        }
    }
    return startTime, finishTime
}

func main() {
    graph := make(map[uint32][]uint32)
    graph[0] = []uint32{1, 2}
    graph[1] = []uint32{3, 4}
    graph[2] = []uint32{5, 6}
    graph[3] = []uint32{}
    graph[4] = []uint32{}
    graph[5] = []uint32{}
    graph[6] = []uint32{}
    fmt.Println(DFS(graph))
}
