import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'
import InvestorResultView from '../investor_result/InvestorResultView'

class AllocationResultView extends React.Component {
  render() {
    const allocationResults = this.props.model.investorResults.map((result) => {
      return <InvestorResultView
        key={result.id}
        model={result}
      />
    })
    return (
      <div className="centeredContainer">
        <h3 className="defaultTitle">Results</h3>
        {allocationResults}
      </div>
    )
  }
}

export default AllocationResultView
