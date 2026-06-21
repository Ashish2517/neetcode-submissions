func groupAnagrams(strs []string) [][]string {
	check := make(map[[26]int][]string)
	for  _,str := range strs{
		var temp[26]int

		for _,char := range str{
			temp[char - 'a']++
		}
		check[temp] = append(check[temp],str)
	}
	var result [][]string
	for _,slice := range check{
		result = append(result,slice)
	}
	return result

}

