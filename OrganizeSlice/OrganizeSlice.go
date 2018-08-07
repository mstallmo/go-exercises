package OrganizeSlice

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func findLongestName(names []string) int {
	longestName := 0
	for _, v := range names {
		if len(v) > longestName {
			longestName = len(v)
		}
	}
	return longestName
}

func Run() {
	sortedNames := make([][]string, findLongestName(names))
	for _, v := range names {
		if (len(v) - 1) > len(sortedNames) {
			newSlice := make([][]string, len(sortedNames)+1)
			copy(newSlice, sortedNames)
			sortedNames = newSlice
		}
		sortedNames[len(v)-1] = append(sortedNames[len(v)-1], v)
	}
	fmt.Print(sortedNames)
}
