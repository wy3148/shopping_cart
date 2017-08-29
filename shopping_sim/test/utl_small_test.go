package shopping_sim_test

import (
	"fmt"
	"github.com/wy3148/shopping_cart/shopping_sim"
	"testing"
)

func Test1GSmall(t *testing.T) {

	s := &shopping_sim.UltSmall{}
	s.Code = "ult_small"
	s.Name = "Unlimited 1GB"
	s.BasicPrice = 24.90

	res := s.ApplyRule(10, 1)

	if len(res) != 1 {
		t.Error("wrong type of expected items")
		t.FailNow()
		return
	}

	//expected number
	total := float32(7 * 24.90)
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
