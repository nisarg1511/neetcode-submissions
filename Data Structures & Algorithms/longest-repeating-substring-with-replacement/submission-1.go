func characterReplacement(s string, k int) int {
    freq:=make(map[rune]int)
    left:=0
    maxFreq:=0
    maxLen:=0
    for i:=0;i<len(s);i++{
        char:=rune(s[i])
        freq[char]++
        window:=i -left + 1
        if  freq[char] > maxFreq{
            maxFreq = freq[char]
        }
        if window - maxFreq > k {
            freq[rune(s[left])]--
            left++
            window=i - left+1
            char = rune(s[left])
        }
        if window > maxLen{
            maxLen = window
        } 
    }
    return maxLen
}   