package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}
type bucket struct {
	head *bucketNode
}
type bucketNode struct {
	key  int
	next *bucketNode
}

func (h *HashTable) Insert(key int) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key int) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key int) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(key int) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	}
}

func (b *bucket) search(key int) bool {
	currentNode := b.head
	for ; currentNode != nil; currentNode = currentNode.next {
		if currentNode.key == key {
			return true
		}
	}

	return false
}

func (b *bucket) delete(key int) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func hash(key int) int {
	return key % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
func main() {
	testHashTable := Init()
	testHashTable.Insert(1)
	testHashTable.Insert(2)
	testHashTable.Insert(3)
	fmt.Println(testHashTable.Search(1))
	fmt.Println(testHashTable.Search(2))
	fmt.Println(testHashTable.Search(3))

	testHashTable.Delete(1)
	testHashTable.Delete(2)
	testHashTable.Delete(3)

	fmt.Println(testHashTable.Search(1))
	fmt.Println(testHashTable.Search(2))
	fmt.Println(testHashTable.Search(3))
}
