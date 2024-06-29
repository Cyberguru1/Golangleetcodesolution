//your song by Rita Ora

func getAncestors(n int, edges [][]int) (res [][]int ){

	gmap := map[int][]int{}

	for _, edge := range edges {
		if _, ok := gmap[edge[1]]; ok {
			gmap[edge[1]] = append(gmap[edge[1]], edge[0])
		}else {
			gmap[edge[1]] = []int{edge[0]}
		}
	}

	for i := 0; i < n; i++ {
		edge := []int{}
		visit := map[int]bool{}
		if con, ok := gmap[i]; ok {
			for len(con) > 0 {
				val := con[0]
				if _, ok := visit[val]; !ok{
					edge = append(edge, val)
					if conn, ok := gmap[val]; ok {
						con = append(con, conn...)
					}
					visit[val] = true
				}
				con = con[1:]
			}
		}
		sort.Ints(edge) // sadly makes it O[N*E log(E)]..where E is edges....sees no way out...cry5
		res = append(res, edge)
	}

	return 
}
