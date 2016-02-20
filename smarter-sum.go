package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"regexp"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(smarterSum(string(bytes)))
}

func smarterSum(s string) string {
	foundNum := false
	var sum big.Float
	highestPrecision := 0
	usedDollarSign := false
	usedCommas := false

	s = regexp.MustCompile("[^-$0-9,.]+").ReplaceAllString(s, " ")
	for _, field := range strings.Fields(s) {
		justNum := regexp.MustCompile("[$,]").ReplaceAllString(field, "")
		float, _, err := big.ParseFloat(justNum, 10, 10000, big.ToNearestEven)
		if err != nil {
			continue
		}
		foundNum = true
		sum.Add(&sum, float)
		p := getPrecision(field)
		if p > highestPrecision {
			highestPrecision = p
		}
		if strings.Contains(field, "$") {
			usedDollarSign = true
		}
		if strings.Contains(field, ",") {
			usedCommas = true
		}
	}
	if !foundNum {
		return ""
	}
	if highestPrecision == 1 && usedDollarSign {
		highestPrecision = 2
	}
	prettyNum := sum.Text('f', highestPrecision)
	if usedCommas {
		prettyNum = addCommas(prettyNum)
	}
	if usedDollarSign {
		prettyNum = addDollarSign(prettyNum)
	}
	return prettyNum
}

func addCommas(s string) (pretty string) {
	end := len(s) - 1
	if strings.Contains(s, ".") {
		end = strings.Index(s, ".") - 1
		pretty = s[end+1:]
	}
	for i, grp := end, 0; i >= 0; i-- {
		if !strings.ContainsAny(string(s[i]), "1234567890") {
			pretty = s[:i+1] + pretty
			break
		}
		if grp == 3 {
			grp = 0
			pretty = "," + pretty
		}
		grp++
		pretty = string(s[i]) + pretty
	}
	return
}

func addDollarSign(s string) string {
	if len(s) == 0 {
		return ""
	}
	if string(s[0]) == "-" {
		return "-$" + s[1:]
	}
	return "$" + s
}

func getPrecision(s string) int {
	m := regexp.MustCompile("^-?[$]?[0-9,]*[.]([0-9]*)$").FindStringSubmatch(s)
	if len(m) != 2 {
		return 0
	}
	return len(m[1])
}
