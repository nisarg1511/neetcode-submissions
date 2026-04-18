func twoSum(nums []int, target int) []int {
    m1:=make(map[int]int)
	m2:=make(map[int]int)
	for i,_ := range nums{
		if nums[i] == target/2 && m1[nums[i]]>0{
			return []int{m1[nums[i]]-1,i}
		}
		m1[target - nums[i]] = i+1
		m2[nums[i]] = i+1

		if m1[nums[i]]>0 && m2[nums[i]]>0 && m1[nums[i]]!=m2[nums[i]] {
			if m1[nums[i]]<m2[nums[i]]{
			return []int{m1[nums[i]]-1,m2[nums[i]]-1}
			}
			return []int{m2[nums[i]]-1,m1[nums[i]]-1}
		}
	}
	return []int{}
}
