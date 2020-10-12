import {calculateInvestmentAllocations} from "../../services/prorate-service";
import {mapErrorResponseMessage} from "../../mappers/error-model-mapper";
import InvestorInputModel from "../investor_input/InvestorInputModel";

class AllocationInputPresenter {
    constructor(view, model, parentPresenter) {
        this.view = view
        this.parentPresenter = parentPresenter
        this.model = model
        this.view.updateState.bind(this, this.model)
    }

    onTotalAllocationChanged(event, maskedValue, numericValue) {
        this.model.availableAllocation = numericValue
        this.view.updateState(this.model)
    }

    onProrateClicked() {
        const validationError = this.model.checkValidationError()
        if (validationError) {
            this.parentPresenter.presentAllocationResult([])
            this.view.displayError(validationError)
            return
        }

        calculateInvestmentAllocations(this.model.generateAPIRequest())
            .then(prorateResponse => {
                const investorAllocations = prorateResponse?.prorate?.investorAllocations
                this.parentPresenter.presentAllocationResult(investorAllocations ? investorAllocations : [])
            })
            .catch(error => {
                this.parentPresenter.presentAllocationResult([])
                const errorMessage = mapErrorResponseMessage(error)
                this.view.displayError(errorMessage)
            })
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
        const index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].name = newName
            this.view.updateState(this.model)
        }
    }

    presentInvestorRequestedAmountChange(id, numericValue) {
        const index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].requestedAmount = numericValue
            this.view.updateState(this.model)
        }
    }

    presentInvestorAverageAmountChange(id, numericValue) {
        const index = this.model.investorInputs.findIndex(input => input.id === id)
        if (index >= 0) {
            this.model.investorInputs[index].averageAmount = numericValue
            this.view.updateState(this.model)
        }
    }

    // endregion parentPresenter functions
}

export default AllocationInputPresenter
