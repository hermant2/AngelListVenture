package test

import (
	"github.com/hermant2/angelventureserver/pkg/applogger"
	"github.com/hermant2/angelventureserver/pkg/entity"
	"github.com/hermant2/angelventureserver/pkg/usecase"
	"github.com/sirupsen/logrus/hooks/test"
)

type ProrateService struct {
	Input                   *usecase.ProrateInput
	InvestorAllocationsStub []*entity.InvestorAllocation
	ErrorStub               error
}

func (service *ProrateService) CalculateInvestorAllocation(input usecase.ProrateInput) ([]*entity.InvestorAllocation, error) {
	service.Input = &input
	return service.InvestorAllocationsStub, service.ErrorStub
}

func AppLogger() *applogger.AppLogger {
	nullLogger, _ := test.NewNullLogger()
	return &applogger.AppLogger{Client: nullLogger}
}
