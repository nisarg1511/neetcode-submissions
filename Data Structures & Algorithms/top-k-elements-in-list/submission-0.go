func topKFrequent(nums []int, k int) []int {
	freqMap:= make(map[int]int)

	for _, num:= range nums{
		freqMap[num]++
	}
	if len(freqMap) == 1{
		return []int{nums[0]}
	}
	inter:= make([][]int,len(nums))

	for j,v :=  range freqMap{
		inter[v] = append(inter[v],j)
	}
	sol:=make([]int,0)
	for i:=len(inter)-1;i>-1;i--{
		if len(inter[i]) > 0{
			for _,v:= range inter[i]{
				sol=append(sol,v)
				if len(sol)==k{
					return sol
				}
			}
		}
	}
	return []int{}
}
