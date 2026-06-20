func groupAnagrams(strs []string) [][]string {
	anagrammap := make(map[[26]int][]string)
	for _,word := range strs{

		var count[26]int

		for _,char := range word{
			count[char - 'a']++
		}
		anagrammap[count] = append(anagrammap[count], word)

	}
	var result [][]string

	for _, group := range anagrammap {
        result = append(result, group)
    }
	return result
}

