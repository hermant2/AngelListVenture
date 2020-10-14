import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import {fireEvent, render, screen} from '@testing-library/react'
import AllocationInputView from "./AllocationInputView"
import AllocationInputPresenter from "./AllocationInputPresenter"

jest.mock('./AllocationInputPresenter')

test('total allocation input changed', () => {
    render(<AllocationInputView />)
    const event = {target: {value: "$500"}}
    const totalAllocationInput = screen.getByTestId("totalAllocation")
    fireEvent.change(totalAllocationInput, event)
    const mockPresenterInstance = AllocationInputPresenter.mock.instances[0]
    const presenterOnTotalAllocationChanged = mockPresenterInstance.onTotalAllocationChanged
    expect(presenterOnTotalAllocationChanged).toHaveBeenCalledTimes(1)
})

test('add investor clicked', () => {
    render(<AllocationInputView />)
    const addInvestorButton = screen.getByTestId("addInvestor")
    fireEvent.click(addInvestorButton)
    const mockPresenterInstance = AllocationInputPresenter.mock.instances[0]
    const presenterOnAddInvestorClicked = mockPresenterInstance.onAddInvestorClicked
    expect(presenterOnAddInvestorClicked).toHaveBeenCalledTimes(1)
})

test('prorate clicked', () => {
    render(<AllocationInputView />)
    const prorateButton = screen.getByTestId("prorate")
    fireEvent.click(prorateButton)
    const mockPresenterInstance = AllocationInputPresenter.mock.instances[0]
    const presenterOnProrateClicked = mockPresenterInstance.onProrateClicked
    expect(presenterOnProrateClicked).toHaveBeenCalledTimes(1)
})
