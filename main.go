package main

func main() {
	minHeap := CreateHeap()
	minHeap.Insert(50)
	minHeap.Insert(30)
	minHeap.Insert(20)
	minHeap.Insert(15)
	minHeap.Insert(10)
	minHeap.Insert(8)
	minHeap.Remove()
}
