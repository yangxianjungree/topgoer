package main

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	e := &Espresso{}
	m := NewMocha(e)
	s := NewSoy(m)
	fmt.Println(s.Desc(), "=", s.Cost())

	h := &HouseBlend{}
	w := NewWhip(h)
	fmt.Println(w.Desc(), "=", w.Cost())
}
