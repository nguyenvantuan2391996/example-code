package main

import "fmt"

type ShippingStrategy interface {
    CalculateFee(weight int) int
}

type GHNShipping struct{}

func (GHNShipping) CalculateFee(weight int) int {
    return 15000 + weight*10
}

type GHTKShipping struct{}

func (GHTKShipping) CalculateFee(weight int) int {
    return 12000 + weight*12
}

type ViettelShipping struct{}

func (ViettelShipping) CalculateFee(weight int) int {
    return 14000 + weight*9
}

type ShippingService struct {
    strategy ShippingStrategy
}

func NewShippingService(strategy ShippingStrategy) *ShippingService {
    return &ShippingService{strategy: strategy}
}

func (s *ShippingService) SetStrategy(strategy ShippingStrategy) {
    s.strategy = strategy
}

func (s *ShippingService) Calculate(weight int) int {
    return s.strategy.CalculateFee(weight)
}

package main

import "fmt"

// Strategy interface
type ShippingStrategy interface {
    CalculateFee(weight int) int
}

// Concrete strategy 1
type GHNShipping struct{}

func (GHNShipping) CalculateFee(weight int) int {
    return 15000 + weight*10
}

// Concrete strategy 2
type GHTKShipping struct{}

func (GHTKShipping) CalculateFee(weight int) int {
    return 12000 + weight*12
}

// Concrete strategy 3
type ViettelShipping struct{}

func (ViettelShipping) CalculateFee(weight int) int {
    return 14000 + weight*9
}

// Context
type ShippingService struct {
    strategy ShippingStrategy
}

func NewShippingService(strategy ShippingStrategy) *ShippingService {
    return &ShippingService{strategy: strategy}
}

func (s *ShippingService) SetStrategy(strategy ShippingStrategy) {
    s.strategy = strategy
}

func (s *ShippingService) Calculate(weight int) int {
    return s.strategy.CalculateFee(weight)
}
