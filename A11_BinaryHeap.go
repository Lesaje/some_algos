package binaryHeap

type Heap struct {
    slice []int
}

func (h *Heap) siftDown(cur int) {
    for 2*cur+1 < len(h.slice) {
        left := 2*cur + 1
        right := 2*cur + 2
        min := left
        if right < len(h.slice) && h.slice[right] < h.slice[left] {
            min = right
        }
        if h.slice[cur] <= h.slice[min] {
            break
        }
        h.slice[cur], h.slice[min] = h.slice[min], h.slice[cur]
        cur = min
    }
}

func (h *Heap) siftUp(cur int) {
    parent := (cur - 1) / 2
    for h.slice[parent] > h.slice[cur] {
        h.slice[parent], h.slice[cur] = h.slice[cur], h.slice[parent]
        cur = parent
        if cur == 0 {
            break
        }
        parent = (cur - 1) / 2
    }
}

//Heapify example of use:
// var h binaryHeap.Heap
// h.Heapify([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
// fmt.Println(h.ExtractMin())
func (h *Heap) Heapify(slice []int) {
    h.slice = slice
    //copy(h.slice, slice)
    middle := len(slice) - 2/2
    for i := middle; i > 0; i-- {
        h.siftDown(i)
    }
}

func (h *Heap) Insert(element int) {
    h.slice = append(h.slice, element)
    h.siftUp(len(h.slice) - 1)
}

//ExtractMin extracts and DELETES min value in the heap
func (h *Heap) ExtractMin() int {
    min := h.slice[0]
    h.slice[0] = h.slice[len(h.slice)-1]
    if len(h.slice) > 1 {
        h.slice = h.slice[:len(h.slice)-1]
        h.siftDown(0)
    } else {
        h.slice = []int{}
    }
    return min
}
