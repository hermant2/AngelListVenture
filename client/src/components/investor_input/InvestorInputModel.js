import uuid from "react-uuid"

class InvestorInputModel {
    id = uuid()
    name = null
    requestedAmount = null
    averageAmount = null
}

export default InvestorInputModel