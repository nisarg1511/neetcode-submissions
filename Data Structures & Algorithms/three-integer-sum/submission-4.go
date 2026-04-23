import "slices"
func threeSum(nums []int) [][]int {
	i:=0
	
	sol:=make([][]int,0)
	slices.Sort(nums)
	for i<len(nums)-2{
		if (i > 0) && (nums[i]==nums[i-1]){
			i++
			continue
		}
		j:=i+1
		k:=len(nums)-1
		for j<k{
			if k< len(nums)-1 && nums[k]==nums[k+1]{
				k--
				continue
			}
			if j > i+1 && nums[j]==nums[j-1]{
				j++
				continue
			}
			if (nums[j]+nums[k]) == (nums[i]*-1){
				sol = append(sol,[]int{nums[i],nums[j],nums[k]})
				j++
				k--
			}else if (nums[j]+nums[k]) < (nums[i]*-1){
				j++
			}else{
				k--
			}
		}
		i++
	}
	return sol
}