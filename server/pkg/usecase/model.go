package usecase

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/hermant2/angelventureserver/pkg/entity"
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
	requestedAmountAllocations []*entity.InvestorAllocation
}

type InvestorInput struct {
	UUID            string
	Name            string
	RequestedAmount decimal.Decimal
	AverageAmount   decimal.Decimal
}

func (prorateCalculationModel prorateCalculation) doInvestorsReceiveRequestedAmounts() bool {
	return len(prorateCalculationModel.requestedAmountAllocations) > 0
}

func (input ProrateInput) validate() error {
	if input.TotalAllocation.LessThanOrEqual(decimal.Zero) {
		return apperror.Unprocessable(apperror.InputZero)
	} else if len(input.Investors) <= 0 {
		return apperror.Unprocessable(apperror.NoInvestors)
	}

	for _, investor := range input.Investors {
		if err := investor.validate(); err != nil {
			return err
		}
	}

	return nil
}

func (input ProrateInput) doesTotalRequestedAmountExceedAvailableAllocation() bool {
	totalRequestedAmount := decimal.Zero

	for _, investorInput := range input.Investors {
		totalRequestedAmount = totalRequestedAmount.Add(investorInput.appliedRequestedAmount(input.TotalAllocation))
		if totalRequestedAmount.GreaterThan(input.TotalAllocation) {
			return true
		}
	}

	return totalRequestedAmount.GreaterThan(input.TotalAllocation)
}

func (input ProrateInput) calculateAllocationFractionDenominator() decimal.Decimal {
	denominator := decimal.Zero

	for _, investorInput := range input.Investors {
		denominator = denominator.Add(investorInput.appliedInvestorAllocationFractionNumerator(input.TotalAllocation))
	}

	return denominator
}

func (input ProrateInput) generateRequestedAmountOutputs() []*entity.InvestorAllocation {
	var outputs []*entity.InvestorAllocation

	for _, investorInput := range input.Investors {
		output := &entity.InvestorAllocation{
			UUID: investorInput.UUID,
			Name: investorInput.Name,
			AppliedAllocation: investorInput.appliedRequestedAmount(input.TotalAllocation)}
		outputs = append(outputs, output)
	}

	return outputs
}

func (input ProrateInput) generateProrateCalculationModel() (*prorateCalculation, error) {
	if input.TotalAllocation.LessThanOrEqual(decimal.Zero) {
		return nil, apperror.Unprocessable(apperror.InputZero)
	} else if len(input.Investors) <= 0 {
		return nil, apperror.Unprocessable(apperror.NoInvestors)
	}
	totalRequestedAmount := decimal.Zero
	weightedDenominator := decimal.Zero
	var requestedAmountOutputs []*entity.InvestorAllocation

	for _, investorInput := range input.Investors {
		if err := investorInput.validate(); err != nil {
			return nil, err
		}

		totalRequestedAmount = totalRequestedAmount.Add(investorInput.appliedRequestedAmount(input.TotalAllocation))
		weightedDenominator = weightedDenominator.Add(investorInput.appliedInvestorAllocationFractionNumerator(input.TotalAllocation))

		if totalRequestedAmount.LessThanOrEqual(input.TotalAllocation) {
			requestedAmountAllocation := &entity.InvestorAllocation{UUID: investorInput.UUID, Name: investorInput.Name,
				AppliedAllocation: investorInput.appliedRequestedAmount(input.TotalAllocation)}
			requestedAmountOutputs = append(requestedAmountOutputs, requestedAmountAllocation)
		}
	}

	if totalRequestedAmount.LessThanOrEqual(input.TotalAllocation) {
		return &prorateCalculation{requestedAmountAllocations: requestedAmountOutputs}, nil
	} else {
		return &prorateCalculation{
			totalAllocation:        input.TotalAllocation,
			weightedDenominator:    weightedDenominator,
			investorInputs:         input.Investors,
			requestedAmountAllocations: nil}, nil
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

func (input InvestorInput) appliedInvestorAllocationFractionNumerator(totalAllocationAmount decimal.Decimal) decimal.Decimal {
	return decimal.Min(input.AverageAmount, input.appliedRequestedAmount(totalAllocationAmount))
}
