import React from 'react'
import InvestorInputView from '../investor_input/InvestorInputView'
import CurrencyInput from 'react-currency-input'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
import '../../styles/Header.css'
import '../../styles/Helper.css'
import AllocationInputPresenter from "./AllocationInputPresenter"
import AllocationInputModel from "./AllocationInputModel"
import {prorateStrings} from "../../strings/localized-strings"

class AllocationInputView extends React.Component {
    constructor(props) {
        super(props)
        const model = new AllocationInputModel()
        this.state = {model: model}
        this.presenter = new AllocationInputPresenter(this, model, this.props.parentPresenter)
    }

    updateState(model) {
        this.setState({model: model})
    }

    displayError(message) {
        alert(message)
    }

    render() {
        const investorInputs = this.state.model.investorInputs.map(input => {
            return <InvestorInputView
                key={input.id}
                parentPresenter={this.presenter}
                model={input}
            />
        })

        return (
            <div className="centeredContainer">
                <h3 className="defaultTitle">{prorateStrings.totalAvailableAllocation}</h3>
                <CurrencyInput
                    data-testid="totalAllocation"
                    className="inputDefault"
                    precision="2"
                    allowNegative={false}
                    allowEmpty={true}
                    prefix={prorateStrings.currencyPrefix}
                    placeholder={prorateStrings.allocation}
                    value={this.state.model.availableAllocation}
                    onChangeEvent={this.presenter.onTotalAllocationChanged.bind(this.presenter)}
                />
                {investorInputs}
                <button
                    data-testid="addInvestor"
                    className="buttonDefault spacingTop"
                    onClick={this.presenter.onAddInvestorClicked.bind(this.presenter)}>{prorateStrings.addInvestor}
                </button>
                <button
                    data-testid="prorate"
                    className="buttonDefault spacingTop"
                    onClick={this.presenter.onProrateClicked.bind(this.presenter)}>{prorateStrings.prorate}
                </button>
            </div>
        )
    }
}

export default AllocationInputView
