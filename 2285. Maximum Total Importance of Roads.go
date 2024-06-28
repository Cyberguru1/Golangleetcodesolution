type Nodes struct {
	node int
	conn int
}

func maximumImportance(n int, roads [][]int) (res int64) {
	mapp := map[int]int{}
	roadN := []Nodes{}

	for _, road := range roads {
		mapp[road[0]]++
		mapp[road[1]]++
	}

	for k, v := range mapp {
		roadN = append(roadN, Nodes{k, v})
	}

	sort.Slice(roadN, func(i, j int)bool{
		return roadN[i].conn < roadN[j].conn
	})

	for i := len(roadN) - 1; i >= 0; i--{
		mapp[roadN[i].node] = n
		n--
	}

	for _, road := range roads {
		res += int64(mapp[road[0]] + mapp[road[1]])
	}

	return
}