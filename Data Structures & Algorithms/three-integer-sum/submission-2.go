import "slices"
func threeSum(nums []int) [][]int {
	i:=0
	j:=1
	sol:=make([][]int,0)
	slices.Sort(nums)
	for i<len(nums)-2{
		if i >0 && nums[i]==nums[i-1]{
			i++
			j=i+1
			continue
		}
		if j<len(nums)-1{
			if j>i+1 && nums[j]==nums[j-1]{
				j++
				continue
			}
			for k:=j+1;k<len(nums);k++{
				if k>j+1 && nums[k]==nums[k-1]{
					continue
				}
				if nums[i]+nums[j]+nums[k]==0{
					sol = append(sol,[]int{nums[i],nums[j],nums[k]})
				}
			}
			j++
		}else{
			i++
			j=i+1
		}
	}
	return sol
}
