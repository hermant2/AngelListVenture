import InvestorResultModel from "../components/investor_result/InvestorResultModel";
import AllocationResultModel from "../components/allocation_result/AllocationResultModel";
import { formatUSDCurrency } from "../formatters/currency-formatter";

export function mapAllocationResult(investmentAllocationResponses) {
    const investorResults = investmentAllocationResponses.map(response => {
        return new InvestorResultModel(response.id, response.name, formatUSDCurrency(response.allocationAmount))
    })

    return new AllocationResultModel(investorResults)
}
