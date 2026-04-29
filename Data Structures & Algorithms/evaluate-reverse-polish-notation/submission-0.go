func evalRPN(tokens []string) int {
	stack:=make([]string,0)

	for _,t:=range tokens{
		switch t {
			case "*":
				n1,_:=strconv.Atoi(stack[len(stack)-2])
				n2,_:=strconv.Atoi(stack[len(stack)-1])
				ans:=n1*n2
				stack = stack[:len(stack)-2]
				stack = append(stack,strconv.Itoa(ans))
				break
			case "/":
				n1,_:=strconv.Atoi(stack[len(stack)-2])
				n2,_:=strconv.Atoi(stack[len(stack)-1])
				ans:=n1/n2
				stack = stack[:len(stack)-2]
				stack = append(stack,strconv.Itoa(ans))
				break
			case "-":
				n1,_:=strconv.Atoi(stack[len(stack)-2])
				n2,_:=strconv.Atoi(stack[len(stack)-1])
				ans:=n1-n2
				stack = stack[:len(stack)-2]
				stack = append(stack,strconv.Itoa(ans))
				break
			case "+":
				n1,_:=strconv.Atoi(stack[len(stack)-2])
				n2,_:=strconv.Atoi(stack[len(stack)-1])
				ans:=n1+n2
				stack = stack[:len(stack)-2]
				stack = append(stack,strconv.Itoa(ans))
				break
			default:
				stack = append(stack,t)
		}
	}
	ans,_:=strconv.Atoi(stack[len(stack)-1])
	return ans
}
