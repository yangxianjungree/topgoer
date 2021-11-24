package main

import (
	"fmt"
	"time"
)

/*
Data struct.
*/
type Data struct {
}

/*
Observer interface.
*/
type Observer interface {
	// update
	Update(*Data)
}

/*
Subject interface.
*/
type Subject interface {
	// register
	RegisterObserver(obs Observer)
	// remove
	RemoveObserver(obs Observer)
	// notify
	NotifyObserver()
}

/*
ConcreteSubject concrete struct.
*/
type ConcreteSubject struct {
	observers []Observer
	data      *Data
}

func (s *ConcreteSubject) RegisterObserver(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *ConcreteSubject) RemoveObserver(obs Observer) {
	idx := -1
	for i := range s.observers {
		if s.observers[i] == obs {
			idx = i
			break
		}
	}

	if idx != -1 {
		s.observers = append(s.observers[:idx], s.observers[idx+1:]...)
	}
}

func (s *ConcreteSubject) NotifyObserver() {
	for _, obs := range s.observers {
		obs.Update(s.data)
	}
}

// ConcreObserver1:
type ConcreObserver1 struct {
}

func (c *ConcreObserver1) Update(data *Data) {
	fmt.Println("Observer 1 got data, time ", time.Now())
}

func NewConcreteObserver1(sub Subject) *ConcreObserver1 {
	c := &ConcreObserver1{}
	sub.RegisterObserver(c)
	return c
}

// ConcreObserver2:
type ConcreObserver2 struct {
}

func (c *ConcreObserver2) Update(data *Data) {
	fmt.Println("Observer 2 got data, now: ", time.Now())
}

func NewConcreteObserver2(sub Subject) *ConcreObserver2 {
	c := &ConcreObserver2{}
	sub.RegisterObserver(c)
	return c
}
