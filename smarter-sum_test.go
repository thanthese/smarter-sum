package main

import "testing"

func TestSmarterSum(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"6", "6"},
		{"6 7", "13"},
		{"1 2 3 4 5", "15"},
		{"1.0 2 3 4 5", "15.0"},
		{"1.0 2 3 4 5.00", "15.00"},
		{"6 test 7", "13"},
		{"1,234 10", "1,244"},
		{"1234.12 10", "1244.12"},
		{"test 1234.12 10 test", "1244.12"},
		{"test 1,234.12 10 test", "1,244.12"},
		{"1\n2\n3\n4\n5", "15"},
		{"$15.12 text $100.10", "$115.22"},
		{"10 1", "11"},
		{"10. 1", "11"},
		{"10.0 1", "11.0"},
		{"10.00 1", "11.00"},
		{"10.000 1", "11.000"},
		{"1,000,000 222,000 333 0.444", "1,222,333.444"},
		{"1000000 222000 333 .444", "1222333.444"},
		{"1000 ,test, 2000", "3000"},

		{"-6", "-6"},
		{"-6 7", "1"},
		{"-1 2 3 4 5", "13"},
		{"-1.0 2 3 4 5", "13.0"},
		{"-1.0 2 3 4 5.00", "13.00"},
		{"-6 test 7", "1"},
		{"-1,234 10", "-1,224"},
		{"-1234.12 10", "-1224.12"},
		{"test -1,234.12 10 test", "-1,224.12"},
		{"test -1234.12 10 test", "-1224.12"},
		{"-1\n2\n3\n4\n5", "13"},
		{"-$15.12 text $100.10", "$84.98"},
		{"-10 1", "-9"},
		{"-10. 1", "-9"},
		{"-10.0 1", "-9.0"},
		{"-10.00 1", "-9.00"},
		{"-10.000 1", "-9.000"},
		{"-1,000,000 -222,000 -333 -0.444", "-1,222,333.444"},
		{"-1000000 -222000 -333 -.444", "-1222333.444"},
		{"-1000 ,test, -2000", "-3000"},
		{"-1000 ,test, 2000", "1000"},

		{"$1 2 3 4", "$10"},
		{"1 2 $3 4", "$10"},
		{"-$1 2 3 4", "$8"},
		{"-$10 2 3 4", "-$1"},
		{"-$10 2.8 3 4", "-$0.20"},
		{"-10 2.8 $3 4", "-$0.20"},
		{"10 -$20 -$30.1", "-$40.10"},
		{"10 -$20 -$30.12", "-$40.12"},
		{"10 -$20 -$30.123", "-$40.123"},
	}
	for _, c := range cases {
		got := smarterSum(c.in)
		if got != c.want {
			t.Errorf("smartSum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestGetPrecision(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"test", 0},
		{"10", 0},
		{"10.", 0},
		{"10.1", 1},
		{"10.12", 2},
		{"$10.12", 2},
		{"-$10.12", 2},
		{"10.113", 3},
		{"0.", 0},
		{"0.1", 1},
		{"0.12", 2},
		{"0.113", 3},
		{".", 0},
		{".1", 1},
		{".12", 2},
		{".113", 3},
		{".113.113", 0},
	}
	for _, c := range cases {
		got := getPrecision(c.in)
		if got != c.want {
			t.Errorf("smartSum(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestAddCommas(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
		{"1234567890", "1,234,567,890"},

		{"1.", "1."},
		{"12.1", "12.1"},
		{"123.12", "123.12"},
		{"1234.123", "1,234.123"},
		{"12345.123", "12,345.123"},
		{"123456.123", "123,456.123"},
		{"1234567.123", "1,234,567.123"},
		{"12345678.123", "12,345,678.123"},
		{"123456789.123", "123,456,789.123"},
		{"1234567890.123", "1,234,567,890.123"},

		{"-1", "-1"},
		{"-12", "-12"},
		{"-123", "-123"},
		{"-1234", "-1,234"},
		{"-12345", "-12,345"},
		{"-123456", "-123,456"},
		{"-1234567", "-1,234,567"},
		{"-12345678", "-12,345,678"},
		{"-123456789", "-123,456,789"},
		{"-1234567890", "-1,234,567,890"},

		{"1244.12", "1,244.12"},
	}
	for _, c := range cases {
		got := addCommas(c.in)
		if got != c.want {
			t.Errorf("smartSum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAddDollarSigns(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"1", "$1"},
		{"12", "$12"},
		{"123", "$123"},
		{"123.4", "$123.4"},
		{"123.45", "$123.45"},
		{"-1", "-$1"},
		{"-12", "-$12"},
		{"-123", "-$123"},
		{"-123.4", "-$123.4"},
		{"-123.45", "-$123.45"},
	}
	for _, c := range cases {
		got := addDollarSign(c.in)
		if got != c.want {
			t.Errorf("smartSum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
