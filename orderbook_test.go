package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrder1 := NewOrder(true, 1)
	buyOrder2 := NewOrder(true, 2)
	buyOrder3 := NewOrder(true, 3)

	l.AddOrder(buyOrder1)
	l.AddOrder(buyOrder2)
	l.AddOrder(buyOrder3)

	l.DeleteOrder(buyOrder1)

	fmt.Println(l)
}

func TestOrderBook(t *testing.T) {
	ob := NewOrderbook()

	buyOrderA := NewOrder(true, 10)
	buyOrderB := NewOrder(true, 12)
	buyOrderC := NewOrder(true, 169)

	ob.PlaceOrder(18_000, buyOrderA)
	ob.PlaceOrder(19_000, buyOrderB)
	ob.PlaceOrder(20_000, buyOrderC)

	//fmt.Printf("%+v", ob)
	for i := 0; i < len(ob.Bids); i++ {
		fmt.Println(ob.Bids[i])
	}
}
