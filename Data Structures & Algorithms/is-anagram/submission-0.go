func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	m1:=make(map[rune]int)
	m2:=make(map[rune]int)

	for i:=range len(s){
		m1[rune(s[i])]++
		m2[rune(t[i])]++
	}

	for _,i:=range s{
		if m1[i]!=m2[i]{
			return false
		}
	}
	return true
}
