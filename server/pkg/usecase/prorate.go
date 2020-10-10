package usecase

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/shopspring/decimal"
)

type ProrateService interface {
	CalculateInvestorAllocation(input ProrateInput) ([]*InvestorOutput, error)
}

type prorateService struct {
}

func NewProrateService() ProrateService {
	return prorateService{}
}

func (service prorateService) CalculateInvestorAllocation(input ProrateInput) ([]*InvestorOutput, error) {
	calculationModel, err := input.generateProrateCalculationModel()
	if err != nil {
		return nil, err
	} else if calculationModel.doInvestorsReceiveRequestedAmounts() {
		return calculationModel.requestedAmountOutputs, nil
	} else {
		return calculateProratedOutputs(*calculationModel)
	}
}

func calculateProratedOutputs(calculationModel prorateCalculation) ([]*InvestorOutput, error) {
	appliedAllocationTotal := decimal.Zero
	var investorOutputs []*InvestorOutput
	for _, investorInput := range calculationModel.investorInputs {
		appliedNumerator := investorInput.appliedAllocationNumerator(calculationModel.totalAllocation)
		investorAllocation := decimal.Min(investorInput.RequestedAmount,
			calculationModel.totalAllocation.Mul(appliedNumerator.Div(calculationModel.weightedDenominator)))
		appliedAllocationTotal = appliedAllocationTotal.Add(investorAllocation)
		output := &InvestorOutput{
			UUID:              investorInput.UUID,
			Name:              investorInput.Name,
			AppliedAllocation: investorAllocation,
			requestedAmount:   investorInput.RequestedAmount,
			averageAmount:     investorInput.AverageAmount}
		investorOutputs = append(investorOutputs, output)
	}

	if appliedAllocationTotal.Round(5).GreaterThan(calculationModel.totalAllocation.Round(5)) {
		return nil, apperror.InternalServerError(apperror.General)
	} else if appliedAllocationTotal.Equal(calculationModel.totalAllocation) {
		return investorOutputs, nil
	} else {
		remaining := calculationModel.totalAllocation.Sub(appliedAllocationTotal)
		allocateRemainingInvestment(remaining, investorOutputs)
		return investorOutputs, nil
	}
}

func allocateRemainingInvestment(remaining decimal.Decimal, outputs []*InvestorOutput) {
	weightedDenominator := calculateRemainingWeightedDenominator(outputs)
	for _, output := range outputs {
		additionalAllocation := remaining.Mul(output.appliedRemainingAllocationNumerator().Div(weightedDenominator))
		output.AppliedAllocation = output.AppliedAllocation.Add(additionalAllocation)
	}
}

func calculateRemainingWeightedDenominator(outputs []*InvestorOutput) decimal.Decimal {
	weightedDenominator := decimal.Zero
	for _, output := range outputs {
		weightedDenominator = weightedDenominator.Add(output.appliedRemainingAllocationNumerator())
	}
	return weightedDenominator
}
