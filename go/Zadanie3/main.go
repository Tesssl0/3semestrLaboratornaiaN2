package main

import "fmt"

func pairToInt(pair string) int {
	if len(pair) < 2 {
		return 0
	}
	return int(pair[0]-'A')*26 + int(pair[1]-'A')
}

func main() {
	var genome1, genome2 string
	fmt.Scan(&genome1, &genome2)

	pairsSet := NewSet(1000)

	// Add all pairs from second genome to set
	for i := 0; i+1 < len(genome2); i++ {
		pair := genome2[i : i+2]
		pairValue := pairToInt(pair)
		pairsSet.Add(pairValue)

	}

	closeness := 0

	// Count pairs from first genome that exist in second
	for i := 0; i+1 < len(genome1); i++ {
		pair := genome1[i : i+2]
		pairValue := pairToInt(pair)

		if pairsSet.Contains(pairValue) {

			closeness++
		}
	}

	fmt.Println(closeness)
}
