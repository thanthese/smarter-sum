package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(smarterSum(string(bytes)))
}

func smarterSum(s string) string {
	foundNum := false
	sum := 0.0
	highestPrecision := 0
	usedCommas := false

	s = regexp.MustCompile("[$]").ReplaceAllString(s, "")
	s = regexp.MustCompile("[^0-9.,-]+").ReplaceAllString(s, " ")
	for _, field := range strings.Fields(s) {
		noCommas := regexp.MustCompile(",").ReplaceAllString(field, "")
		float, err := strconv.ParseFloat(noCommas, 64)
		if err != nil {
			continue
		}
		foundNum = true
		sum += float
		p := getPrecision(field)
		if p > highestPrecision {
			highestPrecision = p
		}
		if strings.Contains(field, ",") {
			usedCommas = true
		}
	}
	if !foundNum {
		return ""
	}
	prettyNum := strconv.FormatFloat(sum, 'f', highestPrecision, 64)
	if usedCommas {
		prettyNum = addCommas(prettyNum)
	}
	return prettyNum
}

func addCommas(s string) (pretty string) {
	end := len(s) - 1
	if strings.Contains(s, ".") {
		end = strings.Index(s, ".") - 1
		pretty = s[end+1:]
	}
	j := 0
	for i := end; i >= 0; i-- {
		if !strings.ContainsAny(string(s[i]), "1234567890") {
			pretty = s[:i+1] + pretty
			break
		}
		if j == 3 {
			j = 0
			pretty = "," + pretty
		}
		j++
		pretty = string(s[i]) + pretty
	}
	return
}

func getPrecision(s string) int {
	matches := regexp.MustCompile("^-?[0-9,]*[.]([0-9]*)$").FindStringSubmatch(s)
	if len(matches) != 2 {
		return 0
	}
	return len(matches[1])
}
