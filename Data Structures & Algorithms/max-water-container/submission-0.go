func maxArea(heights []int) int {
	i,j:=0,len(heights)-1
	largest:=0
	for i<j{
		h:= height(heights[i],heights[j])
		b:=j-i
		a:=h*b
		if a > largest{
			largest = a
		}
		if heights[i]>heights[j]{
			j--
		}else{
			i++
		}
	}
	return largest
}

func height(a,b int) int{
	if a>b{
		return b
	}
	return a
}