package usecase

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/shopspring/decimal"
)

type ProrateInput struct {
	TotalAllocation decimal.Decimal
	Investors       []InvestorInput
}

type prorateCalculation struct {
	totalAllocation        decimal.Decimal
	weightedDenominator    decimal.Decimal
	investorInputs         []InvestorInput
	requestedAmountOutputs []*InvestorOutput
}

type InvestorInput struct {
	UUID            string
	Name            string
	RequestedAmount decimal.Decimal
	AverageAmount   decimal.Decimal
}

type InvestorOutput struct {
	UUID              string
	Name              string
	AppliedAllocation decimal.Decimal
	requestedAmount   decimal.Decimal
	averageAmount     decimal.Decimal
}

func (prorateCalculationModel prorateCalculation) doInvestorsReceiveRequestedAmounts() bool {
	return len(prorateCalculationModel.requestedAmountOutputs) > 0
}

func (input ProrateInput) generateProrateCalculationModel() (*prorateCalculation, error) {
	if input.TotalAllocation.LessThanOrEqual(decimal.Zero) {
		return nil, apperror.Unprocessable(apperror.InputZero)
	} else if len(input.Investors) <= 0 {
		return nil, apperror.Unprocessable(apperror.NoInvestors)
	}
	totalRequestedAmount := decimal.Zero
	weightedDenominator := decimal.Zero
	var requestedAmountOutputs []*InvestorOutput

	for _, investorInput := range input.Investors {
		if err := investorInput.validate(); err != nil {
			return nil, err
		}

		totalRequestedAmount = totalRequestedAmount.Add(investorInput.appliedRequestedAmount(input.TotalAllocation))
		weightedDenominator = weightedDenominator.Add(investorInput.appliedAllocationNumerator(input.TotalAllocation))

		if totalRequestedAmount.LessThanOrEqual(input.TotalAllocation) {
			requestedAmountOutput := &InvestorOutput{UUID: investorInput.UUID, Name: investorInput.Name,
				AppliedAllocation: investorInput.appliedRequestedAmount(input.TotalAllocation)}
			requestedAmountOutputs = append(requestedAmountOutputs, requestedAmountOutput)
		}
	}

	if totalRequestedAmount.LessThanOrEqual(input.TotalAllocation) {
		return &prorateCalculation{requestedAmountOutputs: requestedAmountOutputs}, nil
	} else {
		return &prorateCalculation{
			totalAllocation:        input.TotalAllocation,
			weightedDenominator:    weightedDenominator,
			investorInputs:         input.Investors,
			requestedAmountOutputs: nil}, nil
	}
}

func (input InvestorInput) validate() error {
	if input.RequestedAmount.LessThanOrEqual(decimal.Zero) || input.AverageAmount.LessThanOrEqual(decimal.Zero) {
		return apperror.Unprocessable(apperror.InputZero)
	} else {
		return nil
	}
}

func (input InvestorInput) appliedRequestedAmount(totalInvestmentAmount decimal.Decimal) decimal.Decimal {
	return decimal.Min(input.RequestedAmount, totalInvestmentAmount)
}

func (input InvestorInput) appliedAllocationNumerator(totalAllocationAmount decimal.Decimal) decimal.Decimal {
	return decimal.Min(input.AverageAmount, input.appliedRequestedAmount(totalAllocationAmount))
}

func (output InvestorOutput) appliedRemainingAllocationNumerator() decimal.Decimal {
	if output.AppliedAllocation.LessThan(output.requestedAmount) {
		return decimal.Min(output.requestedAmount, output.averageAmount)
	} else {
		return decimal.Zero
	}
}
