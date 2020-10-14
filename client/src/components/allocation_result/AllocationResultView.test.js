import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import AllocationResultModel from "./AllocationResultModel"
import InvestorResultModel from "../investor_result/InvestorResultModel"
import {render, screen} from "@testing-library/react"
import AllocationResultView from "./AllocationResultView"

test('render AllocationResultView', () => {
    const model = createAllocationResultModel()
    const {getByTestId} = render(<AllocationResultView model={model}/>)
    expect(getByTestId("resultsHeader")).toHaveTextContent("Results")
})

function createAllocationResultModel() {
    const investor1 = new InvestorResultModel("id1", "Name 1", "$5.00")
    const investor2 = new InvestorResultModel("id2", "Name 2", "$10.00")
    const investor3 = new InvestorResultModel("id3", "Name 3", "$20.00")
    return new AllocationResultModel([investor1, investor2, investor3])
}
