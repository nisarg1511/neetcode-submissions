func lengthOfLongestSubstring(s string) int {
	
	length:=0
	hMap:=make(map[rune]int)
	currLen := 0
	
	for i:=0;i<len(s);i++{
		if hMap[rune(s[i])]>0{ 
			if currLen > length {
				length = currLen
			}
			i = hMap[rune(s[i])]
			hMap = make(map[rune]int)
			hMap[rune(s[i])] = i+1
			currLen = 1
		}else{
			currLen++
			hMap[rune(s[i])] = i+1
		}
	}
	if currLen > length{
		return currLen
	}
	return length
}
