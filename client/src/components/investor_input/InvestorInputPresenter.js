class InvestorInputPresenter {
    constructor(view, model, parentPresenter) {
        this.view = view
        this.model = model
        this.parentPresenter = parentPresenter
    }

    onNameChange(event) {
        this.parentPresenter.presentInvestorNameChange(this.model.id, event.target.value)
    }

    onRequestedAmountChanged(event, maskedValue, numericValue) {
        this.parentPresenter.presentInvestorRequestedAmountChange(this.model.id, numericValue)
    }

    onAverageAmountChanged(event, maskedValue, numericValue) {
        this.parentPresenter.presentInvestorAverageAmountChange(this.model.id, numericValue)
    }

    onRemoveInvestorClicked() {
        this.parentPresenter.presentRemoveInvestor(this.model.id)
    }
}

export default InvestorInputPresenter
