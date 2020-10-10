package prorate

import (
	"github.com/shopspring/decimal"
)

type prorateRequest struct {
	AllocationAmount decimal.Decimal   `json:"allocation_amount"`
	InvestorAmounts  []investorRequest `json:"investor_amounts"`
}

type investorRequest struct {
	UUID            string          `json:"uuid"`
	Name            string          `json:"name"`
	RequestedAmount decimal.Decimal `json:"requested_amount"`
	AverageAmount   decimal.Decimal `json:"average_amount"`
}

type prorateResponseWrapper struct {
	Prorate prorateResponse `json:"prorate"`
}

type prorateResponse struct {
	InvestorAllocations []investorAllocationResponse `json:"investor_allocations"`
}

type investorAllocationResponse struct {
	UUID             string          `json:"uuid"`
	Name             string          `json:"name"`
	AllocationAmount decimal.Decimal `json:"allocation_amount"`
}
