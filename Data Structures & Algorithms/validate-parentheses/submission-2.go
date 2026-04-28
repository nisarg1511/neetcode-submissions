func isValid(s string) bool {
    stack:=make([] rune,0)
	// if s[0] == '}' || s[0]==']' || s[0] == ')'{
	// 	return false
	// }
	for _,i:= range s{
		if i=='{' || i=='(' || i=='['{
			stack = append(stack,i)
		}else if len(stack) > 0 && isClosingof(stack[len(stack)-1],i){
			stack = stack[:len(stack)-1]
		}else{
			return false
		}
	}
	if len(stack) <1{
		return true
	}
	return false
}

func isClosingof(c,s rune) bool{
	if (c == '{' && s=='}') ||( c=='(' && s==')') || (c=='[' && s==']'){
		return true
	}
	return false
}
