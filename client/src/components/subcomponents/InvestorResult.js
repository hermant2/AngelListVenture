import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'

class InvestorResult extends React.Component {
  render() {
    return (
      <div className="inputContainer">
        <h4>{this.props.name}</h4>
        <h4>{this.props.allocation}</h4>
      </div>
    )
  }
}

export default InvestorResult
