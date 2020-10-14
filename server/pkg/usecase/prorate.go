package usecase

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/hermant2/angelventureserver/pkg/constants"
	"github.com/hermant2/angelventureserver/pkg/entity"
	"github.com/shopspring/decimal"
)

type ProrateService interface {
	CalculateInvestorAllocation(input ProrateInput) ([]*entity.InvestorAllocation, error)
}

type prorateService struct {
}

func NewProrateService() ProrateService {
	return prorateService{}
}

func (service prorateService) CalculateInvestorAllocation(input ProrateInput) ([]*entity.InvestorAllocation, error) {
	if err := input.validate(); err != nil {
		return nil, err
	} else if input.doesTotalRequestedAmountExceedAvailableAllocation() {
		return calculateProratedOutputs(input)
	} else {
		return input.generateRequestedAmountOutputs(), nil
	}
}

/*
* NOTE: I commented out the following lines of code (as well as in the model.go file in this package), as it was another
* algorithmic structure I was torn between. There are many things we need to know up front before performing the complex
* version of the algorithm. First, we need to determine if the input is valid. Due to this being undefined in the specs,
* I went ahead and assumed that input values could not be zero, and that each investment required at least one investor.
* Second, we need to know if the total requested amount exceeds the total allocation that is available. If it does not,
* Then we can go ahead and return everybody their requested amounts. Otherwise complex version of the algorithm needs to
* be performed. All of these upfront calculations could be performed with a single iteration through the InvestorInputs
* array, and I used a prorateCalculationModel to store those required values. The version of the algorithm that I left
* uncommented simply has separate methods to perform each of these up-front checks separately. Though this iterates over
* the InvestorInputs array more times than necessary, I found this to be less in violation of the single responsibility
* principle, making the code easier to read & reason about. I also assumed that the number of investors would be relatively
* small, so that a couple of more iterations wouldn't impact performance significantly. However, if I were wrong about
* these assumptions, I would return to the algorithmic structure that limited the number of iterations, and would perhaps
* take more time to review & refactor the code so that it is easier for other team members to understand and reason about.
* Would be very curious to discuss and get feedback about this.
 */

//func (service prorateService) CalculateInvestorAllocation(input ProrateInput) ([]*entity.InvestorAllocation, error) {
//	calculationModel, err := input.generateProrateCalculationModel()
//	if err != nil {
//		return nil, err
//	} else if calculationModel.doInvestorsReceiveRequestedAmounts() {
//		return calculationModel.requestedAmountAllocations, nil
//	} else {
//		return calculateProratedOutputs(*calculationModel)
//	}
//}
//
//func calculateProratedOutputs(calculationModel prorateCalculation) ([]*entity.InvestorAllocation, error) {
//	appliedAllocationTotal := decimal.Zero
//	var investorAllocations []*entity.InvestorAllocation
//	for _, investorInput := range calculationModel.investorInputs {
//		appliedNumerator := investorInput.appliedInvestorAllocationFractionNumerator(calculationModel.totalAllocation)
//		investorAllocationAmount := decimal.Min(investorInput.RequestedAmount,
//			calculationModel.totalAllocation.Mul(appliedNumerator.Div(calculationModel.weightedDenominator)))
//		appliedAllocationTotal = appliedAllocationTotal.Add(investorAllocationAmount)
//		allocation := &entity.InvestorAllocation{
//			UUID:              investorInput.UUID,
//			Name:              investorInput.Name,
//			AppliedAllocation: investorAllocationAmount,
//			RequestedAmount:   investorInput.RequestedAmount,
//			AverageAmount:     investorInput.AverageAmount}
//		investorAllocations = append(investorAllocations, allocation)
//	}
//
//	if appliedAllocationTotal.Round(constants.DecimalPrecision).
//		GreaterThan(calculationModel.totalAllocation.Round(constants.DecimalPrecision)) {
//		return nil, apperror.InternalServerError(apperror.General)
//	} else if appliedAllocationTotal.Equal(calculationModel.totalAllocation) {
//		return investorAllocations, nil
//	} else {
//		remaining := calculationModel.totalAllocation.Sub(appliedAllocationTotal)
//		allocateRemainingInvestment(remaining, investorAllocations)
//		return investorAllocations, nil
//	}
//}

// region private
func calculateProratedOutputs(input ProrateInput) ([]*entity.InvestorAllocation, error) {
	appliedAllocationTotal := decimal.Zero
	var investorOutputs []*entity.InvestorAllocation

	weightedAllocationDenominator := input.calculateAllocationFractionDenominator()
	for _, investorInput := range input.Investors {
		appliedNumerator := investorInput.appliedInvestorAllocationFractionNumerator(input.TotalAllocation)
		investorAllocation := decimal.Min(investorInput.RequestedAmount,
			input.TotalAllocation.Mul(appliedNumerator.Div(weightedAllocationDenominator)))
		appliedAllocationTotal = appliedAllocationTotal.Add(investorAllocation)
		output := &entity.InvestorAllocation{
			UUID:              investorInput.UUID,
			Name:              investorInput.Name,
			AppliedAllocation: investorAllocation,
			RequestedAmount:   investorInput.RequestedAmount,
			AverageAmount:     investorInput.AverageAmount}
		investorOutputs = append(investorOutputs, output)
	}

	if appliedAllocationTotal.Round(constants.DecimalPrecision).
		GreaterThan(input.TotalAllocation.Round(constants.DecimalPrecision)) {
		return nil, apperror.InternalServerError(apperror.General)
	} else if appliedAllocationTotal.Equal(input.TotalAllocation) {
		return investorOutputs, nil
	} else {
		remaining := input.TotalAllocation.Sub(appliedAllocationTotal)
		allocateRemainingInvestment(remaining, investorOutputs)
		return investorOutputs, nil
	}
}

func allocateRemainingInvestment(remaining decimal.Decimal, allocations []*entity.InvestorAllocation) {
	weightedDenominator := calculateRemainingAllocationDenominator(allocations)
	for _, allocation := range allocations {
		remainingAllocationNumerator := appliedRemainingInvestorAllocationNumerator(allocation)
		additionalAllocation := remaining.Mul(remainingAllocationNumerator.Div(weightedDenominator))
		allocation.AppliedAllocation = allocation.AppliedAllocation.Add(additionalAllocation)
	}
}

func calculateRemainingAllocationDenominator(allocations []*entity.InvestorAllocation) decimal.Decimal {
	weightedDenominator := decimal.Zero
	for _, allocation := range allocations {
		remainingAllocationNumerator := appliedRemainingInvestorAllocationNumerator(allocation)
		weightedDenominator = weightedDenominator.Add(remainingAllocationNumerator)
	}
	return weightedDenominator
}

func appliedRemainingInvestorAllocationNumerator(allocation *entity.InvestorAllocation) decimal.Decimal {
	if allocation.AppliedAllocation.LessThan(allocation.RequestedAmount) {
		return decimal.Min(allocation.RequestedAmount, allocation.AverageAmount)
	} else {
		return decimal.Zero
	}
}

// endregion private
