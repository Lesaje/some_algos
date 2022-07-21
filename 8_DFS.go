package dfs

// DFS returns enter and exit "timestamp" for every vertex in graph
// and topologically sorted vertices (only for acyclic DAG)
func DFS(graph map[uint32][]uint32) ([]uint32, []uint32, []uint32) {
    startTime := make([]uint32, len(graph))
    finishTime := make([]uint32, len(graph))
    isVisited := make([]bool, len(graph))
    sorted := make([]uint32, 0, len(graph))
    time := uint32(0)

    //declare "closure" to avoid passing too many variables to dfsVisit function
    //not a real closure, because I need recursion
    var dfsVisit func(vertex uint32)
    dfsVisit = func(vertex uint32) {
        time += 1
        startTime[vertex] = time
        isVisited[vertex] = true
        for _, v := range graph[vertex] {
            if isVisited[v] == false {
                dfsVisit(v)
            }
        }
        time += 1
        finishTime[vertex] = time
        sorted = append(sorted, vertex)
    }

    for i := 0; i < len(graph); i++ {
        if isVisited[i] == false {
            dfsVisit(uint32(i))
        }
    }
    return startTime, finishTime, sorted
}
