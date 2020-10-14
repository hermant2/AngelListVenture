import React from 'react'
import '@testing-library/jest-dom/extend-expect'
import {render} from '@testing-library/react'
import InvestorResultView from "./InvestorResultView"
import InvestorResultModel from "./InvestorResultModel"

test('render InvestorResultView', () => {
    const model = new InvestorResultModel("id", "Name", "$20.00")
    const {getByTestId} = render(<InvestorResultView model={model}/>)
    expect(getByTestId("name")).toHaveTextContent(model.name)
    expect(getByTestId("amount")).toHaveTextContent(model.amount)
})
