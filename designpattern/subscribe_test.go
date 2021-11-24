package main

import (
	"fmt"
	"testing"
)

func TestSubscribe(t *testing.T) {
	sub := ConcreteSubject{}
	obs1 := NewConcreteObserver1(&sub)
	obs2 := NewConcreteObserver2(&sub)

	sub.NotifyObserver()
	fmt.Println(obs1, obs2)
}

// go test -v
