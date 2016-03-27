// Products
//
// Contains communication logic to handle
// products negotiations with the server.
package printhouse

// Returns the list of products available
// through Printhouse (canvas, photo paper,
// posters, etc. in all their sizes).
//
// Usage:
//	1: ph.GetProducts(nil): Retrieve all
// products.
//	2: ph.GetProducts("kind=canvas&format=4x3"):
// Retrieve products filtered by optional
// parameters such as kind, size, and format
func (ph *Printhouse) GetProducts(filters string) []Product {
	return nil
}

// Returns a single products' full information.
//
// Usage:
//	1: ph.GetProduct("54808653b7017c7f795a6ac6")
func (ph *Printhouse) GetProduct(id string) *Product {
	return nil
}


