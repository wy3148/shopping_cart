package main

import "github.com/wy3148/shopping_cart/shopping_sim"

func main() {

	cart := shopping_sim.NewShoppingCart()

	// ult_small Unlimited 1GB $24.90
	// ult_medium Unlimited 2GB $29.90
	// ult_large Unlimited 5GB $44.90
	// 1gb 1 GB Data-pack $9.90

	cart.Add("ult_small", 1)
	cart.Add("ult_medium", 1)
	cart.Add("ult_large", 4)
	cart.Add("1gb", 1, "I<3AMAYSIM")
	cart.PrintOrder()
}
