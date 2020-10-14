package entity

import "github.com/shopspring/decimal"

type InvestorAllocation struct {
	UUID              string
	Name              string
	AppliedAllocation decimal.Decimal
	RequestedAmount   decimal.Decimal
	AverageAmount     decimal.Decimal
}
