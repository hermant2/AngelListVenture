import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
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
                    className="inputDefault"
                    type="text"
                    name="name"
                    placeholder="Investor Name"
                    value={this.props.model.name}
                    onChange={this.presenter.onNameChange.bind(this.presenter)}
                />

                <input
                    className="inputDefault"
                    type="number"
                    name="requestedAmount"
                    placeholder="Requested Amount"
                    value={this.props.model.requestedAmount}
                    onChange={this.presenter.onRequestedAmountChanged.bind(this.presenter)}
                />

                <input
                    className="inputDefault"
                    type="number"
                    name="averageAmount"
                    placeholder="Average Amount"
                    value={this.props.model.averageAmount}
                    onChange={this.presenter.onAverageAmountChanged.bind(this.presenter)}
                />

                <button
                    className="buttonDefault"
                    onClick={this.presenter.onRemoveInvestorClicked}>Remove
                </button>
            </div>
        )
    }
}

export default InvestorInputView
