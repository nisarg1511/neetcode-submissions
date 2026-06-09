func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	heap := maxHeap(stones)

	for len(heap) > 1 {
		s1, h := popMax(heap)
		heap = h

		s2, h := popMax(heap)
		heap = h

		if s1 != s2 {
			heap = append(heap, s1-s2)

			// sift up
			j := len(heap) - 1
			for j > 0 {
				parent := (j - 1) / 2
				if heap[parent] < heap[j] {
					heap[parent], heap[j] = heap[j], heap[parent]
					j = parent
				} else {
					break
				}
			}
		}
	}

	if len(heap) == 0 {
		return 0
	}

	return heap[0]
}

func popMax(heap []int) (int, []int) {
	n := len(heap)

	maxVal := heap[0]

	if n == 1 {
		return maxVal, heap[:0]
	}

	heap[0] = heap[n-1]
	heap = heap[:n-1]

	heap = heapify(heap)

	return maxVal, heap
}

func maxHeap(arr []int) []int {
	heap := make([]int, 0, len(arr))

	for _, v := range arr {
		heap = append(heap, v)

		j := len(heap) - 1
		for j > 0 {
			parent := (j - 1) / 2
			if heap[parent] < heap[j] {
				heap[parent], heap[j] = heap[j], heap[parent]
				j = parent
			} else {
				break
			}
		}
	}

	return heap
}

func heapify(arr []int) []int {
	i := 0

	for {
		left := (i * 2) + 1
		right := (i * 2) + 2

		largest := i

		if left < len(arr) && arr[left] > arr[largest] {
			largest = left
		}

		if right < len(arr) && arr[right] > arr[largest] {
			largest = right
		}

		if largest == i {
			break
		}

		arr[i], arr[largest] = arr[largest], arr[i]
		i = largest
	}

	return arr
}