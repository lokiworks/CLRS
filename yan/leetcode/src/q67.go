package main

import "strconv"

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	remainder := 0
	result := ""
	i, j := aLen-1, bLen-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		av, bv := int(a[i]-'0'), int(b[j]-'0')
		temp := (av + bv + remainder) % 2
		remainder = (av + bv + remainder) / 2
		result = result + strconv.Itoa(temp)
	}
	for ; i >= 0; i = i - 1 {
		temp := (int(a[i]-'0') + remainder) % 2
		remainder = (int(a[i]-'0') + remainder) / 2
		result = result + strconv.Itoa(temp)
	}
	for ; j >= 0; j = j - 1 {
		temp := (int(b[j]-'0') + remainder) % 2
		remainder = (int(b[j]-'0') + remainder) / 2
		result = result + strconv.Itoa(temp)
	}

	if remainder != 0 {
		result = result + strconv.Itoa(remainder)
	}

	reverseResult := ""

	for _, v := range result {
		reverseResult = string(v) + reverseResult
	}
	return reverseResult

}
