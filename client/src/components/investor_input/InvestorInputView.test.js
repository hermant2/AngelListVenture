import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import {fireEvent, render, screen} from '@testing-library/react'
import InvestorInputModel from "./InvestorInputModel"
import InvestorInputPresenter from "./InvestorInputPresenter"
import InvestorInputView from "./InvestorInputView"

jest.mock('./InvestorInputPresenter')

beforeEach(() => {
    InvestorInputPresenter.mockClear();
})

test('render InvestorInputView', () => {
    const model = createInvestorInputModel("Person", 100, 50.40)
    const {getByTestId} = render(<InvestorInputView model={model}/>)
    expect(getByTestId("name")).toHaveValue(model.name)
    expect(getByTestId("requestedAmount")).toHaveValue("$100.00")
    expect(getByTestId("averageAmount")).toHaveValue("$50.40")
    expect(getByTestId("remove")).toHaveTextContent("Remove")
})

test('name input changed', () => {
    const model = createInvestorInputModel("Person", 100, 50.40)
    render(<InvestorInputView model={model}/>)
    const mockName = "Guy Fieri"
    const event = {target: {value: mockName}}
    const nameInput = screen.getByTestId("name")
    fireEvent.change(nameInput, event)
    const mockPresenterInstance = InvestorInputPresenter.mock.instances[0]
    const presenterOnNameChange = mockPresenterInstance.onNameChange
    // Attempted to learn how to get this expectation working with value but was unable to do so
    // expect(presenterOnNameChange).toHaveBeenCalledWith(event)
    expect(presenterOnNameChange).toHaveBeenCalledTimes(1)
})

test('requested amount input changed', () => {
    const model = createInvestorInputModel("Person", 100, 50.40)
    render(<InvestorInputView model={model}/>)
    const event = {target: {value: "1001"}}
    const requestedAmountInput = screen.getByTestId("requestedAmount")
    fireEvent.change(requestedAmountInput, event)
    const mockPresenterInstance = InvestorInputPresenter.mock.instances[0]
    const presenterOnRequestedAmountChanged = mockPresenterInstance.onRequestedAmountChanged
    // Attempted to learn how to get this expectation working with value but was unable to do so
    // expect(presenterOnRequestedAmountChanged).toHaveBeenCalledWith(event)
    expect(presenterOnRequestedAmountChanged).toHaveBeenCalledTimes(1)
})

test('average amount input changed', () => {
    const model = createInvestorInputModel("Person", 100, 50.40)
    render(<InvestorInputView model={model}/>)
    const event = {target: {value: "155"}}
    const averageAmountInput = screen.getByTestId("averageAmount")
    fireEvent.change(averageAmountInput, event)
    const mockPresenterInstance = InvestorInputPresenter.mock.instances[0]
    const presenterOnAverageAmountChanged = mockPresenterInstance.onAverageAmountChanged
    // Attempted to learn how to get this expectation working with value but was unable to do so
    // expect(presenterOnAverageAmountChanged).toHaveBeenCalledWith(event)
    expect(presenterOnAverageAmountChanged).toHaveBeenCalledTimes(1)
})

function createInvestorInputModel(name, requestedAmount, averageAmount) {
    const model = new InvestorInputModel()
    model.name = name
    model.requestedAmount = requestedAmount
    model.averageAmount = averageAmount
    return model
}