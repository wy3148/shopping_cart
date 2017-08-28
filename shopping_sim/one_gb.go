package shopping_sim

type OneGbPack ProductItem

const (
	OneGbProductCode = "1gb"
	OneGbProductName = "1 GB Data-pack"
)

func init() {
	var s OneGbPack
	s.Code = OneGbProductCode
	s.Name = OneGbProductName
	s.BasicPrice = 9.90
	registerPriceRule(s.Code, &s)
}

func (s *OneGbPack) ApplyRule(num int, discount float32) []ExpectedProduct {
	var expected []ExpectedProduct
	var t ExpectedProduct

	t.Code = s.Code
	t.Name = s.Name
	t.Num = num
	t.TotalPrice = s.BasicPrice * float32(num) * discount
	return append(expected, t)
}

func (s *OneGbPack) ProductName() string {
	return s.Name
}
