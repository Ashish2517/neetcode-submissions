func twoSum(nums []int, target int) []int {

	check := make(map[int]int)

	for ind,val := range nums{
		reqno := target - val
		checkind,exist := check[reqno]
		if exist{
			return []int{checkind,ind}
		}
		check[val] = ind
		
	}
	return nil
}
