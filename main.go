package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func getData(fileName string) []string {

	var train []string
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		train = append(train, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return train
}

func loadLexicon() map[string]bool {

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
	lexicon := loadLexicon()
	data := getData("training.txt")
	regex := regexp.MustCompile(`(?m) \w+`)

	var sum = 0
	fmt.Println(data[6875])
	for _, match := range regex.FindAllString(strings.ToLower(data[6875][1:]), -1) {
		_, found := lexicon[match[1:]]
		if found {
			if lexicon[match[1:]] {
				sum++
			} else {
				sum--
			}
		}
	}
	fmt.Println(sum)
}
