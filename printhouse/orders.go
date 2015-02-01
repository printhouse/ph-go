// Orders
// Contains communication logic to handle
// orders negotiations with the server.
package printhouse

// Lists all the orders in the account.
//
// Usage:
//	1: ph.GetOrders()
func (ph *Printhouse) GetOrders() []Order {
	return nil
}

// Returns a single order's information
//
// Usage:
//	1: ph.GetOrder("54808653b7017c7f795a6ac6")
func (ph *Printhouse) GetOrder(id string) Order {
	return nil
}

// Creates an order, with the inner order_items needed.
//
// If positive, the Order will have a new Id.
// Usage:
//	1: ph.PostOrder(&Order{...})
func (ph *Printhouse) PostOrder(order *Order) *Order {
	return nil
}

// Updates an order, if it hasn't been confirmed.
//
// Usage:
//	1: ph.PutOrder(&Order{...})
func (ph *Printhouse) PutOrder(order *Order) (*Order, error) {
	return nil, nil
}

// Charges the developer's account and sends
// an order to production.
//
// An order cannot be updated once it has
// been confirmed
//
// Usage:
//	1: ph.ConfirmOrder(&Order{...})
func (ph *Printhouse) ConfirmOrder(order *Order) bool {
	return false
}

