package scrabble

import (
	"fmt"
	"reflect"
	"testing"
)

func TestScore(t *testing.T) {
	for _, test := range scrabbleScoreTests {
		if actual := Score(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %d, Actual %d", test.input, test.expected, actual)
		} else {
			fmt.Printf("PASS: Score(%q) %d\n", test.input, actual)
		}
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func TestSplit(t *testing.T){
	var given   string="D, G"
	var expected = []rune{'d','g'}
	var  actual = Split(given)
	fmt.Printf("Split(%q) expected %d, Actual %d", given, expected, actual)
	if !reflect.DeepEqual(expected , actual ){
		t.Errorf("split(%q) expected %d, Actual %d", given, expected, actual)

	}


}
