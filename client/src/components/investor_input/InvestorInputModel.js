import uuid from "react-uuid"

class InvestorInputModel {
    id = uuid()
    name = ""
    requestedAmount = null
    averageAmount = null
}

export default InvestorInputModel