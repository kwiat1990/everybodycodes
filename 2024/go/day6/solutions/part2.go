package solutions

import (
	"maps"
	"strings"
)

func Part2(input string) string {
	graph := map[string][]string{}

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, ":")
		graph[items[0]] = strings.Split(items[1], ",")
	}

	paths := bfs(graph, "RR")
	lookup := map[int][]string{}

	var builder strings.Builder
	for _, path := range paths {
		builder.Reset()
		for _, p := range path {
			builder.WriteByte(p[0])
		}
		s := builder.String()
		lookup[len(s)] = append(lookup[len(s)], s)
	}

	for val := range maps.Values(lookup) {
		if len(val) == 1 {
			return val[0]
		}
	}

	return ""
}

func bfs(graph map[string][]string, node string) [][]string {
	var paths [][]string
	var queue [][]string

	queue = append(queue, []string{node})

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == "@" {
			paths = append(paths, path)
		}

		for _, neighbor := range graph[current] {
			var newPath []string
			newPath = append(newPath, path...)
			newPath = append(newPath, neighbor)
			queue = append(queue, newPath)
		}
	}

	return paths
}
