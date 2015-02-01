// Shipping Quotes
// Contains communication logic to handle
// shipping quotes negotiations with the server.
package printhouse

// Generates a shipping quote based on the information submitted
//
// If positive, the ShippingQuote will have a new Id.
// Usage:
//	1: ph.PostShippingQuote(&ShippingQuote{...})
func (ph *Printhouse) PostShippingQuote(sq *ShippingQuote) *ShippingQuote {
	return nil
}

// Lists all shipping quotes for the user
//
// Usage:
//	1: ph.GetShippingQuotes()
func (ph *Printhouse) GetShippingQuotes() []ShippingQuote {
	return nil
}

// Get's the specified shipping quote for the user
//
// Usage:
//	1: ph.GetShippingQuote("54808653b7017c7f795a6ac6")
func (ph *Printhouse) GetShippingQuote(id string) ShippingQuote {
	return nil
}



