import axios from 'axios'

export function calculateInvestmentAllocations(request) {
    return axios.post("/api/v1/prorate", request)
        .then(response => response.data)
}
