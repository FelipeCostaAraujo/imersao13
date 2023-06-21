package entity

type OrderQueue struct {
	Orders []*Order
}

// Less is a function that compares two orders
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap is a function that swaps two orders
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len is a function that returns the length of the order queue
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push is a function that adds an order to the order queue
func (oq *OrderQueue) Push(x interface{}) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Pop is a function that removes an order from the order queue
func (oq *OrderQueue) Pop() interface{} {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
