package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	var err error = nil
	var dis int = 0
	if len(a) != len(b) {
		return 0, fmt.Errorf("length mismatch")
	}
	for i, _ := range a{
		elemA := a[i]
		elemB := b[i]
		if elemA != elemB {
			dis++
		}
		i++
	}
	return dis, err
}
