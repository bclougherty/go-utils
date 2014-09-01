package stringutils

import (
	"fmt"
	"testing"
)

type StringPair struct {
    A, B string
}

func TestCamelToSpinal(t *testing.T) {
	fmt.Print("TestCamelToSpinal:\n")

	tests := make([]StringPair, 1)

	tests = append(tests, StringPair{
		"TestCamelCaseStringWithLotsOfWordsOfVaryingLengths",
		"test-camel-case-string-with-lots-of-words-of-varying-lengths",
	})

	tests = append(tests, StringPair{
		"ALLCAPS",
		"a-l-l-c-a-p-s",
	})

	tests = append(tests, StringPair{
		"alllowercase",
		"alllowercase",
	})


	for i := 0; i < len(tests); i++ {
		test := tests[i]

		fmt.Printf("\t\"%s\" should become \"%s\"\n", test.A, test.B)

		output := CamelToSpinal(test.A)

		if output != test.B {
			t.Errorf("Expected %s, got %s", test.B, output)
		}
	}

	fmt.Print("\n")
}

func TestSpinalToCamel(t *testing.T) {
	fmt.Print("TestSpinalToCamel:\n")

	tests := make([]StringPair, 1)

	tests = append(tests, StringPair{
		"test-camel-case-string-with-lots-of-words-of-varying-lengths",
		"TestCamelCaseStringWithLotsOfWordsOfVaryingLengths",
	})

	tests = append(tests, StringPair{
		"a-l-l-c-a-p-s",
		"ALLCAPS",
	})

	tests = append(tests, StringPair{
		"alloneword",
		"Alloneword"})


	for i := 0; i < len(tests); i++ {
		test := tests[i]

		fmt.Printf("\t\"%s\" should become \"%s\"\n", test.A, test.B)

		output := SpinalToCamel(test.A)

		if output != test.B {
			t.Errorf("Expected %s, got %s", test.B, output)
		}
	}

	fmt.Print("\n")
}