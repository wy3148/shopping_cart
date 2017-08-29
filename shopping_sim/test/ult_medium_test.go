package shopping_sim_test

import (
	"fmt"
	"github.com/wy3148/shopping_cart/shopping_sim"
	"testing"
)

func Test2G(t *testing.T) {

	s := &shopping_sim.UltMedium{}
	s.Code = "ult_medium"
	s.Name = "Unlimited 2GB"
	s.BasicPrice = 29.90

	res := s.ApplyRule(10, 0.9)

	//get 1GB pack for free!
	if len(res) != 2 {
		t.Error("wrong type of expected items")
		t.FailNow()
		return
	}

	//expected number
	total := float32(10 * 29.90 * 0.9)
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
