package main

type MaxHeap struct {
	slice []int
}

func (h *MaxHeap) insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)

}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// (index -1) /2 = parent index
func parent(i int) int {
	return (i - 1) / 2

}

// index * 2 + 1 = (left child index)
func left(i int) int {
	return i*2 + 1
}

// index * 2 + 1 = (right child index)
func right(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]
}

func main() {

}
