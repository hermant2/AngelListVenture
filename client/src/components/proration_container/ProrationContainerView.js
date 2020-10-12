import React from 'react'
import {Col, Container, Row} from 'react-grid-system'
import AllocationInputView from '../allocation_input/AllocationInputView'
import AllocationResultView from '../allocation_result/AllocationResultView'
import ProrationContainerPresenter from "./ProrationContainerPresenter";
import ProrationContainerModel from "./ProrationContainerModel";

class ProrationContainerView extends React.Component {
    constructor(props) {
        super(props)
        const model = new ProrationContainerModel()
        this.state = {model: model}
        this.presenter = new ProrationContainerPresenter(this, model)
    }

    updateState(model) {
        this.setState({model: model})
    }

    render() {
        return (
            <div className="centerAlignment">
                <Container className="containerSpacing centerAlignment">
                    <Row>
                        <Col sm={6}>
                            <AllocationInputView
                                parentPresenter={this.presenter}
                            />
                        </Col>

                        <Col sm={6}>
                            <AllocationResultView
                                model={this.state.model.allocationResult}
                            />
                        </Col>
                    </Row>
                </Container>
            </div>
        )
    }
}

export default ProrationContainerView
