import React from 'react'
import uuid from 'react-uuid'
import axios from 'axios'
import InvestorInput from './subcomponents/InvestorInput'
import '../styles/Widgets.css'
import '../styles/Container.css'
import '../styles/Header.css'
import '../styles/Helper.css'

class AllocationInput extends React.Component {
  constructor(props) {
    super(props)

    this.removeInvestor = this.removeInvestor.bind(this)
    this.handleInvestorChange = this.handleInvestorChange.bind(this)

    this.state = {
      availableAllocation: 0,
      investorInputs: [{id: uuid()}]
    }
  }

  addInvestor() {
    let newInvestorInputs = this.state.investorInputs.concat({id: uuid()})
    this.setState({"investorInputs": newInvestorInputs})
  }

  prorate() {
    let request = {
      allocationAmount: this.state.availableAllocation,
      investorAmounts: this.state.investorInputs
    }
    axios.post("http://localhost:8080/api/v1/prorate", request)
        .then(response => this.props.onAllocationResult(response))
        .catch(error => console.error(error))
  }

  removeInvestor(id) {
    let newInputs = this.state.investorInputs.filter(input => input.id !== id)
    this.setState({"investorInputs": newInputs})
  }

  handleInvestorChange(index, field, value) {
    let updatedInput = this.state.investorInputs[index]
    updatedInput[field] = value
    let newInputs = [
        ...this.state.investorInputs.slice(0, index),
        updatedInput,
        ...this.state.investorInputs.slice(index + 1)
    ]
    this.setState({"investorInputs": newInputs})
  }

  handleChange(field, event) {
    this.setState({[field]: event.target.value})
  }

  render() {
    const investorInputs = this.state.investorInputs.map((input, index) => {
      return <InvestorInput
        key={input.id}
        id={input.id}
        index={index}
        removeAction={this.removeInvestor}
        onChange={this.handleInvestorChange}
        nameValue={input.name}
        requestedAmountValue={input.requestedAmount}
        averageAmountValue={input.averageAmount}
       />
    })

    return (
      <div className="centeredContainer">
        <h3 className="defaultTitle">Total Available Allocation</h3>
        <input
          className="inputDefault"
          type="number"
          name="name"
          placeholder="Allocation"
          onChange={this.handleChange.bind(this, "availableAllocation")}
        />
        {investorInputs}
        <button className="buttonDefault spacingTop" onClick={this.addInvestor.bind(this)}>Add Investor</button>
        <button className="buttonDefault spacingTop" onClick={this.prorate.bind(this)}>Prorate</button>
      </div>
    )
  }
}

export default AllocationInput
