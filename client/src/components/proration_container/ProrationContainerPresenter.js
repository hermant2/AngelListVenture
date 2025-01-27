import { mapAllocationResult } from "../../mappers/allocation-model-mapper";

class ProrationContainerPresenter {
    constructor(view, model) {
        this.view = view
        this.model = model
    }

    presentAllocationResult(results) {
        this.model.allocationResult = mapAllocationResult(results)
        this.view.updateState(this.model)
    }
}

export default ProrationContainerPresenter
