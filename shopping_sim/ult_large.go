package shopping_sim

type UltLarge ProductItem

func init() {
	var s UltLarge
	s.Code = "ult_large"
	s.Name = "Unlimited 5GB"
	s.BasicPrice = 44.90
	s.BulkPrice = 39.90
	registerPriceRule(s.Code, &s)
}

// The Unlimited 5GB Sim will have a bulk discount applied; whereby the price will drop to $39.90 each for the first month, if the customer
// buys more than 3.
func (s *UltLarge) ApplyRule(num int, discount float32) []ExpectedProduct {

	var expected []ExpectedProduct
	var t ExpectedProduct

	if num > 3 {
		t.TotalPrice = s.BulkPrice * float32(num) * discount
	} else {
		t.TotalPrice = s.BasicPrice * float32(num) * discount
	}

	t.Code = s.Code
	t.Name = s.Name
	t.Num = num
	return append(expected, t)
}

func (s *UltLarge) ProductName() string {
	return s.Name
}
