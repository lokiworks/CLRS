package main

type Queue []int

func (s *Queue) Enqueue(x int) {
	*s = append(*s, x)
}

func (s *Queue) Dequeue() int {
	if len(*s) == 0 {
		return 0
	}

	x := (*s)[0]
	*s = (*s)[1:len(*s)]
	return x
}

func main() {

}
