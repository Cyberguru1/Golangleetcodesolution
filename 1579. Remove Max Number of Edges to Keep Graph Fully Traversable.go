type ufind struct {
	parent    []int
	rank      []int
	localNode int
}

func maxNumEdgesToRemove(n int, edges [][]int) int {

	ali, bob := newUf(n), newUf(n)
	count := 0

	for _, edge := range edges {
		if edge[0] == 3 {
			uA := ali.Union(edge[1], edge[2])
			uB := bob.Union(edge[1], edge[2])
			if !uA && !uB {
				count++
			}
		}
	}

	for _, edge := range edges {

		if edge[0] == 1 {
			if !ali.Union(edge[1], edge[2]) {
				count++
			}
		} else if edge[0] == 2 {
			if !bob.Union(edge[1], edge[2]) {
				count++
			}
		}
	}

	if ali.localNode == n - 1 && bob.localNode == n - 1 {
		return count
	}


	return -1
}

func newUf(size int) *ufind {
	parents := make([]int, size+1)
	ranks := make([]int, size+1)

	for i := 1; i < size; i++ {
		parents[i] = i
		ranks[i] = 1
	}

	return &ufind{parent: parents, rank: ranks}
}

func (uf *ufind) Find(node int) int {
	if uf.parent[node] != node {
		return uf.Find(uf.parent[node])
	}

	return node
}

func (uf *ufind) Union(node1, node2 int) bool {
	p1, p2 := uf.Find(node1), uf.Find(node2)

	if p1 == p2 {
		return false
	} else if uf.rank[p1] > uf.rank[p2] {
		uf.parent[p2] = p1
		uf.rank[p1]++
	} else {
		uf.parent[p1] = p2
		uf.rank[p2] = p1
	}

	uf.localNode++
	return true
}