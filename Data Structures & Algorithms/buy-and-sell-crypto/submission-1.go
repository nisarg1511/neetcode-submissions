func maxProfit(prices []int) int {
	profit:=0
	curr:=prices[0]
	for i:=0;i<len(prices)-1;i++{
		
		if i==0 || prices[i]<curr{
			curr = prices[i]
			for j:=i+1;j<len(prices);j++{
				if prices[j]-prices[i] > profit{
					profit = prices[j]-prices[i]
				}
			}
		}
		
	}	
	return profit
}
