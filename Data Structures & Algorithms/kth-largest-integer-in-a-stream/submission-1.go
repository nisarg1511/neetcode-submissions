type KthLargest struct {
	heap []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		heap: make([]int, 0),
		k:    k,
	}

	for _, num := range nums {
		obj.Add(num)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
	if len(this.heap) < this.k {
		this.push(val)
	} else if val > this.heap[0] {
		this.heap[0] = val
		this.heapifyDown(0)
	}

	return this.heap[0]
}

func (this *KthLargest) push(val int) {
	this.heap = append(this.heap, val)

	i := len(this.heap) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if this.heap[parent] <= this.heap[i] {
			break
		}

		this.heap[parent], this.heap[i] = this.heap[i], this.heap[parent]
		i = parent
	}
}

func (this *KthLargest) heapifyDown(i int) {
	n := len(this.heap)

	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < n && this.heap[left] < this.heap[smallest] {
			smallest = left
		}

		if right < n && this.heap[right] < this.heap[smallest] {
			smallest = right
		}

		if smallest == i {
			break
		}

		this.heap[i], this.heap[smallest] = this.heap[smallest], this.heap[i]
		i = smallest
	}
}