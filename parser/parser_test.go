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
			Input:  `{"name":  123 }`,
			Expect: []string{"{", "name", ":", "123", "}"},
		},
		{
			Input:  `{"name":"zhangsan"}`,
			Expect: []string{"{", "name", ":", "zhangsan", "}"},
		},
		{
			Input:  `{"name":null}`,
			Expect: []string{"{", "name", ":", "null", "}"},
		},
	}

	for _, item := range testTable {
		l := NewLex(item.Input)
		token := l.NextToken()
		t.Logf("token %s", token)
		for token != EOF {
			token = l.NextToken()
			t.Logf("token %s", token)
		}
	}
}

func TestString(t *testing.T) {
	token := "name:null"

	t.Logf("0-3: %s", token[0:4])

	if token[0:4] != "name" {
		t.Error("error")
	}
}
