package shopping_sim_test

import (
	"fmt"
	"github.com/wy3148/shopping_cart/shopping_sim"
	"testing"
)

func Test1gb(t *testing.T) {

	s := &shopping_sim.OneGbPack{}
	s.Code = shopping_sim.OneGbProductCode
	s.Name = shopping_sim.OneGbProductName
	s.BasicPrice = 9.90

	res := s.ApplyRule(10, 0.9)

	if len(res) != 1 {
		t.Error("wrong type of expected items")
		t.FailNow()
		return
	}

	if res[0].Num != 10 {
		t.Error("wrong number of expected items")
		t.FailNow()
	}

	//expected number
	total := float32(10 * 9.90 * 0.9)
	realTotal := float32(0.00)

	for _, v := range res {
		realTotal += v.TotalPrice
	}

	if total != realTotal {
		t.Error("wrong total price:", total, realTotal)
		t.FailNow()
		return
	}

	// Output:
	// ok
	fmt.Println("ok")
}
