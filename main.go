package main

import (
	"github.com/wy3148/shopping_cart/shopping_sim"
	"log"
)

func main() {

	// ult_small Unlimited 1GB $24.90
	// ult_medium Unlimited 2GB $29.90
	// ult_large Unlimited 5GB $44.90
	// 1gb 1 GB Data-pack $9.90

	log.Println("Scenario 1:")
	cart1 := shopping_sim.NewShoppingCart()
	cart1.Add("ult_small", 3)
	cart1.Add("ult_large", 1)
	cart1.PrintOrder()
	log.Print("\n\n")

	log.Println("Scenario 2:")
	cart2 := shopping_sim.NewShoppingCart()
	cart2.Add("ult_small", 2)
	cart2.Add("ult_large", 4)
	cart2.PrintOrder()
	log.Print("\n\n")

	log.Println("Scenario 3:")
	cart3 := shopping_sim.NewShoppingCart()
	cart3.Add("ult_small", 1)
	cart3.Add("ult_medium", 2)
	cart3.PrintOrder()
	log.Print("\n\n")

	log.Print("Scenario 4:")
	cart4 := shopping_sim.NewShoppingCart()
	cart4.Add("ult_small", 1, "I<3AMAYSIM")
	cart4.Add("1gb", 1)
	cart4.PrintOrder()

}
