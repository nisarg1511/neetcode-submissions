func longestConsecutive(nums []int) int {
    set:=make(map[int]string)
    for _,n:=range nums{
        set[n]="1"
    }
    current:=1
    longest :=0
    for n,_:=range set{
        if set[n] !="" && set[n-1]==""{
            curr:=n
            for {
                if set[curr+1]!=""{
                    current++
                    set[curr] = ""
                    curr++
                }else{
                    if current>longest{
                        longest = current
                    }
                    current = 1
                    break
                }
            }
        }
    }
    return longest
}