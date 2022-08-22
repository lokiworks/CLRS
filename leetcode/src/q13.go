package main

var (
	gMap = make(map[rune]int)
)

func init() {
	gMap['I'] = 1
	gMap['V'] = 5
	gMap['X'] = 10
	gMap['L'] = 50
	gMap['C'] = 100
	gMap['D'] = 500
	gMap['M'] = 1000

}

func romanToInt(s string) int {
	r := 0
	n := 0
	for _, e := range s {
		t, ok := gMap[e]
		if ok {
			if n < t {
				r = r - 2*n + t
			} else {
				r += t
			}
		}
		n = gMap[e]
	}

	return r
}
