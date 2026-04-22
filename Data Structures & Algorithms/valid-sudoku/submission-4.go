func isValidSudoku(board [][]byte) bool {
    rlen:=len(board)
    clen:=len(board[0])

    for i:=0;i<rlen;i++{
        hashMap:=make(map[byte]bool)
        for j:=0;j<clen;j++{
            if string(board[i][j])=="."{
                continue
            }            
            if _,ok:=hashMap[board[i][j]];!ok{
                hashMap[board[i][j]] = true
            }else {
                return false
            }
        }
    }

    for i:=0;i<clen;i++{
        hashMap:=make(map[byte]bool)
        for j:=0;j<rlen;j++{
            if string(board[j][i])=="."{
                continue
            }  
            if _,ok:=hashMap[board[j][i]];!ok{
                hashMap[board[j][i]] = true
            }else {
                return false
            }
        }
    }
    boxStartPoints := [][2]int{{0, 0}, {3,0}, {6,0}, {0,3}, {3,3}, {6,3}, {0,6}, {3,6}, {6,6}}
    // box check
    for i := 0; i < len(boxStartPoints); i++ {
        r, c := boxStartPoints[i][0], boxStartPoints[i][1]
        hashmap := make(map[byte]bool)
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                if string(board[r+j][c+k]) == "." {
                    continue
                }
                // fmt.Println(r+j, c+k)
                // fmt.Println(string(board[r+j][c+k]))
        
                if _, ok := hashmap[board[r+j][c+k]]; !ok {
                    hashmap[board[r+j][c+k]] = true
                } else {
                    return false
                }

            }
            
        }
    }

    return true
}
