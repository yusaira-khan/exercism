package hamming

func Distance(a, b string) (int, error) {
	var err error = nil
	var dis int = 0
	if len(a) == 1 {
		if a[0] != b[0] {
			dis = 1
		}
	}
	return dis, err
}
