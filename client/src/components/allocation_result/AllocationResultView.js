import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
import {prorateStrings} from "../../strings/localized-strings"
import InvestorResultView from '../investor_result/InvestorResultView'

class AllocationResultView extends React.Component {
    render() {
        const allocationResults = this.props.model.investorResults.map((result) => {
            return <InvestorResultView
                data-testid="investorResult"
                key={result.id}
                model={result}
            />
        })
        return (
            <div className="centeredContainer">
                <h3 data-testid="resultsHeader" className="defaultTitle">{prorateStrings.results}</h3>
                {allocationResults}
            </div>
        )
    }
}

export default AllocationResultView
