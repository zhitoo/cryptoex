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

}
