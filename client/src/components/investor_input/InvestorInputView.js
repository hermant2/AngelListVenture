import React from 'react'
import CurrencyInput from 'react-currency-input'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
import {investorStrings, prorateStrings} from "../../strings/localized-strings";
import InvestorInputPresenter from "./InvestorInputPresenter";

class InvestorInputView extends React.Component {
    constructor(props) {
        super(props)
        this.presenter = new InvestorInputPresenter(this, this.props.model, this.props.parentPresenter)
    }

    render() {
        return (
            <div className="inputContainer">
                <input
                    data-testid="name"
                    className="inputDefault"
                    type="text"
                    placeholder={investorStrings.investorName}
                    value={this.props.model.name}
                    onChange={this.presenter.onNameChange.bind(this.presenter)}
                />

                <CurrencyInput
                    data-testid="requestedAmount"
                    className="inputDefault"
                    precision="2"
                    allowNegative={false}
                    allowEmpty={true}
                    prefix={prorateStrings.currencyPrefix}
                    placeholder={investorStrings.requestedAmount}
                    value={this.props.model.requestedAmount}
                    onChangeEvent={this.presenter.onRequestedAmountChanged.bind(this.presenter)}
                />

                <CurrencyInput
                    data-testid="averageAmount"
                    className="inputDefault"
                    precision="2"
                    allowNegative={false}
                    allowEmpty={true}
                    prefix={prorateStrings.currencyPrefix}
                    placeholder={investorStrings.averageAmount}
                    value={this.props.model.averageAmount}
                    onChangeEvent={this.presenter.onAverageAmountChanged.bind(this.presenter)}
                />

                <button
                    data-testid="remove"
                    className="buttonDefault"
                    onClick={this.presenter.onRemoveInvestorClicked.bind(this.presenter)}>{investorStrings.remove}
                </button>
            </div>
        )
    }
}

export default InvestorInputView
