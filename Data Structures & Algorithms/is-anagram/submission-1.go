func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}

	counts := make(map[rune]int)
	for _,i := range s{
		counts[i]++
	}
	for _,j := range t{
		counts[j]--	
	}
	for _,count := range counts{
		if count != 0{
			return false
		}
	}
	return true


}
