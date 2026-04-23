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
			
			if (nums[j]+nums[k]) == (nums[i]*-1){
				sol = append(sol,[]int{nums[i],nums[j],nums[k]})
				j++
				k--
                for nums[j]==nums[j-1] && j<k{
                    j++
                }
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