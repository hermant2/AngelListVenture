import axios from 'axios'

export function calculateInvestmentAllocations(request) {
    return axios.post("http://localhost:8080/api/v1/prorate", request)
        .then(response => response.data)
}
