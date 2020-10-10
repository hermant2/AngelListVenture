package prorate

import (
	"github.com/hermant2/angelventureserver/pkg/usecase"
)

func mapProrateInput(request prorateRequest) usecase.ProrateInput {
	input := usecase.ProrateInput{TotalAllocation: request.AllocationAmount}
	for _, investorRequest := range request.InvestorAmounts {
		investorInput := usecase.InvestorInput{
			UUID:            investorRequest.UUID,
			Name:            investorRequest.Name,
			RequestedAmount: investorRequest.RequestedAmount,
			AverageAmount:   investorRequest.AverageAmount}
		input.Investors = append(input.Investors, investorInput)
	}

	return input
}

func mapProrateResponseWrapper(investorOutputs []*usecase.InvestorOutput) prorateResponseWrapper {
	response := prorateResponse{}
	for _, output := range investorOutputs {
		allocationResponse := investorAllocationResponse{
			UUID:             output.UUID,
			Name:             output.Name,
			AllocationAmount: output.AppliedAllocation.Round(5)}
		response.InvestorAllocations = append(response.InvestorAllocations, allocationResponse)
	}

	return prorateResponseWrapper{response}
}
