package shopping_sim

type ProductItem struct {
	Code       string
	Name       string
	BasicPrice float32
	BulkPrice  float32
	priceRule
}

type ExpectedProduct struct {
	Code       string
	Name       string
	Num        int
	TotalPrice float32
}

type priceRule interface {

	//num, number of of current items
	//pass 1 if there is no promot code or discount
	//most of time return one product type, however 2G return mulitple types
	//so we return slice
	ApplyRule(num int, discount float32) []ExpectedProduct

	ProductName() string
}

var priceManager map[string]priceRule

func registerPriceRule(code string, r priceRule) {

	if priceManager == nil {
		priceManager = make(map[string]priceRule)
	}

	priceManager[code] = r
}
