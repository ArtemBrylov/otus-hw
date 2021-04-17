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

var splitterSpaces *regexp.Regexp = regexp.MustCompile("\\s+")

func Top10(text string) []string {

	textInOneRow := strings.ReplaceAll(text, "\n", " ")
	textWithoutTabs := strings.ReplaceAll(textInOneRow, "\t", "")
	textWithSingleSpaceBetween := splitterSpaces.ReplaceAllString(textWithoutTabs, " ")

	textSplitted := strings.Split(textWithSingleSpaceBetween, " ")

	words := countWordsInSlice(textSplitted)

	sort.Slice(words, func(i, j int) bool { return words[i].Amount > words[j].Amount })

	words = dd(words)

	return charItemsToStrings(words)
}

func countWordsInSlice(text []string) []ChartItem {
	words := []ChartItem{}

	for _, word := range text {
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

func dd(items []ChartItem) []ChartItem {
	result := [][]ChartItem{}
	var resultSlice []ChartItem

	fmt.Println(items)

	for i, item := range items {

		switch {
		case len(resultSlice) == 0:
			resultSlice = append(resultSlice, item)
		case i == len(items)-1:
			resultSlice = append(resultSlice, item)
			result = append(result, resultSlice)
		case resultSlice[0].Amount == item.Amount:
			resultSlice = append(resultSlice, item)
		case resultSlice[0].Amount > item.Amount:
			result = append(result, resultSlice)
			resultSlice = nil
			resultSlice = append(resultSlice, item)
		}
	}

	for _, val := range result {
		sort.Slice(val, func(i, j int) bool { return val[i].Word < val[j].Word })
	}
	fmt.Println(result)

	var top10 []ChartItem
	j := 0
	for _, val := range result {
		for _, val2 := range val {
			if j == 10 {
				break
			}
			top10 = append(top10, val2)
			j++
		}
	}
	fmt.Println(top10)

	return top10
}

func charItemsToStrings(chars []ChartItem) []string {
	s := []string{}
	for _, val := range chars {
		s = append(s, val.Word)
	}
	return s
}
