import "slices"

func groupAnagrams(strs []string) [][]string {
	indexMap:= make(map[string][]string)
	for _,str := range strs{
		indexMap[sortString(str)] = append(indexMap[sortString(str)],str)
	}

	sol:=make([][]string,len(indexMap))
	i:=0
	for _,v:=range indexMap{
			sol[i]=v
			i++
	}	
	return sol
}


func sortString(unsorted string)  string{
	runes := []rune(unsorted)
	slices.Sort(runes)
	sorted := string(runes)
	return sorted
}