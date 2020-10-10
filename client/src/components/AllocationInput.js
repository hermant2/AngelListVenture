import React from 'react'
import uuid from 'react-uuid'
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
    let results = this.state.investorInputs.map(input => {
      return {
        id: input.id,
        name: input.name,
        allocation: input.requestedAmount
      }
    })
    this.props.onAllocationResult(results)
  }

  removeInvestor(id) {
    let newInputs = this.state.investorInputs.filter(input => input.id !== id)
    this.setState({"investorInputs": newInputs})
  }

  handleInvestorChange(index, field, value) {
    let updatedInput = this.state.investorInputs[index]
    updatedInput[field] = value
    let newInputs = this.state.investorInputs.slice(index, 1, updatedInput)
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
