// Safest Path soln

type UF struct {
	edgeTo, weight []int
	r, c   int
}

var moves = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func (uf *UF) neighbors(r, c int) []int {
	neigh := make([]int, 0, 4)
	for _, move := range moves {
		nR, nC := move[0]+r, move[1]+c
		i, err := uf.Point1D(nR, nC)
		if err == nil {
			neigh = append(neigh, i)
		}
	}
	return neigh
}

func (uf *UF) find(i int) int {
	if uf.edgeTo[i] == i {
		return i
	}
	uf.edgeTo[i] = uf.find(uf.edgeTo[i])
	return uf.edgeTo[i]
}

func (uf *UF) union(i, j int) {
	parI, parJ := uf.find(i), uf.find(j)
	if parI == parJ || min(uf.weight[parI], uf.weight[parJ]) == 0 {
		return
	}
	if uf.weight[parI] < uf.weight[parJ] {
		parI, parJ = parJ, parI
	}
	uf.edgeTo[parJ] = parI
	uf.weight[parI] += uf.weight[parJ]
}

func NewUF(r, c int) *UF {
	uf := UF{edgeTo: make([]int, r*c), weight: make([]int, r*c), r: r, c: c}
	for i := range len(uf.edgeTo) {
		uf.edgeTo[i] = i
	}
	return &uf
}

func (uf *UF) Point1D(r, c int) (int, error) {
	if r < 0 || r >= uf.r || c < 0 || c >= uf.c {
		return 0, errors.New("out of bounds")
	}
	return r*uf.c + c, nil
}

func (uf *UF) Add(r, c int) {
	i, _ := uf.Point1D(r, c)
	uf.weight[i] = 1
	for _, neigh := range uf.neighbors(r, c) {
		uf.union(i, neigh)
	}
}

func (uf *UF) Connects() bool {
	a, b := uf.find(0), uf.find(len(uf.edgeTo)-1)
	return a == b && min(uf.weight[a], uf.weight[b]) > 0
}

func maximumSafenessFactor(grid [][]int) int {
	uf := NewUF(len(grid), len(grid[0]))
	mlq := make([][][2]int, 0)

	mlq = append(mlq, make([][2]int, 0))
	for r := range len(grid) {
		for c, val := range grid[r] {
			if val == 1 {
				mlq[0] = append(mlq[0], [2]int{r, c})
				grid[r][c] = -1
			}
		}
	}

	for len(mlq[len(mlq)-1]) > 0 {
		last := mlq[len(mlq)-1]
		cur := make([][2]int, 0)

		for _, pt := range last {
			for _, delta := range moves {
				next := [2]int{pt[0] + delta[0], pt[1] + delta[1]}
				_, err := uf.Point1D(next[0], next[1])
				if err == nil && grid[next[0]][next[1]] >= 0 {
					cur = append(cur, next)
					grid[next[0]][next[1]] = -1
				}
			}
		}
		mlq = append(mlq, cur)
	}

	safe := len(mlq)
	for !uf.Connects() {
		safe--
		for _, pt := range mlq[safe] {
			uf.Add(pt[0], pt[1])
		}
	}

	return safe
}