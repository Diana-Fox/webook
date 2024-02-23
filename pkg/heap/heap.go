package heap

//
//type HeapNode struct {
//	Id int
//}
//
//// type Heap map[int]any
//// *MaxHeap 来个大顶堆
//type MaxHeap[T HeapNode] struct {
//	nodes []HeapNode
//}
//
//func NewMaxHeap(nodes []HeapNode) MaxHeap {
//	return MaxHeap{
//		nodes: nodes,
//	}
//}
//
//func (m *MaxHeap[T]) Len() int {
//	return len(m.nodes)
//}
//
//func (m *MaxHeap[T]) Less(i, j int) bool {
//	return m.nodes[i].Id > m.nodes[j].Id
//}
//
//func (m *MaxHeap[T]) Swap(i, j int) {
//	m.nodes[i], m.nodes[j] = m.nodes[j], m.nodes[i]
//}
//
//func (m *MaxHeap[T]) Push(x any) {
//	m.nodes:=append(m.nodes, T.(x))
//}
//
//func (m *MaxHeap[T]) Pop() any {
//	v:=(m.nodes)[m.Len()-1]
//	m.nodes = m.nodes[:m.Len()-1]
//	return v
//}
//
////func (m *MaxHeap) Len() int {
////	return len(m.nodes)
////}
////
////// Less 大顶堆
////func (m *MaxHeap) Less(i, j int) bool {
////	return m.nodes[i].Id > m.nodes[j].Id
////}
////
////func (m *MaxHeap) Swap(i, j int) {
////	m.nodes[i], m.nodes[j] = m.nodes[j], m.nodes[i]
////}
////
////// 送入
////func (m *MaxHeap) Push(node *HeapNode) {
////	m.nodes = append(m.nodes, node)
////}
////
////// 弹出
////func (m *MaxHeap) Pop() *HeapNode {
////	node := m.nodes[m.Len()-1]
////	m.nodes = m.nodes[:m.Len()-1]
////	return node
////}
