package main

var (
	gMap1 = make(map[byte]int)
)

func init() {
	// init gMap1
	//---------------
	gMap1['('] = -1
	gMap1[')'] = 1
	gMap1['{'] = -10
	gMap1['}'] = 10
	gMap1['['] = -20
	gMap1[']'] = 20
	gMap1[' '] = 0
	//---------------

}

func isValid(s string) bool {
	r := 0
	l := []int{}

	for _, e := range s {
		if r > 0 {
			return false
		}

		v := gMap1[byte(e)]
		if v < 0 {
			l = append(l, v)
		} else {
			if len(l)-1 >= 0 {
				prev := l[len(l)-1]
				if prev+v != 0 {
					return false
				}
				l = l[:len(l)-1]
			}

		}

		r += v

	}
	return r == 0
}
