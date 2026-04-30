func carFleet(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	remainingTime := make([][]float64, len(position))

	for i := 0; i < len(position); i++ {
		remainingTime[i] = make([]float64, 2)
		remainingTime[i][0] = float64(position[i])
		remainingTime[i][1] = calculateRemainingTime(position[i], target, speed[i])
	}
	sortByPosition(remainingTime, 0, len(remainingTime)-1)
	fleet := 1
	lastTime := remainingTime[len(remainingTime)-1][1]
	for i := len(remainingTime) - 2; i >= 0; i-- {
		if remainingTime[i][1] > lastTime {
			fleet++
			lastTime = remainingTime[i][1]
		}
	}
	return fleet
}

func calculateRemainingTime(position, target, speed int) float64 {
	return float64(target-position) / float64(speed)
}

func sortByPosition(arr [][]float64, left, right int) {
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	sortByPosition(arr, left, mid)
	sortByPosition(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr [][]float64, left, mid, right int) {
	n1, n2 := mid-left+1, right-mid
	l := make([][]float64, n1)
	r := make([][]float64, n2)
	for i := 0; i < n1; i++ {
		l[i] = arr[left+i]
	}
	for i := 0; i < n2; i++ {
		r[i] = arr[mid+1+i]
	}
	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if l[i][0] < r[j][0] {
			arr[k] = l[i]
			i++
		} else {
			arr[k] = r[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = l[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = r[j]
		j++
		k++
	}
}