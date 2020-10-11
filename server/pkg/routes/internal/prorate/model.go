package prorate

import (
	"github.com/shopspring/decimal"
)

type prorateRequest struct {
	AllocationAmount decimal.Decimal   `json:"allocationAmount"`
	InvestorAmounts  []investorRequest `json:"investorAmounts"`
}

type investorRequest struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	RequestedAmount decimal.Decimal `json:"requestedAmount"`
	AverageAmount   decimal.Decimal `json:"averageAmount"`
}

type prorateResponseWrapper struct {
	Prorate prorateResponse `json:"prorate"`
}

type prorateResponse struct {
	InvestorAllocations []investorAllocationResponse `json:"investorAllocations"`
}

type investorAllocationResponse struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	AllocationAmount decimal.Decimal `json:"allocationAmount"`
}
