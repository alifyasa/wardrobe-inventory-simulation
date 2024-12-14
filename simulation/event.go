package simulation

type Event struct {
	Hour      int
	EventFunc func()
}

type EventQueue []*Event

func (eq EventQueue) Len() int {
	return len(eq)
}

func (eq EventQueue) Less(i int, j int) bool {
	return eq[i].Hour < eq[j].Hour
}

func (eq EventQueue) Swap(i int, j int) {
	eq[i], eq[j] = eq[j], eq[i]
}

func (pq *EventQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Event))
}

func (pq *EventQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
