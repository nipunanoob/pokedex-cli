package main

import "testing"

func TestCleanInput(t *testing.T){
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello world",
			expected: []string{"hello","world"},
		},
		{
			input: "bye word    ",
			expected: []string{"bye","word"},
		},
		{
			input: "hi to the greatest man alive",
			expected: []string{"hi","to","the","greatest","man","alive",},
		},
		{
			input: "",
			expected: []string{},
		},
	}

	for _,c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("Expected %d words but got %d",len(c.expected),len(actual))	
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word{
				t.Errorf("Word %s does not match expected %s",word,expectedWord)
			}
		}
	}
}