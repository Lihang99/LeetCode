func maxProfit(prices []int) int {
	minPrice,res := prices[0],0
	for i:=1;i<len(prices);i++{
		if prices[i]<minPrice{
			minPrice = prices[i]
		}else{
			if prices[i]-minPrice > res{
				res = prices[i]-minPrice
			}
		}
	}
	return res
}