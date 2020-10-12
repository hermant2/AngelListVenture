import InvestorInputModel from "../investor_input/InvestorInputModel";
import {errorStrings} from "../../strings/localized-strings";

class AllocationInputModel {
    availableAllocation = null
    investorInputs = [new InvestorInputModel()]

    generateAPIRequest() {
        return {
            allocationAmount: this.availableAllocation,
            investorAmounts: this.investorInputs
        }
    }

    checkValidationError() {
        const totalAllocation = this.availableAllocation ? this.availableAllocation : 0
        if (totalAllocation <= 0) {
            return errorStrings.inputZero
        }

        if (this.investorInputs.length <= 0) {
            return errorStrings.noInvestors
        }

        for (let i = 0; i < this.investorInputs.length; i++) {
            const input = this.investorInputs[i]
            const requestedAmount = input.requestedAmount ? input.requestedAmount : 0
            const averageAmount = input.averageAmount ? input.averageAmount : 0
            if (requestedAmount <= 0 || averageAmount <= 0) {
                return errorStrings.inputZero
            }
        }

        return null
    }
}

export default AllocationInputModel
