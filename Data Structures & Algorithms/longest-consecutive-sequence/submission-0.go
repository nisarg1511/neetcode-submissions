func longestConsecutive(nums []int) int {
	groups:=make(map[int] string)
	freq:=make(map[string] int)

	for _,n:=range nums{
		updateGroup(n,&groups)
	}
	for _,v:=range groups{
		freq[v]+=1
	}
	longest:=0
	for _,v:=range freq{
		if v>longest{
			longest=v
		}
	}
	return longest
}

func updateGroup(num int,groups *map[int]string) {
	if (*groups)[num-1]=="" {
		(*groups)[num] = strconv.Itoa(num)
	}else{
		(*groups)[num] = (*groups)[num-1]
	}
	if (*groups)[num+1]!=""{
		updateGroup(num+1,groups)
	}
	return
}
