import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'

class InvestorResultView extends React.Component {
  render() {
    return (
      <div className="inputContainer">
        <h4 data-testid="name">{this.props.model.name}</h4>
        <h4 data-testid="amount">{this.props.model.amount}</h4>
      </div>
    )
  }
}

export default InvestorResultView
