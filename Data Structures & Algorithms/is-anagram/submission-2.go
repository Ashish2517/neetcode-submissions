func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	check := make(map[rune]int)
	for _,chars := range s{
		check[chars]++
	}
	for _,chart:= range t{
		check[chart]--
	}
	for _,value := range check{
		if value != 0{
			return false
		}
	}
	return true
}
