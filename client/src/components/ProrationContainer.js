import React from 'react'
import { Col, Row, Container } from 'react-grid-system'
import AllocationInput from './AllocationInput'
import AllocationResult from './AllocationResult'

class ProrationContainer extends React.Component {
  constructor(props) {
    super(props)

    this.handleAllocationUpdate = this.handleAllocationUpdate.bind(this)

    this.state = {
      allocationResults: []
    }
  }

  handleAllocationUpdate(results) {
    this.setState({"allocationResults": results})
  }

  render() {
    return (
      <div className="centerAlignment">
        <Container className="containerSpacing centerAlignment" >
          <Row>
            <Col sm={6}>
              <AllocationInput
                onAllocationResult={this.handleAllocationUpdate}
               />
            </Col>

            <Col sm={6}>
              <AllocationResult
                results={this.state.allocationResults}
               />
            </Col>
          </Row>
        </Container>
      </div>
    )
  }
}

export default ProrationContainer
