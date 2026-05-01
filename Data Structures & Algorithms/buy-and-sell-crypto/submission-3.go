func maxProfit(prices []int) int {
    if len(prices) ==  2{
        if prices[0] > prices[1]{
            return 0
        }
        return prices[1] - prices[0]
    }
    
	min:=prices[0]
	profit:=0
	for i:=1;i<len(prices);i++{
		if prices[i] < min{
			min = prices[i]
		}
		if prices[i]-min > profit{
			profit = prices[i]-min
		}
	}
	return profit
}
