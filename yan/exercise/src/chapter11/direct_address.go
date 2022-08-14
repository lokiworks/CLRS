package main

type DirectAddressObject struct {
	k int
	v string
}

func directAddressSearch(T []DirectAddressObject, k int) DirectAddressObject {
	return T[k]
}

func directAddressInsert(T []DirectAddressObject, x DirectAddressObject) {
	T[x.k] = x
}

func directAddressDelete(T []DirectAddressObject, x DirectAddressObject) {
	T[x.k] = DirectAddressObject{}
}

func main() {

}
