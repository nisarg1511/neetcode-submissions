func trap(height []int) int {
    l,r:=1,len(height)-2
    lmax,rmax:=height[l-1],height[r+1]
    total:=0
    for l<=r{
        if lmax < rmax{
            for lmax > height[l]{
                total+=lmax-height[l]
                l++
            }
            lmax = height[l]
            l++
        }else {
            for rmax > height[r]{
                total+=rmax-height[r]
                r--
            }
            rmax=height[r]
            r--
        }
    }
    return total
}