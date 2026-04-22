import "slices"
func isValidSudoku(board [][]byte) bool {
	nums:=make(map[string][]int)
	for i,arr:= range board{
		for j,v:= range arr{
			s:=string(v)
			if s!="."{
				n,_:=strconv.Atoi(s)
				row:=strconv.Itoa(i)+"r"
				col:=strconv.Itoa(j)+"c"
				nums[row]=append(nums[row],n)
				nums[col]=append(nums[col],n)
				a:=strconv.Itoa((i/3)*3+j/3)
				nums[a]=append(nums[a],n)
				// if i <=2 && j<=2{
				// 	nums["9b"] = append(nums["9b"],n)
				// }else if i <=5 && j<=2{
				// 	nums["8b"] = append(nums["8b"],n)
				// }else if i <=8 && j<=2{
				// 	nums["7b"] = append(nums["7b"],n)
				// }else if i <=2 && j <=5{
				// 	nums["6b"] =  append(nums["6b"],n)
				// }else if i<=5 && j <=5{
				// 	nums["5b"] = append(nums["5b"],n)
				// }else if i<=8 && j<=5{
				// 	nums["4b"] = append(nums["4b"],n)
				// }else if i<=2 && j<=8{
				// 	nums["3b"] = append(nums["3b"],n)
				// }else if i<=5 && j<=8{
				// 	nums["2b"] = append(nums["2b"],n)
				// }else {
				// 	nums["1b"] = append(nums["1b"],n)
				// }
			}
		}
	}
	for _,v:=range nums{
		slices.Sort(v)
	}
	for _,v:=range nums{
		for i:=0;i<len(v)-1;i++{
			if v[i] == v[i+1]{
				return false
			}
		}
	}
	return true
}
