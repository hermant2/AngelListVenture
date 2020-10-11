import React from 'react'
import InvestorInputView from '../investor_input/InvestorInputView'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
import '../../styles/Header.css'
import '../../styles/Helper.css'
import AllocationInputPresenter from "./AllocationInputPresenter";
import AllocationInputModel from "./AllocationInputModel";

class AllocationInputView extends React.Component {
    constructor(props) {
        super(props)
        let model = new AllocationInputModel()
        this.state = model
        this.presenter = new AllocationInputPresenter(this, model, this.props.parentPresenter)
    }

    updateState(model) {
        this.setState(model)
    }

    render() {
        const investorInputs = this.state.investorInputs.map((input, index) => {
            return <InvestorInputView
                key={input.id}
                parentPresenter={this.presenter}
                model={input}
            />
        })

        return (
            <div className="centeredContainer">
                <h3 className="defaultTitle">Total Available Allocation</h3>
                <input
                    className="inputDefault"
                    type="number"
                    name="name"
                    placeholder="Allocation"
                    onChange={this.presenter.onTotalAllocationChanged.bind(this.presenter)}
                />
                {investorInputs}
                <button
                    className="buttonDefault spacingTop"
                    onClick={this.presenter.onAddInvestorClicked.bind(this.presenter)}>Add Investor
                </button>
                <button
                    className="buttonDefault spacingTop"
                    onClick={this.presenter.onProrateClicked.bind(this.presenter)}>Prorate
                </button>
            </div>
        )
    }
}

export default AllocationInputView
