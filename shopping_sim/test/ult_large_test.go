package shopping_sim_test

import (
	"fmt"
	"github.com/wy3148/shopping_cart/shopping_sim"
	"testing"
)

func Test5G(t *testing.T) {

	s := &shopping_sim.UltLarge{}
	s.Code = "ult_large"
	s.Name = "Unlimited 5GB"
	s.BasicPrice = 44.90
	s.BulkPrice = 39.90

	res := s.ApplyRule(10, 1)

	//buy 10 X 5G : output 10 X 5G
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
	total := float32(10 * 39.90)
	realTotal := float32(0.00)

	for _, v := range res {
		realTotal += v.TotalPrice
	}

	if total != realTotal {
		t.Error("wrong total price: %f, %f", total, realTotal)
		t.FailNow()
		return
	}

	// Output:
	// ok
	fmt.Println("ok")
}
