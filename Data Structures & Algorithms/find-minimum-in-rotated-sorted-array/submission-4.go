func findMin(nums []int) int {
	left:=0
	right:=len(nums)-1
	lowest:=nums[right]
	if nums[left] < nums[right]{
		return nums[left]
	}
	for left<=right{
		mid:=(left+right)/2

		if nums[mid] < lowest{
			lowest =  nums[mid]
		}

		if nums[mid] < nums[left]{
			// if nums[left] > nums[right]{
				right = mid - 1
			// }else{
			// 	left = mid+1
			// }
		}else{
			if nums[left] < nums[right]{
				right = mid - 1
			}else{
				left = mid+1
			}
		}
	} 
	return lowest
}
