package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	mergeArray := append(arrayA, arrayB...)
	arrayClone := make(map[string]int)
	for _, val := range mergeArray {
		arrayClone[val] = 1
	}
	var mergeArrayFinal []string
	for newValue := range arrayClone {
		mergeArrayFinal = append(mergeArrayFinal, newValue)
	}
	return mergeArrayFinal
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}
