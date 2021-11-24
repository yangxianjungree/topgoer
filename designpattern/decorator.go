package main

type Beverage interface {
	Cost() float64
	Desc() string
}

type Espresso struct {
}

func (e *Espresso) Cost() float64 {
	return 1.9
}

func (e *Espresso) Desc() string {
	return "Espresso"
}

type HouseBlend struct {
}

func (h *HouseBlend) Cost() float64 {
	return 2.0
}

func (h *HouseBlend) Desc() string {
	return "HouseBlend"
}

// Add decorator---------------------------

type Mocha struct {
	beverage Beverage
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.1
}

func (m *Mocha) Desc() string {
	return m.beverage.Desc() + " Mocha"
}

func NewMocha(b Beverage) Beverage {
	return &Mocha{b}
}

type Whip struct {
	beverage Beverage
}

func (w *Whip) Cost() float64 {
	return w.beverage.Cost() + 0.2
}

func (w *Whip) Desc() string {
	return w.beverage.Desc() + " Whip"
}

func NewWhip(b Beverage) Beverage {
	return &Whip{b}
}

type Soy struct {
	beverage Beverage
}

func (s *Soy) Cost() float64 {
	return s.beverage.Cost() + 0.3
}

func (s *Soy) Desc() string {
	return s.beverage.Desc() + " Soy"
}

func NewSoy(b Beverage) Beverage {
	return &Soy{b}
}
