const US_LOCALE = "en-US"
const STYLE = "currency"
const CURRENCY_TYPE = "USD"

export function formatUSDCurrency(money) {
    return new Intl.NumberFormat(US_LOCALE, {
        style: STYLE,
        currency: CURRENCY_TYPE
    }).format(money)
}
