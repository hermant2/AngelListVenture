package prorate_test

import (
	"encoding/json"
	"fmt"
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/entity"
	"github.com/hermant2/angelventureserver/pkg/routes/internal/prorate"
	"github.com/hermant2/angelventureserver/pkg/test"
	"github.com/hermant2/angelventureserver/pkg/usecase"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var mockProrateService *test.ProrateService
var mockAppLogger *applogger.AppLogger

func createController() prorate.Controller {
	mockProrateService = &test.ProrateService{}
	mockAppLogger = test.AppLogger()
	return prorate.NewController(mockProrateService, mockAppLogger)
}

func TestController_CalculateAllocation(t *testing.T) {
	mockProrateRequest := map[string]interface{}{
		"allocationAmount": 125.0,
		"investorAmounts": []map[string]interface{}{
			mockInvestorRequest("id1", "Name1", 29.55, 12.10),
			mockInvestorRequest("id2", "Name2", 125, 100.50)},
	}
	mockRequest, _ := test.PostRequest(mockProrateRequest)

	t.Run("ServiceCalculateInvestorAllocationSuccess_RenderResponse", func(t *testing.T) {
		controller := createController()
		mockProrateService.ErrorStub = nil
		mockProrateService.InvestorAllocationsStub = []*entity.InvestorAllocation{
			{UUID: "id1", Name: "Name1", AppliedAllocation: decimal.NewFromFloat(25)},
			{UUID: "id2", Name: "Name2", AppliedAllocation: decimal.NewFromFloat(100)}}

		writer := httptest.NewRecorder()
		controller.CalculateAllocation(writer, mockRequest)

		assertions := assert.New(t)
		assertProrateInput(assertions, mockProrateRequest, *mockProrateService.Input)
		assertProrateResponse(assertions, writer, mockProrateService.InvestorAllocationsStub)
		assertions.Equal(writer.Result().StatusCode, http.StatusOK)
	})

	t.Run("ServiceCalculateInvestorAllocationError_RenderError", func(t *testing.T) {
		controller := createController()
		appErr := apperror.Unprocessable(apperror.General)
		mockProrateService.ErrorStub = appErr
		mockProrateService.InvestorAllocationsStub = nil

		writer := httptest.NewRecorder()
		controller.CalculateAllocation(writer, mockRequest)

		assertions := assert.New(t)

		assertProrateInput(assertions, mockProrateRequest, *mockProrateService.Input)
		test.AssertErrorResponse(assertions, writer, appErr)
		assertions.Equal(writer.Result().StatusCode, http.StatusUnprocessableEntity)
	})

	t.Run("InvalidJSON_RenderBadRequestError", func(t *testing.T) {
		controller := createController()

		invalidRequest, _ := test.InvalidPostRequest()
		writer := httptest.NewRecorder()

		controller.CalculateAllocation(writer, invalidRequest)

		assertions := assert.New(t)
		assertions.Nil(mockProrateService.Input)
		test.AssertErrorResponse(assertions, writer, apperror.BadRequest(apperror.General))
	})
}

func assertProrateInput(assertions *assert.Assertions, request map[string]interface{}, input usecase.ProrateInput) {
	assertions.Equal(decimal.NewFromFloat(request["allocationAmount"].(float64)), input.TotalAllocation)
	fmt.Println(reflect.TypeOf(request["investorAmounts"]))
	investorAmounts := request["investorAmounts"].([]map[string]interface {})
	assertions.Equal(len(investorAmounts), len(input.Investors))

	for i := range input.Investors {
		investorRequest := investorAmounts[i]
		investorInput := input.Investors[i]
		assertions.Equal(investorRequest["id"], investorInput.UUID)
		assertions.Equal(investorRequest["name"], investorInput.Name)
		assertions.Equal(decimal.NewFromFloat(investorRequest["requestedAmount"].(float64)), investorInput.RequestedAmount)
		assertions.Equal(decimal.NewFromFloat(investorRequest["averageAmount"].(float64)), investorInput.AverageAmount)
	}
}

func assertProrateResponse(assertions *assert.Assertions, writer *httptest.ResponseRecorder, allocationEntities []*entity.InvestorAllocation) {
	var jsonResponseMap map[string]interface{}
	json.NewDecoder(writer.Result().Body).Decode(&jsonResponseMap)

	prorateResponse := jsonResponseMap["prorate"].(map[string]interface{})
	investorAllocationResponses := prorateResponse["investorAllocations"].([]interface{})
	assertions.Equal(len(investorAllocationResponses), len(allocationEntities))
	for i := range investorAllocationResponses {
		allocationResponse := investorAllocationResponses[i].(map[string]interface{})
		allocationEntity := allocationEntities[i]
		assertions.Equal(allocationResponse["id"], allocationEntity.UUID)
		assertions.Equal(allocationResponse["name"], allocationEntity.Name)
		assertions.Equal(allocationResponse["allocationAmount"], allocationEntity.AppliedAllocation.String())
	}
}

func mockInvestorRequest(id string, name string, requestedAmount float64, averageAmount float64) map[string]interface{} {
	return map[string]interface{}{
		"id":              id,
		"name":            name,
		"requestedAmount": requestedAmount,
		"averageAmount":   averageAmount}
}
