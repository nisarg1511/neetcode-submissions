func productExceptSelf(nums []int) []int {
	rev:=make([]int, len(nums))
	non:=make([]int,len(nums))
	sol:=make([]int, len(nums))
	rev[0]=nums[len(nums)-1]
	non[0] = nums[0]
	for i:=1;i<len(nums);i++{
		rev[i] = rev[i-1] * nums[(len(nums)-i)-1]
		non[i] = nums[i]*non[i-1]
	}
	for i:=1;i<len(nums)-1;i++{
		sol[i] = non[i-1]*rev[len(nums)-i-2]
	}
	sol[len(nums)-1] = non[len(nums)-2]
	sol[0] = rev[len(rev)-2]
	return sol
}
