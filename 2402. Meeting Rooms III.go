// coming back to understand the solution

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mostBooked(n int, meetings [][]int) int {
	if n == 1 {
		return 0
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	counts := make([]int32, n)

	unusedRooms := make(MinHeap, n)
	for i := 0; i < n; i++ {
		unusedRooms[i] = i
	}
	heap.Init(&unusedRooms)

	usedRooms := make(MinHeap, 0, n)

	for _, meeting := range meetings {
		for start := (meeting[0] + 1) << 8; len(usedRooms) != 0 && usedRooms[0] < start; {
			room := heap.Pop(&usedRooms).(int)
			room &= 0xFF
			heap.Push(&unusedRooms, room)
		}

		var room int
		var roomIdx int
		if len(unusedRooms) != 0 {
			room = heap.Pop(&unusedRooms).(int)
			roomIdx = room
			room |= meeting[1] << 8
		} else {
			room = heap.Pop(&usedRooms).(int)
			roomIdx = room & 0xFF
			room = (room>>8+meeting[1]-meeting[0])<<8 | roomIdx
		}
		counts[roomIdx]++
		heap.Push(&usedRooms, room)
	}

	result := 0
	maxCount := int32(1)
	for i, count := range counts {
		if count > maxCount {
			maxCount = count
			result = i
		}
	}
	return result
}