package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type ChartItem struct {
	Word   string
	Amount int
}

func Top10(text string) []string {
	// r := fmt.Sprint(len(text))
	// fmt.Println("Length=" + r)

	textSplitted := strings.Split(text, " ")
	fmt.Println("==================")
	fmt.Println(textSplitted)
	fmt.Println("==================")

	words := []ChartItem{}

	for i, val := range textSplitted {
		fmt.Println(i)
		fmt.Println(val)

		// Count how many times current word is presented into text

		words = append(words, ChartItem{val, i})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].Amount > words[j].Amount })

	for i, val := range words {
		fmt.Println(i)
		fmt.Println(val)
	}

	return nil
}
