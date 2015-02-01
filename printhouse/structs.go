// Structs to be used on this library,
// contains analog models to the ones
// on api server side.
package printhouse

type Product struct{
	Id          string
	Kind        string
	Description string
	Price       float64
	XInches     int16
	YInches     int16
	XPixels     int16
	YPixels     int16
	Format      string
}

type PrintFile struct{
	Product  string
	FileUrl  string
	Id       string
}

type Order struct{

}
type ShippingQuote struct{

}
