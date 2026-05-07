func checkInclusion(s1 string, s2 string) bool {
	sMap:=make(map[rune]int)
	freqMap:=make(map[rune]int)
	for _,s:=range s1{
		sMap[s]++
	}

	left := 0
	for i:=0;i<len(s2);i++{
		freqMap[rune(s2[i])]++
		if i-left+1 == len(s1){
			if isPerm(sMap,freqMap) {
				return true
			}
			freqMap[rune(s2[left])]--
			left++
		}
	}
	return false
}

func isPerm(m1,m2 map[rune]int) bool{
	for k,v := range m1{
		if v!= m2[k]{
			return false
		}
	}
	return true
}
