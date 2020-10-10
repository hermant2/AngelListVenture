import React from 'react'
import '../../styles/Widgets.css'
import '../../styles/Container.css'

class InvestorInput extends React.Component {
  removeInvestor() {
    this.props.removeAction(this.props.id)
  }

  handleChange(field, event) {
    this.props.onChange(this.props.index, field, event.target.value)
  }

  render() {
    return (
      <div className="inputContainer">
        <input
          className="inputDefault"
          type="text"
          name="name"
          placeholder="Investor Name"
          value={this.props.nameValue}
          onChange={this.handleChange.bind(this, "name")}
        />

        <input
          className="inputDefault"
          type="number"
          name="requestedAmount"
          placeholder="Requested Amount"
          value={this.props.requestedAmountValue}
          onChange={this.handleChange.bind(this, "requestedAmount")}
        />

        <input
          className="inputDefault"
          type="number"
          name="averageAmount"
          placeholder="Average Amount"
          value={this.props.averageAmountValue}
          onChange={this.handleChange.bind(this, "averageAmount")}
        />

        <button
          className="buttonDefault"
          onClick={this.removeInvestor.bind(this)}>Remove</button>
      </div>
    )
  }
}

export default InvestorInput
