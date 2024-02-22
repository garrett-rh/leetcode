package adjacent

func restoreArray(adjacentPairs [][]int) []int {
	if len(adjacentPairs) == 1 {
		return adjacentPairs[0]
	}

	// iterate through pairs, create list of "neighbors"
	graph := make(map[int][]int)
	for _, v := range adjacentPairs {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	visited := make(map[int]bool)
	resp := make([]int, len(adjacentPairs)+1)
	for k := range graph {
		if len(graph[k]) == 1 {
			dfs(k, 0, resp, visited, graph)
			break
		}
	}

	return resp
}

func dfs(num int, k int, ans []int, visited map[int]bool, graph map[int][]int) {
	if _, ok := visited[num]; ok {
		return
	}

	visited[num] = true
	ans[k] = num
	for _, v := range graph[num] {
		dfs(v, k+1, ans, visited, graph)
	}
}
