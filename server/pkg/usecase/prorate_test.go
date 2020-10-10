package usecase_test

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/hermant2/angelventureserver/pkg/usecase"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestProrateService_CalculateInvestorAllocation(t *testing.T) {
	service := usecase.NewProrateService()

	t.Run("ComplexInput1_Results", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(100),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(95)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(2), AverageAmount: decimal.NewFromInt(1)},
				{UUID: "id3", Name: "Name3", RequestedAmount: decimal.NewFromInt(1), AverageAmount: decimal.NewFromInt(4)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(err)

		assertions.Equal(len(outputs), 3)

		expectedAllocation1, _ := decimal.NewFromString("97.96875")
		assertions.Equal(outputs[0].UUID, "id1")
		assertions.Equal(outputs[0].Name, "Name1")
		assertions.True(outputs[0].AppliedAllocation.Round(5).Equal(expectedAllocation1))

		expectedAllocation2, _ := decimal.NewFromString("1.03125")
		assertions.Equal(outputs[1].UUID, "id2")
		assertions.Equal(outputs[1].Name, "Name2")
		assertions.True(outputs[1].AppliedAllocation.Round(5).Equal(expectedAllocation2))

		expectedAllocation3, _ := decimal.NewFromString("1")
		assertions.Equal(outputs[2].UUID, "id3")
		assertions.Equal(outputs[2].Name, "Name3")
		assertions.True(outputs[2].AppliedAllocation.Round(5).Equal(expectedAllocation3))
	})

	t.Run("ComplexInput2_Results", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(100),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(150), AverageAmount: decimal.NewFromInt(95)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(1), AverageAmount: decimal.NewFromInt(1)},
				{UUID: "id3", Name: "Name3", RequestedAmount: decimal.NewFromInt(1), AverageAmount: decimal.NewFromInt(4)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(err)

		assertions.Equal(len(outputs), 3)

		expectedAllocation1, _ := decimal.NewFromString("98")
		assertions.Equal(outputs[0].UUID, "id1")
		assertions.Equal(outputs[0].Name, "Name1")
		assertions.True(outputs[0].AppliedAllocation.Round(5).Equal(expectedAllocation1))

		expectedAllocation2, _ := decimal.NewFromString("1")
		assertions.Equal(outputs[1].UUID, "id2")
		assertions.Equal(outputs[1].Name, "Name2")
		assertions.True(outputs[1].AppliedAllocation.Round(5).Equal(expectedAllocation2))

		expectedAllocation3, _ := decimal.NewFromString("1")
		assertions.Equal(outputs[2].UUID, "id3")
		assertions.Equal(outputs[2].Name, "Name3")
		assertions.True(outputs[2].AppliedAllocation.Round(5).Equal(expectedAllocation3))
	})

	t.Run("SimpleInput1_Results", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(100),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(100)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(25), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(err)

		assertions.Equal(len(outputs), 2)

		expectedAllocation1, _ := decimal.NewFromString("80")
		assertions.Equal(outputs[0].UUID, "id1")
		assertions.Equal(outputs[0].Name, "Name1")
		assertions.True(outputs[0].AppliedAllocation.Round(5).Equal(expectedAllocation1))

		expectedAllocation2, _ := decimal.NewFromString("20")
		assertions.Equal(outputs[1].UUID, "id2")
		assertions.Equal(outputs[1].Name, "Name2")
		assertions.True(outputs[1].AppliedAllocation.Round(5).Equal(expectedAllocation2))
	})

	t.Run("SimpleInput2_Results", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(200),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(100)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(25), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(err)

		assertions.Equal(len(outputs), 2)

		expectedAllocation1, _ := decimal.NewFromString("100")
		assertions.Equal(outputs[0].UUID, "id1")
		assertions.Equal(outputs[0].Name, "Name1")
		assertions.True(outputs[0].AppliedAllocation.Round(5).Equal(expectedAllocation1))

		expectedAllocation2, _ := decimal.NewFromString("25")
		assertions.Equal(outputs[1].UUID, "id2")
		assertions.Equal(outputs[1].Name, "Name2")
		assertions.True(outputs[1].AppliedAllocation.Round(5).Equal(expectedAllocation2))
	})

	t.Run("TotalAllocationZero_ReturnError", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(0),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(100)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(25), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(outputs)

		appError := err.(apperror.Standard)
		assertions.Equal(appError.Status, http.StatusUnprocessableEntity)
		assertions.Equal(appError.Code, apperror.InputZero)
	})

	t.Run("TotalAllocationZero_ReturnError", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(0),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(100)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(25), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(outputs)

		appError := err.(apperror.Standard)
		assertions.Equal(appError.Status, http.StatusUnprocessableEntity)
		assertions.Equal(appError.Code, apperror.InputZero)
	})

	t.Run("NoInvestors_ReturnError", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(100),
			Investors: []usecase.InvestorInput{}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(outputs)

		appError := err.(apperror.Standard)
		assertions.Equal(appError.Status, http.StatusUnprocessableEntity)
		assertions.Equal(appError.Code, apperror.NoInvestors)
	})

	t.Run("InvestorRequestedAmountZero_ReturnError", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(100),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(100)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(0), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(outputs)

		appError := err.(apperror.Standard)
		assertions.Equal(appError.Status, http.StatusUnprocessableEntity)
		assertions.Equal(appError.Code, apperror.InputZero)
	})

	t.Run("InvestorAverageAmountZero_ReturnError", func(t *testing.T) {
		input := usecase.ProrateInput{
			TotalAllocation: decimal.NewFromInt(0),
			Investors: []usecase.InvestorInput{
				{UUID: "id1", Name: "Name1", RequestedAmount: decimal.NewFromInt(100), AverageAmount: decimal.NewFromInt(0)},
				{UUID: "id2", Name: "Name2", RequestedAmount: decimal.NewFromInt(25), AverageAmount: decimal.NewFromInt(25)}}}
		outputs, err := service.CalculateInvestorAllocation(input)

		assertions := assert.New(t)
		assertions.Nil(outputs)

		appError := err.(apperror.Standard)
		assertions.Equal(appError.Status, http.StatusUnprocessableEntity)
		assertions.Equal(appError.Code, apperror.InputZero)
	})
}
