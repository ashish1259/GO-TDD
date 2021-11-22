package refleection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	city string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"chris"},
			[]string{"chris"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
