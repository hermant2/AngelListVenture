class InvestorInputPresenter {
    constructor(view, model, parentPresenter) {
        this.view = view
        this.model = model
        console.log(parentPresenter)
        this.parentPresenter = parentPresenter
    }

    onNameChange(event) {
        this.parentPresenter.presentInvestorNameChange(this.model.id, event.target.value)
    }

    onRequestedAmountChanged(event) {
        this.parentPresenter.presentInvestorRequestedAmountChange(this.model.id, event.target.value)
    }

    onAverageAmountChanged(event) {
        this.parentPresenter.presentInvestorAverageAmountChange(this.model.id, event.target.value)
    }

    onRemoveInvestorClicked() {
        this.parentPresenter.presentRemoveInvestor(this.model.id)
    }
}

export default InvestorInputPresenter
