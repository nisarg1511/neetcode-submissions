func isPalindrome(s string) bool {
    s = strings.ReplaceAll(s," ","")
    s = strings.ToLower(s)
    i:=0
    j:=len(s)-1
    for {
        if i>=j{
            return true
        }
        r:=s[i]
        p:=s[j]
        validI:=(r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
        validJ:=(p >= 'a' && p <= 'z') || (p >= '0' && p <= '9')
        if validI && validJ {
            if r!=p{
                return false
            }
            i++
            j--
        }else if !validI{
            i++
        } else{
            j--
        }
    }
}
