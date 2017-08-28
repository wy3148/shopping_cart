package shopping_sim

type UltMedium ProductItem

func init() {
	var s UltMedium
	s.Code = "ult_medium"
	s.Name = "Unlimited 2GB"
	s.BasicPrice = 29.90
	registerPriceRule(s.Code, &s)
}

func (s *UltMedium) ApplyRule(num int, discount float32) []ExpectedProduct {

	var expected []ExpectedProduct
	var m ExpectedProduct
	var oneGbFree ExpectedProduct

	m.TotalPrice = s.BasicPrice * float32(num) * discount
	m.Code = s.Code
	m.Name = s.Name
	m.Num = num
	expected = append(expected, m)

	oneGbFree.Code = OneGbProductCode
	oneGbFree.Name = OneGbProductName
	oneGbFree.TotalPrice = 0
	oneGbFree.Num = num
	expected = append(expected, oneGbFree)
	return expected
}

func (s *UltMedium) ProductName() string {
	return s.Name
}
