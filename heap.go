package algorithmsingo
import (
	"math"
)

type Heap struct {
	items []int
}

func NewHeap() *Heap {
	return &Heap{
		items: []int{},
	}
}

func (heap *Heap) Add(item int) {
	heap.items = append(heap.items, item)
	currentIndex := len(heap.items) - 1
	heap.bubbleUp(currentIndex)
}

func (heap *Heap) Get() int {
	minItem := heap.Peak()
	heapSize := heap.Size()
	heap.items[0] = heap.items[heapSize - 1]
	heap.items = heap.items[:heapSize - 1]

	if (heapSize - 1 > 1) {
		heap.bubbleDown(0)
	}
	return minItem
}

func (heap *Heap) Peak() int {
	return heap.items[0]
}

func (heap *Heap) Size() int {
	return len(heap.items)
}

func (heap *Heap) bubbleUp(currentIndex int) {
	parentIndex := getParentIndex(currentIndex)
	if currentIndex != 0 && heap.items[parentIndex] > heap.items[currentIndex] {
		parentItem := heap.items[parentIndex]
		heap.items[parentIndex] = heap.items[currentIndex]
		heap.items[currentIndex] = parentItem
		heap.bubbleUp(parentIndex)
	}
}

func (heap *Heap) bubbleDown(currentIndex int) {
	heapSize := heap.Size()
	currentItem := heap.items[currentIndex]
	leftValue := math.MaxInt64
	rightValue := math.MaxInt64

	leftChildIndex, rightChildIndex := getChildIndexes(currentIndex)
	if leftChildIndex < heapSize {
		leftValue = heap.items[leftChildIndex]
	}
	if rightChildIndex < heapSize {
		rightValue = heap.items[rightChildIndex]
	}

	isRightMin := rightValue < leftValue

	if isRightMin && rightChildIndex < heapSize {
		heap.items[rightChildIndex] = currentItem
		heap.items[currentIndex] = rightValue
		heap.bubbleDown(rightChildIndex)
	} else if !isRightMin && leftChildIndex < heapSize {
		heap.items[leftChildIndex] = currentItem
		heap.items[currentIndex] = leftValue
		heap.bubbleDown(leftChildIndex)
	}
}

func getParentIndex(index int) int {
	return int(math.Ceil(float64(index) / 2.0)) - 1
}

func getChildIndexes(index int) (int, int) {
	return index * 2 + 1, index * 2 + 2
}