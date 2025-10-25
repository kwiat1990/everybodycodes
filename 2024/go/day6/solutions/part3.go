package solutions

import (
	"maps"
	"slices"
	"strings"
)

func Part3(input string) string {
	graph := map[string][]string{}

	for _, line := range strings.Split(input, "\n") {
		items := strings.Split(line, ":")
		graph[items[0]] = strings.Split(items[1], ",")
	}

	paths := dfs(graph, "RR")
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

func dfs(graph map[string][]string, start string) [][]string {
	var paths [][]string

	type stackItem struct {
		node    string
		path    []string
		visited map[string]struct{}
	}

	initialVisited := map[string]struct{}{start: {}}
	stack := []stackItem{{node: start, path: []string{start}, visited: initialVisited}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item.node == "@" {
			completePath := make([]string, len(item.path))
			copy(completePath, item.path)
			paths = append(paths, completePath)
			continue
		}

		neighbors := graph[item.node]
		slices.Reverse(neighbors)
		for _, neighbor := range neighbors {
			if _, exists := item.visited[neighbor]; !exists {
				visitedInBranch := maps.Clone(item.visited)
				visitedInBranch[neighbor] = struct{}{}

				newPath := make([]string, len(item.path), len(item.path)+1)
				copy(newPath, item.path)
				newPath = append(newPath, neighbor)

				stack = append(stack, stackItem{
					node:    neighbor,
					path:    newPath,
					visited: visitedInBranch,
				})
			}
		}
	}

	return paths
}
