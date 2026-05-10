func checkInclusion(s1 string, s2 string) bool {
	
	if len(s1) > len(s2){
		return false
	}

	m1:=make([]int,26)
	m2:=make([]int,26)

	for i,v:=range s1{
		m1[v - 'a']++
		m2[rune(s2[i]) - 'a']++
	}

	matches:=0

	for i:=0;i<26;i++{
		if m1[i] == m2[i]{
			matches++
		}
	}

	left:=0
	right:=len(s1)
	
	for right<len(s2){
		if matches == 26{
			return true
		}

		char:=rune(s2[right]) - 'a'
		m2[char]++
		if m2[char]-1 == m1[char]{
			matches--
		} else if m2[char] == m1[char]{
			matches++
		}
		right++

		char=rune(s2[left]) - 'a'
		m2[char]--
		if m2[char] + 1 == m1[char]{
			matches--
		}else if m2[char] == m1[char]{
			matches ++
		}
		left++
	}
	if matches == 26{
		return true
	}
	return false
}
