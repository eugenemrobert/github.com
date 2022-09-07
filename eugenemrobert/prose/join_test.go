package prose

import (
	"testing"
)

type testData struct {
	phrases []string
	want    string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{phrases: []string{}, want: ""},
		testData{phrases: []string{"my parents"}, want: "my parents"},
		testData{phrases: []string{"my parents", "a rodeo clown"}, want: "my parents and a rodeo clown"},
		testData{phrases: []string{"my parents", "a rodeo clown", "a prize bull"}, want: "my parents, a rodeo clown, and a prize bull"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.phrases)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.phrases, got, test.want)
		}
	}
}
