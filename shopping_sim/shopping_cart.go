package shopping_sim

import (
	"errors"
	"log"
)

const (
	promoCodeDef = "I<3AMAYSIM"
)

type ShoppingCart struct {
	ordered            map[string]int
	total              float32
	outputOrderedItems []ExpectedProduct
	discount           float32
}

func NewShoppingCart() *ShoppingCart {
	c := new(ShoppingCart)
	c.total = 0
	c.ordered = make(map[string]int)
	c.outputOrderedItems = nil
	c.discount = 1
	return c
}

func (c *ShoppingCart) Add(productCode string, num int, promoCode ...string) error {

	if num <= 0 {
		return errors.New("number of items can not be less than zero:")
	}

	if _, ok := priceManager[productCode]; !ok {
		return errors.New("failed to find the product with code:" + productCode)
	}

	if promoCode != nil && promoCode[0] == promoCodeDef {
		c.discount = 0.9
	}

	if _, ok := c.ordered[productCode]; ok {
		c.ordered[productCode] = c.ordered[productCode] + num
	} else {
		c.ordered[productCode] = num
	}

	return nil
}

func (c *ShoppingCart) PrintOrder() {

	log.Println("Items Added 		Expected Cart Total		Expected Cart Items")
	for productCode, num := range c.ordered {
		if rule, ok := priceManager[productCode]; ok {

			log.Println(num, " x ", rule.ProductName())

			res := rule.ApplyRule(num, c.discount)

			for _, p := range res {
				c.total += p.TotalPrice
				c.outputOrderedItems = append(c.outputOrderedItems, p)
			}
		}
	}

	log.Printf("					%0.2f", c.total)

	for _, p := range c.outputOrderedItems {
		log.Println("								", p.Num, " x "+p.Name)
	}
}
