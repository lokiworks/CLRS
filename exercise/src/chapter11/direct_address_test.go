package main

import (
	"testing"
)

func Test_directAddressSearch(t *testing.T) {
	arr := make([]DirectAddressObject, 10)
	o1 := DirectAddressObject{k: 0, v: "0"}
	o2 := DirectAddressObject{k: 1, v: "1"}
	o3 := DirectAddressObject{k: 2, v: "2"}

	directAddressInsert(arr, o1)
	directAddressInsert(arr, o2)
	directAddressInsert(arr, o3)

	act01 := directAddressSearch(arr, 0)
	if act01.k != 0 {
		t.Errorf("search() = %v, want %v", act01.k, 0)
	}

	directAddressDelete(arr, o1)
	directAddressDelete(arr, o2)
	directAddressDelete(arr, o3)

}
