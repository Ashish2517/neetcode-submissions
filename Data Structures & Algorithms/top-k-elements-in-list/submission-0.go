func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
    for _, num := range nums {
        countMap[num]++
    }

    buckets := make([][]int, len(nums)+1)
    for num, freq := range countMap {
        buckets[freq] = append(buckets[freq], num)
    }

    var result []int
    for i := len(buckets) - 1; i >= 0; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
            if len(result) >= k {
                return result[:k]
            }
        }
    }
    return result

}
