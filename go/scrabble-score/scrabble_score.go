package scrabble

import "strings"

var copied = map[string]int{"A, E, I, O, U, L, N, R, S, T": 1,
	"D, G":                               2,
	"B, C, M, P":                         3,
	"F, H, V, W, Y":                       4,
	"K":                                  5,
	"J, X" :                              8,
	"Q, Z":                               10,
}
var Letters = getScoresTable()
func Score(word string) int{
	currentScore:=0
	word=strings.ToLower(word)
	for _,letter :=range word{
		currentScore+=Letters[letter]
	}
	return currentScore
}
func firstRune(str string) (r rune) {
	for _, r = range str {
		return
	}
	return
}
func toRunes(sL []string)[]rune{
	rL := make([]rune, len(sL))
    for i,s := range sL{
    	rL[i]=firstRune(s)
    }
    return rL
}

func Split(m string)[]rune{
  return toRunes(strings.Split(strings.ToLower(m),", "))
}

func getScoresTable()map[rune]int {
	l := make(map[rune]int,1)
	for letters,score := range copied {
		ll:=Split(letters)
		for _,letter := range ll{
			l[letter] = score
		}
	}
	return l
}