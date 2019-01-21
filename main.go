package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadCorpse() map[string]bool {

	// Open CSV file
	f, err := os.Open("BingLiu.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	corpse := make(map[string]bool)
	for _, line := range lines {

		if line[1] == "positive" {
			corpse[line[0]] = true
		} else {
			corpse[line[0]] = false
		}
	}

	return corpse
}

func main() {
	corpse := loadCorpse()
	fmt.Println(corpse)
}
