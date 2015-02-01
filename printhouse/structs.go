// Structs to be used on this library,
// contains analog models to the ones
// on api server side.
package printhouse

type Product struct{
	id          string
	kind        string
	description string
	price       float64
	x_inches    int16
	y_inches    int16
	x_pixels    int16
	y_pixels    int16
	format      string
}

type PrintFile struct{
	product  string
	file_url string
	id       string
}

type Order struct{

}
type ShippingQuote struct{

}
