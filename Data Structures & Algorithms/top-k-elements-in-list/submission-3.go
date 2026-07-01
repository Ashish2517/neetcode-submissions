
func topKFrequent(nums []int, k int) []int {
	frequencymap := make(map[int]int)
    for _,num := range nums{
        frequencymap[num]++
    }
    buckets := make([][]int,len(nums)+1)
    for num,frequency := range frequencymap{
        buckets[frequency] = append(buckets[frequency],num)
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