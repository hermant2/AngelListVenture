import uuid from "react-uuid"
import InvestorInputModel from "../investor_input/InvestorInputModel";

class AllocationInputModel {
    availableAllocation = 0
    investorInputs = [new InvestorInputModel()]

    generateAPIRequest() {
        return {
            allocationAmount: this.availableAllocation,
            investorAmounts: this.investorInputs
        }
    }
}

export default AllocationInputModel
