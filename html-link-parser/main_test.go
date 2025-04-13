package main

import (
	"fmt"
	"html-link-parser/parser"
	"html-link-parser/utils"
	"strings"
	"testing"
)

type testCase struct {
	path     string
	expected string
}

func TestExamples(t *testing.T) {
	testCases := []testCase{
		{
			path:     "ex.html",
			expected: "{\n\tHref: \"/dog\",\n\tText: \"Something in a span Text not in a span Bold text!\",\n},",
		},
		{
			path:     "ex1.html",
			expected: "{\n\tHref: \"/other-page\",\n\tText: \"A link to another page\",\n},",
		},
		{
			path:     "ex2.html",
			expected: "{\n\tHref: \"https://www.twitter.com/joncalhoun\",\n\tText: \"Check me out on twitter\",\n},\n{\n\tHref: \"https://github.com/gophercises\",\n\tText: \"Gophercises is on Github !\",\n},",
		},
		{
			path:     "ex3.html",
			expected: "{\n\tHref: \"#\",\n\tText: \"Login\",\n},\n{\n\tHref: \"/lost\",\n\tText: \"Lost? Need help?\",\n},\n{\n\tHref: \"https://twitter.com/marcusolsson\",\n\tText: \"@marcusolsson\",\n},",
		},
		{
			path:     "ex4.html",
			expected: "{\n\tHref: \"/dog-cat\",\n\tText: \"dog cat\",\n},",
		},
	}

	for _, testC := range testCases {
		res := parser.Parse(testC.path)
		stringRes := make([]string, len(*res))

		for i, n := range *res {
			stringRes[i] += n.String()
		}
		strRes := strings.Join(stringRes, "\n")

		if strRes != testC.expected {
			t.Errorf("\nParse(%q)\n\tGot      |  %q\n\tExpected |  %q\n", testC.path, strRes, testC.expected)
		} else {
			fmt.Printf("\nParse(%q)\n\tGot      |  %q\n\tExpected |  %q\n", testC.path, strRes, testC.expected)
		}
	}
}

func TestCorrectPaths(t *testing.T) {
	testCases := []string{
		"ex.html",
		"ex1.html",
		"ex2.html",
		"ex3.html",
		"ex4.html",
	}

	for _, testC := range testCases {
		_, err := utils.ReaderFromFile(testC)

		if err != nil {
			t.Errorf("\nReaderFromFile(%q)\n\tGot      |  %q\n\tExpected |  nil\n", testC, err)
		} else {
			fmt.Printf("\nReaderFromFile(%q)\n\tGot      |  %q\n\tExpected |  nil\n", testC, err)
		}
	}
}

func TestIncorrectPaths(t *testing.T) {
	testCases := []string{
		"ex5.html",
		"example.html",
		"resources",
		"res/ex.html",
		"ex.xml",
		"ex.1html",
	}

	for _, testC := range testCases {
		_, err := utils.ReaderFromFile(testC)

		if err == nil {
			t.Errorf("\nReaderFromFile(%q)\n\tGot      |  %q\n\tExpected |  \"open res/%s: The system cannot find the file specified.\"\n", testC, err, testC)
		} else {
			fmt.Printf("\nReaderFromFile(%q)\n\tGot      |  %q\n\tExpected |  \"open res/%s: The system cannot find the file specified.\"\n", testC, err, testC)
		}
	}
}
