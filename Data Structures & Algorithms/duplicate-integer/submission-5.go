func hasDuplicate(nums []int) bool {
	check := make(map[int]int)
	for ind,num := range nums{
		_,exist := check[num]
		if exist{
			return true
		}
		check[num]=ind
	}
	return false
}