package shopping_sim

type UltSmall ProductItem

func init() {
	var s UltSmall
	s.Code = "ult_small"
	s.Name = "Unlimited 1GB"
	s.BasicPrice = 24.9
	registerPriceRule(s.Code, &s)
}

func (s *UltSmall) ApplyRule(num int, discount float32) []ExpectedProduct {

	var expected []ExpectedProduct
	var t ExpectedProduct

	t.Code = s.Code
	t.Name = s.Name
	t.Num = num
	t.TotalPrice = float32(num-num/3) * s.BasicPrice * discount

	return append(expected, t)
}

func (s *UltSmall) ProductName() string {
	return s.Name
}
