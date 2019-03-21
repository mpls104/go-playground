package cart

import "testing"

func TestAddAndGetProductsInCart(t *testing.T) {
	c := New()
	c.Add("りんご")
	c.Add("みかん")

	products := c.GetAll()
	if len(products) != 2 {
		t.Fatalf("商品の数が想定と違う")
	}
}