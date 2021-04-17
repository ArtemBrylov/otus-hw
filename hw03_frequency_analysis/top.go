package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type ChartItem struct {
	Word   string
	Amount int
}

// var splitter *regexp.Regexp = regexp.MustCompile(` *, *`)
var splitterSpaces *regexp.Regexp = regexp.MustCompile("\\s+")

func Top10(text string) []string {

	textInOneRow := strings.ReplaceAll(text, "\n", "")
	textWithoutTabs := strings.ReplaceAll(textInOneRow, "\t", "")
	textWithSingleSpaceBetween := splitterSpaces.ReplaceAllString(textWithoutTabs, " ")

	textSplitted := strings.Split(textWithSingleSpaceBetween, " ")
	// textSplitted := splitter.Split(textWithSingleSpaceBetween, -1)
	fmt.Println("==================")
	fmt.Println(textSplitted)
	fmt.Println("==================")

	words := countWordsInSlice(textSplitted)

	sort.Slice(words, func(i, j int) bool { return words[i].Amount > words[j].Amount })

	// fmt.Println("Print inner values")
	// fmt.Println(len(words))
	for i, val := range words {
		fmt.Println(i)
		fmt.Println(val)
	}
	// fmt.Println("End printing inner values")

	return charItemsToStrings(words)
}

func countWordsInSlice(text []string) []ChartItem {
	words := []ChartItem{}

	for _, word := range text {
		// fmt.Println(i)
		// fmt.Println(val)

		// Count how many times current word is presented into text
		if !contains(words, word) {
			amount := 0
			for _, item := range text {
				if word == item {
					amount += 1
				}
			}
			words = append(words, ChartItem{word, amount})
		}
	}
	return words
}

func contains(slice []ChartItem, item string) bool {
	for _, word := range slice {
		if word.Word == item {
			return true
		}
	}
	return false
}

func charItemsToStrings(chars []ChartItem) []string {
	s := []string{}
	for _, val := range chars {
		s = append(s, val.Word)
	}
	return s
}
