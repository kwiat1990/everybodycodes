package solutions

import (
	"maps"
	"strings"
)

func Part1(input string) string {
	graph := map[string][]string{}

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, ":")
		graph[items[0]] = strings.Split(items[1], ",")
	}

	paths := getPath(graph, "RR")
	lookup := map[int][][]string{}

	for _, path := range paths {
		lookup[len(path)] = append(lookup[len(path)], path)
	}

	for val := range maps.Values(lookup) {
		if len(val) == 1 {
			return strings.Join(val[0], "")
		}
	}

	return ""
}

// Explanation: getPath is bottom-up version of DFS:
//
// getPath(graph, "RR")
// └─ neighbor "B"
//
//	└─ getPath(graph, "B")
//	  ├─ neighbor "@"
//	  │  └─ getPath(graph, "@")  // Returns [["@"]]
//	  └─ builds ["B", "@"]
//	└─ builds ["RR", "B", "@"]
func getPath(graph map[string][]string, node string) [][]string {
	var paths [][]string

	if node == "@" {
		return [][]string{{node}}
	}

	for _, neighbor := range graph[node] {
		for _, p := range getPath(graph, neighbor) {
			path := []string{node}
			path = append(path, p...)
			paths = append(paths, path)
		}
	}
	return paths
}

// func dfs(graph map[string][]string, node string, paths *[]string, path *[]string) {
// 	*path = append(*path, node)
//
// 	if node == "@" {
// 		output := make([]string, len(*path))
// 		copy(output, *path)
// 		*paths = append(*paths, strings.Join(output, ""))
// 	} else {
// 		for _, neighbor := range graph[node] {
// 			dfs(graph, neighbor, paths, path)
// 		}
// 	}
//
// 	*path = (*path)[:len(*path)-1]
// }
