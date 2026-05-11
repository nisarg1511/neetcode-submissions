func searchMatrix(matrix [][]int, target int) bool {
	low:=0
	high:=len(matrix)-1


	
	for low<=high{
	mid:=(low+high)/2

		if matrix[mid][0] <= target && matrix[mid][len(matrix[mid])-1] >= target{
			l:=0
			h:=len(matrix[mid])-1
			for l<=h{
			m:=(l+h)/2
				if matrix[mid][m]== target{
					return true
				}
				if matrix[mid][m] > target{
					h = m-1
				}else if matrix[mid][m] < target{
					l = m+1
				}
			}
			return false
		}
		if matrix[mid][0] > target {
			high = mid -1
		}else if matrix[mid][len(matrix[mid])-1] < target {
			low  =  mid+1
		}
	}
	return false
}
