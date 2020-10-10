import React from 'react'
import '../styles/Widgets.css'
import '../styles/Container.css'
import InvestorResult from './subcomponents/InvestorResult'

class AllocationResult extends React.Component {
  render() {
    const allocationResults = this.props.results.map((result) => {
      return <InvestorResult
        key={result.id}
        name={result.name}
        allocation={result.allocation}
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

export default AllocationResult
