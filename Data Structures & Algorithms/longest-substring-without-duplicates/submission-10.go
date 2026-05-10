func lengthOfLongestSubstring(s string) int {
	maxLen:=0
	currLen:=0
	
	freqMap:=make(map[rune]int)
	left:=0

	for i,v:=range s{
		if freqMap[v] > 0 && freqMap[v] >  left  {
			if currLen > maxLen{
				maxLen = currLen
			}
			left = freqMap[v] 
			currLen = i - left + 1
		}else {
			currLen++
		}
		freqMap[v] = i+1
	}

	if currLen > maxLen{
		return currLen
	}
	return maxLen
}
