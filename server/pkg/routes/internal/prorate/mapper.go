package prorate

import (
	"github.com/hermant2/angelventureserver/pkg/constants"
	"github.com/hermant2/angelventureserver/pkg/entity"
	"github.com/hermant2/angelventureserver/pkg/usecase"
)

func mapProrateInput(request prorateRequest) usecase.ProrateInput {
	input := usecase.ProrateInput{TotalAllocation: request.AllocationAmount}
	for _, investorRequest := range request.InvestorAmounts {
		investorInput := usecase.InvestorInput{
			UUID:            investorRequest.ID,
			Name:            investorRequest.Name,
			RequestedAmount: investorRequest.RequestedAmount,
			AverageAmount:   investorRequest.AverageAmount}
		input.Investors = append(input.Investors, investorInput)
	}

	return input
}

func mapProrateResponseWrapper(investorAllocations []*entity.InvestorAllocation) prorateResponseWrapper {
	response := prorateResponse{}
	for _, allocation := range investorAllocations {
		allocationResponse := investorAllocationResponse{
			ID:               allocation.UUID,
			Name:             allocation.Name,
			AllocationAmount: allocation.AppliedAllocation.Round(constants.DecimalPrecision)}
		response.InvestorAllocations = append(response.InvestorAllocations, allocationResponse)
	}

	return prorateResponseWrapper{response}
}
