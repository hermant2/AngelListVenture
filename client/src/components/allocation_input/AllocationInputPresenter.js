import {calculateInvestmentAllocations} from "../../services/ProrateService";
import InvestorInputModel from "../investor_input/InvestorInputModel";
import AllocationInputModel from "./AllocationInputModel";

class AllocationInputPresenter {
    constructor(view, model, parentPresenter) {
        this.view = view
        this.parentPresenter = parentPresenter
        this.model = new AllocationInputModel()
        this.view.updateState.bind(this, this.model)
    }

    onTotalAllocationChanged(event) {
        this.model.availableAllocation = event.target.value
        this.view.updateState(this.model)
    }

    onProrateClicked() {
        calculateInvestmentAllocations(this.model.generateAPIRequest())
            .then(prorateResponse => {
                this.parentPresenter.presentAllocationResult(prorateResponse.prorate.investorAllocations)
            })
            .catch(error => console.error(error))
    }

    onAddInvestorClicked() {
        this.model.investorInputs.push(new InvestorInputModel())
        this.view.updateState(this.model)
    }

    // region parentPresenter functions
    presentRemoveInvestor(id) {
        this.model.investorInputs = this.model.investorInputs.filter(input => input.id !== id)
        this.view.updateState(this.model)
    }

    presentInvestorNameChange(id, newName) {
        let index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].name = newName
            this.view.updateState(this.model)
        }
    }

    presentInvestorRequestedAmountChange(id, newValue) {
        let index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].requestedAmount = newValue
            this.view.updateState(this.model)
        }
    }

    presentInvestorAverageAmountChange(id, newValue) {
        let index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].averageAmount = newValue
            this.view.updateState(this.model)
        }
    }

    // endregion parentPresenter functions
}

export default AllocationInputPresenter
