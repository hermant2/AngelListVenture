import LocalizedStrings from 'react-localization'

export const generalStrings = new LocalizedStrings({
    en: {
        siteTitle: "AngelList Venture"
    }
})

export const prorateStrings = new LocalizedStrings({
    en: {
        allocation: "Allocation",
        currencyPrefix: "$",
        totalAvailableAllocation: "Total Available Allocation",
        results: "Results",
        addInvestor: "Add Investor",
        prorate: "Prorate"
    }
})

export const investorStrings = new LocalizedStrings({
    en: {
        investorName: "Investor Name",
        requestedAmount: "Requested Amount",
        averageAmount: "Average Amount",
        remove: "Remove"
    }
})

export const errorStrings = new LocalizedStrings({
    en: {
        oopsTitle: "Oops!",
        somethingWentWrong: "Something went wrong.",
        inputZero: "All monetary inputs must be greater than zero.",
        noInvestors: "There must be at least one investor."
    }
})
