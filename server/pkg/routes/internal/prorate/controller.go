package prorate

import (
	"github.com/hermant2/angelventureserver/pkg/apperror"
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/routes/internal/api"
	"github.com/hermant2/angelventureserver/pkg/usecase"
	"net/http"
)

type Controller interface {
	CalculateAllocation(writer http.ResponseWriter, request *http.Request)
}

type controller struct {
	service usecase.ProrateService
	logger  *applogger.AppLogger
}

func NewController(service usecase.ProrateService, logger *applogger.AppLogger) Controller {
	return controller{service, logger}
}

func (controller controller) CalculateAllocation(writer http.ResponseWriter, request *http.Request) {
	var prorateRequest prorateRequest
	err := api.DecodeRequest(request.Body, &prorateRequest)
	if err != nil {
		controller.logger.Error("decode_request", err)
		api.RenderError(writer, request, apperror.BadRequest(apperror.General))
		return
	}

	input := mapProrateInput(prorateRequest)
	outputs, err := controller.service.CalculateInvestorAllocation(input)
	if err != nil {
		controller.logger.Error("calculate_allocation", err)
		api.RenderError(writer, request, err)
	} else {
		responseWrapper := mapProrateResponseWrapper(outputs)
		apiResponse := api.Response{Model: responseWrapper, Status: http.StatusOK}
		api.RenderResponse(writer, request, apiResponse)
	}
}
