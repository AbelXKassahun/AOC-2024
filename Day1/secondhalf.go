package Day1

func SecondHalf(filepath string) int {
	// similarityScore := calcSimilarityScore(filepath)
	return calcSimilarityScore(filepath)
}

func calcSimilarityScore(filepath string) int {
	left, right, _ := FirstHalf(filepath)
	var similarityScore int
	var count int
	for _, val := range left {
		for _, val1 := range right{
			if val == val1 {
				count++
			} else if val < val1 {
				mult := val * count
				similarityScore += mult
				break
			}
		} 
		count = 0
	}
	return similarityScore
}