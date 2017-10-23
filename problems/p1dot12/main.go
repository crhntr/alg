package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var filename = "input.txt"

func init() {
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
}

func main() {
	fileContents, err := ioutil.ReadFile(filename)
	mustBeNil(err)
	infix, postfix, prefix := parseFile(string(fileContents))

	// displayFix("infix", infix)
	// displayFix("prefix", prefix)
	// displayFix("postfix", postfix)

	if len(infix) == 0 && len(prefix) == len(postfix) {
		infix = GetInfix(prefix, postfix)
	} else if len(prefix) == 0 && len(infix) == len(postfix) {
		prefix = GetPrefix(infix, postfix)
	} else if len(postfix) == 0 && len(prefix) == len(infix) {
		postfix = GetPostfix(infix, prefix)
	}

	displayFix("infix", infix)
	displayFix("prefix", prefix)
	displayFix("postfix", postfix)
}

func parseFile(input string) (infix, postfix, prefix []int) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if nums, ok := parseLine("prefix", line); ok {
			prefix = nums
			continue // goes to next line
		}
		if nums, ok := parseLine("postfix", line); ok {
			postfix = nums
			continue // goes to next line
		}
		if nums, ok := parseLine("infix", line); ok {
			infix = nums
			continue // goes to next line
		}
	}

	return
}

func parseLine(fixname, line string) ([]int, bool) {
	if !strings.HasPrefix(fixname, ": ") {
		fixname = fixname + ": "
	}
	if !strings.HasPrefix(line, fixname) {
		return []int{}, false
	}

	nums := []int{}
	line = strings.TrimPrefix(line, fixname)
	for _, numStr := range strings.Split(line, ",") {
		numStr = strings.TrimSpace(numStr)
		num, err := strconv.Atoi(numStr)
		mustBeNil(err)
		nums = append(nums, num)
	}
	return nums, true
}

func mustBeNil(err error) {
	if err != nil {
		panic(err)
	}
}

func displayFix(fixName string, arr []int) {
	fmt.Printf("%s: ", fixName)

	if len(arr) > 1 {
		fmt.Print(arr[0])
		for _, num := range arr[1:] {
			fmt.Printf(",%d", num)
		}
	}
	fmt.Println()
}
