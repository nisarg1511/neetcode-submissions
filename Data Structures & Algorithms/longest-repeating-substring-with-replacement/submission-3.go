func characterReplacement(s string, k int) int {
	left:=0
	maxLen:=0
	maxOcc:=0
	m:=make([]int,26)

	for right,v:=range s{
		m[v - 'A']+=1
		if m[v - 'A'] > maxOcc{
			maxOcc = m[v-'A']
		}
		if (right - left+ 1) - maxOcc > k{
			m[rune(s[left])-'A']-=1
			left++
		}
		if right - left + 1 > maxLen{
			maxLen = right -left + 1
		}
	}
		
	return maxLen
}
