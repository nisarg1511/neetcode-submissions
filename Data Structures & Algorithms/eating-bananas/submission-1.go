import "slices"
func minEatingSpeed(piles []int, h int) int {
	slices.Sort(piles)
	low:=1
	high:=piles[len(piles)-1]
	
	minValidRate:=piles[len(piles)-1]

	for low<=high{
		timeToEat:=0
		mid:=(low+high)/2

		for i:=0;i<len(piles);i++{
			currPile:=float64(piles[i])
			selectedRate:=float64(mid)
			timeToEat+=int(math.Ceil(currPile/selectedRate))
		}

		if timeToEat <= h {
			high = mid-1
			minValidRate = mid
		}else {
			low = mid+1
		}
	}
	
	return minValidRate
}
