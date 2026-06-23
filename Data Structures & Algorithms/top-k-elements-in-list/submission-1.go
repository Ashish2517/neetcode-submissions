func topKFrequent(nums []int, k int) []int {
	check := make(map[int]int)
    for _,num := range nums{
        check[num]++
    }
    buckets := make([][]int,len(nums)+1)
    for number,frequency := range check{
       buckets[frequency] = append(buckets[frequency],number)
    }
    var result []int
    for i := len(buckets)-1;i>=0;i--{
        if len(buckets[i])>0{
            result = append(result,buckets[i]...)
        }
    }
    if len(result)>=k{
        return result[:k]
    }
    return nil
}
