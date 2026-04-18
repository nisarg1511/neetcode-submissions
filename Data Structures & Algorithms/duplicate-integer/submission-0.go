func hasDuplicate(nums []int) bool {
    m1:= make(map[int]int)

	for _,i:=range nums{
		m1[i]++
		if m1[i] > 1 {
			return true
		}
	}
	return false
}
