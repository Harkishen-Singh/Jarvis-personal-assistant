package utils

import (
	"fmt"
	"strings"
)

// RatingData as a type for ratings.
type RatingData struct {
	target string
	rating int
}

func compareTwoStrings(first string, second string) int {
	first = strings.Replace(first, " ", "", -1)
	second = strings.Replace(second, " ", "", -1)

	if len(first) == 0 && len(second) == 0 {
		return 100 // if both are empty strings
	}
	if len(first) == 0 || len(second) == 0 {
		return 0 // if only one is empty string
	}
	if first == second {
		return 100 // identical
	}
	if len(first) == 1 && len(second) == 1 {
		return 0 // both are 1-letter strings
	}
	if len(first) < 2 || len(second) < 2 {
		return 0 // if either is a 1-letter string
	}

	var firstBigrams map[string]int
	var bigram string
	var i, count int
	firstBigrams = make(map[string]int)
	for i = 0; i < len(first)-1; i++ {
		bigram = first[i : i+2]
		_, ok := firstBigrams[bigram]
		if ok {
			count = firstBigrams[bigram] + 1
		} else {
			count = 1
		}
		firstBigrams[bigram] = count
	}
	intersectionSize := 0

	for i = 0; i < len(second)-1; i++ {
		bigram = second[i : i+2]
		_, ok := firstBigrams[bigram]
		if ok {
			count = firstBigrams[bigram]
		} else {
			count = 0
		}
		if count > 0 {
			firstBigrams[bigram] = count - 1
			intersectionSize++
		}
	}

	return ((2.0 * intersectionSize * 100) / (len(first) + len(second) - 2))
}

//FindBestMatch to check for the ratings of every string and choose the best possible string
func FindBestMatch(mainString string, targetStrings []string) RatingData {

	var ratings RatingData
	var bestMatch RatingData
	var ratingsArray []RatingData
	bestMatchIndex := 0

	for i := 0; i < len(targetStrings); i++ {
		currentTargetString := targetStrings[i]
		currentRating := compareTwoStrings(mainString, currentTargetString)
		ratings.target = currentTargetString
		ratings.rating = currentRating
		ratingsArray = append(ratingsArray, ratings)
		if currentRating > ratingsArray[bestMatchIndex].rating {
			bestMatchIndex = i
		}
	}

	bestMatch = ratingsArray[bestMatchIndex]
	fmt.Println("BestMatch: ", bestMatch)

	return bestMatch
}
