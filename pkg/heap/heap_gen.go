package heap

//
//type HeapNodeGen struct {
//	Id int
//}
//
//// *MaxHeap 来个大顶堆
//type MaxHeapGen[T HeapNodeGen] struct {
//	nodes []T
//}
//
//func NewMaxHeapGen[T HeapNodeGen](nodes []T) MaxHeapGen[T] {
//	return MaxHeapGen[T]{
//		nodes: nodes,
//	}
//}
//func (m MaxHeapGen[T]) Len() int {
//	return len(m.nodes)
//}
////本来准备直接扩展，但是发现有点问题
//func (m MaxHeapGen[T]) Less(i, j int) bool {
//	return m.nodes[i].Id > m.nodes[j].Id
//}
//
//func (m MaxHeapGen[T]) Swap(i, j int) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m MaxHeapGen[T]) Push(x any) {
//
//	//TODO implement me
//	panic("implement me")
//}
//
//func (m MaxHeapGen[T]) Pop() any {
//	//TODO implement me
//	panic("implement me")
//}
