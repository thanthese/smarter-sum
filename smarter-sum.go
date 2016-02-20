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
	for _, field := range getFields(s) {
		float, err := strconv.ParseFloat(removeCommas(field), 64)
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

func addCommas(s string) string {
	pretty := ""
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
	return pretty
}

func getFields(s string) []string {
	s = regexp.MustCompile("[$]").ReplaceAllString(s, "")
	s = regexp.MustCompile("[^0123456789.,-]+").ReplaceAllString(s, " ")
	return strings.Fields(s)
}

func removeCommas(s string) string {
	return regexp.MustCompile(",").ReplaceAllString(s, "")
}

func getPrecision(s string) int {
	matches := regexp.MustCompile("^-?[0-9,]*[.]([0-9]*)$").FindStringSubmatch(s)
	if len(matches) != 2 {
		return 0
	}
	return len(matches[1])
}
