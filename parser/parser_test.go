package parser

import "testing"

func TestSwitch(t *testing.T) {

	strs := []string{"1", "2", "3"}

	for _, item := range strs {
		switch item {

		case "1", "2", "3":
			t.Logf("value: %s", item)
		}
	}
}

func TestLexer(t *testing.T) {
	testTable := []struct {
		Input  string
		Expect []string
	}{
		{
			Input:  `{"name":"zhangsan"}`,
			Expect: []string{"{", "name", ":", "zhangsan", "}"},
		},
		{
			Input:  `{"name":null}`,
			Expect: []string{"{", "name", ":", "null", "}"},
		},
		{
			Input:  `{"name":123}`,
			Expect: []string{"{", "name", ":", "123", "}"},
		},
	}

	for _, item := range testTable {
		l := NewLex(item.Input)
		token := l.NextToken()

		for token != EOF {
			t.Logf("token %s", token)
			token = l.NextToken()
		}
	}
}
