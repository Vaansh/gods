package heap

type MinHeap struct {
	heap     []int
	size     int
	heapsize int
}

func NewMinHeap(h int) *MinHeap {
	minheap := &MinHeap{
		heap:     []int{},
		size:     0,
		heapsize: h,
	}
	return minheap
}

func (m *MinHeap) Insert(i int) {
	if m.size >= m.heapsize {
		return
	}
	m.heap = append(m.heap, i)
	m.size++
	m.UpHeapify(m.size - 1)
}

func (m *MinHeap) Remove() int {
	root := m.heap[0]
	m.heap[0] = m.heap[m.size-1]
	m.heap = m.heap[:m.size-1]
	m.size--
	m.DownHeapify(0)
	return root
}

func (m *MinHeap) UpHeapify(i int) {
	for m.heap[i] < m.heap[m.Parent(i)] {
		m.Swap(i, m.Parent(i))
	}
}

func (m *MinHeap) DownHeapify(i int) {
	if m.Leaf(i) {
		return
	}
	min := i
	left := m.Left(i)
	right := m.Right(i)
	if left < m.size && m.heap[left] < m.heap[min] {
		min = left
	}
	if right < m.size && m.heap[right] < m.heap[min] {
		min = right
	}
	if min != i {
		m.Swap(i, min)
		m.DownHeapify(min)
	}
}

func (m *MinHeap) ConstructMinHeap() {
	for i := ((m.size / 2) - 1); i >= 0; i-- {
		m.DownHeapify(i)
	}
}

func (m *MinHeap) Swap(a, b int) {
	temp := m.heap[a]
	m.heap[a] = m.heap[b]
	m.heap[b] = temp
}

func (m *MinHeap) Leaf(i int) bool {
	if i >= (m.size/2) && i <= m.size {
		return true
	}
	return false
}

func (m *MinHeap) Parent(i int) int {
	return (i - 1) / 2
}

func (m *MinHeap) Left(i int) int {
	return 2*i + 1
}

func (m *MinHeap) Right(i int) int {
	return 2*i + 2
}

//func main() {
//in := []int{6, 5, 3, 7, 2, 8}
//min := NewMinHeap(len(in))
//for i := 0; i < len(in); i++ {
//min.Insert(in[i])
//}
//min.ConstructMinHeap()
//for i := 0; i < len(in); i++ {
//fmt.Println(min.Remove())
//}
//}
