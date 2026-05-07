func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2){
		return false
	}
	
	arr:=make([]int,26)
	arr1:=make([]int,26)

	for i:=0;i<len(s1);i++{
		arr[s1[i] - 'a'] ++
		arr1[s2[i] - 'a'] ++
	}

	matches:=0

	for i:=0;i<26;i++{
		if arr[i]==arr1[i]{
			matches++
		}
	}

	
	left := 0
	for right:=len(s1);right<len(s2);right++{
		if matches == 26{ 
			return true
		}

		index:= s2[right] - 'a'
		arr1[index]++
		if arr1[index] == arr[index]+1{
			matches--
		}else if arr1[index] == arr[index]{
			matches++
		}

		index = s2[left] - 'a'
		arr1[index]--
		if arr1[index]+1 == arr[index]{
			matches--
		}else if arr1[index] == arr[index]{
			matches++
		}
		left++
	}

	return matches==26
}
