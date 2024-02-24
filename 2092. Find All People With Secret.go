// clearly not my solution
// still pondering on it
type Meet struct {
	person int
	time   int
}

type MinHeap []Meet

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Meet))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make([][]Meet, n)
	for _, meeti := range meetings {
		graph[meeti[0]] = append(graph[meeti[0]], Meet{meeti[1], meeti[2]})
		graph[meeti[1]] = append(graph[meeti[1]], Meet{meeti[0], meeti[2]})
	}

	visited := make([]bool, n)
	queue := MinHeap{{0, 0}, {firstPerson, 0}}

	for len(queue) > 0 {
		ff := heap.Pop(&queue)
		if visited[ff.(Meet).person] {
			continue
		}
		visited[ff.(Meet).person] = true

		for _, v := range graph[ff.(Meet).person] {
			if v.time >= ff.(Meet).time && !visited[v.person]{
				heap.Push(&queue, v)
			}
		}
	}

	ans := []int{}

	for i, v := range visited {
		if v {
			ans = append(ans, i)
		}
	}

	return ans
}
