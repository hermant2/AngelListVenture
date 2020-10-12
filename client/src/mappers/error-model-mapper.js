import {errorStrings} from "../strings/localized-strings";

const ErrorCodes = {
    GENERAL: 18930,
    INPUT_ZERO: 18931,
    NO_INVESTORS: 18932
}

export function mapErrorResponseMessage(error) {
    const errorResponse = error?.response?.data?.error
    if (!errorResponse) {
        return errorStrings.somethingWentWrong
    }

    switch (errorResponse.code) {
        case ErrorCodes.INPUT_ZERO:
            return errorStrings.inputZero
        case ErrorCodes.NO_INVESTORS:
            return errorStrings.noInvestors
        case ErrorCodes.GENERAL:
        default:
            return errorStrings.somethingWentWrong
    }
}
