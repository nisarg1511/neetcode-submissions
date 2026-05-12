func search(nums []int, target int) int {
	left:=0
	right:=len(nums)-1	

	for left <=right{
		mid:=(left +  right)/2
		
		if nums[mid]== target{
			return mid
		}else if nums[left] == target{
			return left
		}else if nums[right] == target {
			return right
		}
		
		if nums[left] < nums[right]{
			if target > nums[mid]{
				left =  mid +1
			}else {
				right =  mid - 1 
			}
		}else if nums[mid] > nums[left]  {
				if nums[mid] > target{
					if nums[right] >= target{
						left = mid+1
					}else{
						right = mid - 1
					}
				}else {
				 	left= mid+1
				}	
		}else {
				if target > nums[mid] && target <= nums[right]{
					left  = mid+1
				}else{
					right = mid -1
				}
		}
	}

	return -1
}
