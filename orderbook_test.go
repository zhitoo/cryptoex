package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	var buyOrder1, buyOrder2, buyOrder3 *Order
	buyOrder1 = NewOrder(true, 1)
	buyOrder2 = NewOrder(true, 2)
	buyOrder3 = NewOrder(true, 3)

	l.AddOrder(buyOrder1)
	l.AddOrder(buyOrder2)
	l.AddOrder(buyOrder3)

	l.DeleteOrder(buyOrder1)

	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()
	sellOrderA := NewOrder(false, 10)
	sellOrderB := NewOrder(false, 100)
	ob.PlaceLimitOrder(65_000, sellOrderA)
	ob.PlaceLimitOrder(67_000, sellOrderB)
	assert(t, len(ob.Asks()), 2)
}

func TestMarketPlaceOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrder := NewOrder(false, 20)
	ob.PlaceLimitOrder(10_000, sellOrder)

	buyOrder := NewOrder(true, 10)
	matches := ob.PlaceMarketOrder(buyOrder)

	assert(t, len(matches), 1)
	assert(t, len(ob.asks), 1)
	assert(t, ob.AskTotalVolume(), 10)
	assert(t, matches[0].Ask, sellOrder)
	assert(t, matches[0].Bid, buyOrder)
	assert(t, matches[0].SizeFilled, 10.0)
	assert(t, matches[0].Price, 10_000.0)
	assert(t, buyOrder.IsFilled(), true)

	fmt.Printf("%v", matches)

}
