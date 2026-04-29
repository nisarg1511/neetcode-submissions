func dailyTemperatures(temperatures []int) []int {
	sol:=make([]int,len(temperatures))
	stack:=make([]int,0)
	for i,j:=range temperatures{
		if len(stack) == 0 || j < temperatures[stack[len(stack)-1]]{
			stack=append(stack,i)
		}else {
			for len(stack)>0 && j > temperatures[stack[len(stack)-1]] {
				index:= stack[len(stack)-1]
				sol[index] = i-index
				stack = stack[:len(stack)-1]
			}
			stack = append(stack,i)
		}
	}
	return sol
}
