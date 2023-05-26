package calculate

type Graph struct {
	adjList map[string][]string
	visited map[string]bool
	path    []string
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[string][]string),
		visited: make(map[string]bool),
		path:    make([]string, 0),
	}
}

func (g *Graph) AddEdge(startNode, endNode string) {
	g.adjList[startNode] = append(g.adjList[startNode], endNode)
}

func (g *Graph) DFS(node string) {
	g.visited[node] = true
	g.path = append(g.path, node)

	for _, destination := range g.adjList[node] {
		if !g.visited[destination] {
			g.DFS(destination)
		}
	}
}

func FindPath(nodes [][]string) []string {
	graph := NewGraph()
	startingPoint := make(map[string]int)
	for _, path := range nodes {
		graph.AddEdge(path[0], path[1])
		startingPoint[path[0]]++
		startingPoint[path[1]]--
	}
	var source string
	for key, value := range startingPoint {
		if value > 0 {
			source = key
			break
		}
	}
	graph.DFS(source)
	result := []string{graph.path[0], graph.path[len(graph.path)-1]}
	return result
}
