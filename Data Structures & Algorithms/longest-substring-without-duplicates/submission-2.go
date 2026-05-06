func lengthOfLongestSubstring(s string) int {
	
	length:=0
	hMap:=make(map[rune]int)
	currLen := 0
	lastDup := 0 
	for i:=0;i<len(s);i++{
		if hMap[rune(s[i])]>0 && hMap[rune(s[i])] > lastDup{ 
            lastDup = hMap[rune(s[i])]
			if currLen > length {
				length = currLen
			}
            currLen = i - hMap[rune(s[i])]+1
			hMap[rune(s[i])] = i+1
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
